package filesystem

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/pkg/errors"

	"golang.org/x/sys/windows"

	"github.com/hectane/go-acl"
	aclapi "github.com/hectane/go-acl/api"
)

// pathSeparator is a byte representation of the OS path separator. We rely on
// this being a single byte for performance in ensureNoPathSeparator. For each
// platform where we make this assumption, we have a test to ensure that it's
// valid.
var pathSeparator = byte(os.PathSeparator)

// pathSeparatorAlternate is a byte representation of the alternate OS path
// separator. We rely on this being a single byte for performance in
// ensureNoPathSeparator. For each platform where we make this assumption, we
// have a test to ensure that it's valid.
var pathSeparatorAlternate = byte('/')

// ensureValidName verifies that the provided name does not reference the
// current directory, the parent directory, or contain a path separator
// character.
func ensureValidName(name string) error {
	// Verify that the name does not reference the directory itself or the
	// parent directory.
	if name == "." {
		return errors.New("name is directory reference")
	} else if name == ".." {
		return errors.New("name is parent directory reference")
	}

	// Verify that neither of the path separator characters appears in the name.
	if strings.IndexByte(name, pathSeparator) != -1 {
		return errors.New("path separator appears in name")
	} else if strings.IndexByte(name, pathSeparatorAlternate) != -1 {
		return errors.New("alternate path separator appears in name")
	}

	// Success.
	return nil
}

// Directory represents a directory on disk and provides race-free operations on
// the directory's contents. All of its operations avoid the traversal of
// symbolic links.
type Directory struct {
	// handle is the underlying Windows HANDLE object that has been opened
	// without the FILE_SHARE_DELETE to ensure that the directory is immovable.
	handle windows.Handle
	// file is the underlying os.File object corresponding to the directory. On
	// Windows systems, it is not actually a wrapper around the handle object,
	// but rather around a search handle generated by the FindFirstFileExW
	// function, hence the reason we need to hold open a separate HANDLE. It is
	// guaranteed that the value returned from the file's Name function will be
	// an absolute path.
	file *os.File
}

// Close closes the directory.
func (d *Directory) Close() error {
	// Close the file object.
	if err := d.file.Close(); err != nil {
		windows.CloseHandle(d.handle)
		return errors.Wrap(err, "unable to close file object")
	}

	// Close the handle.
	if err := windows.CloseHandle(d.handle); err != nil {
		return errors.Wrap(err, "unable to close file handle")
	}

	// Success.
	return nil
}

// CreateDirectory creates a new directory with the specified name inside the
// directory. The directory will be created with user-only read/write/execute
// permissions.
func (d *Directory) CreateDirectory(name string) error {
	// Verify that the name is valid.
	if err := ensureValidName(name); err != nil {
		return err
	}

	// Create the directory.
	return os.Mkdir(filepath.Join(d.file.Name(), name), 0700)
}

// CreateTemporaryFile creates a new temporary file using the specified name
// pattern inside the directory. Pattern behavior follows that of
// io/ioutil.TempFile. The file will be created with user-only read/write
// permissions.
func (d *Directory) CreateTemporaryFile(pattern string) (string, WritableFile, error) {
	// Verify that the name is valid. This should still be a sensible operation
	// for pattern specifications.
	if err := ensureValidName(pattern); err != nil {
		return "", nil, err
	}

	// Create the temporary file using the standard io/ioutil implementation.
	file, err := ioutil.TempFile(d.file.Name(), pattern)
	if err != nil {
		return "", nil, err
	}

	// Extract the base name of the file.
	name := filepath.Base(file.Name())

	// Success.
	return name, file, nil
}

// CreateSymbolicLink creates a new symbolic link with the specified name and
// target inside the directory. The symbolic link is created with the default
// system permissions (which, generally speaking, don't apply to the symbolic
// link itself).
func (d *Directory) CreateSymbolicLink(name, target string) error {
	// Verify that the name is valid.
	if err := ensureValidName(name); err != nil {
		return err
	}

	// Create the symbolic link.
	return os.Symlink(target, filepath.Join(d.file.Name(), name))
}

// SetPermissions sets the permissions on the content within the directory
// specified by name. Ownership information is set first, followed by
// permissions extracted from the mode using ModePermissionsMask. Ownership
// setting can be skipped completely by providing a nil OwnershipSpecification
// or a specification with both components unset. An OwnershipSpecification may
// also include only certain components, in which case only those components
// will be set. Permission setting can be skipped by providing a mode value that
// yields 0 after permission bit masking.
func (d *Directory) SetPermissions(name string, ownership *OwnershipSpecification, mode Mode) error {
	// Verify that the name is valid.
	if err := ensureValidName(name); err != nil {
		return err
	}

	// Compute the target path.
	path := filepath.Join(d.file.Name(), name)

	// Set ownership information, if specified.
	if ownership != nil && (ownership.userSID != nil || ownership.groupSID != nil) {
		// Compute the information that we're going to set.
		var information uint32
		if ownership.userSID != nil {
			information |= aclapi.OWNER_SECURITY_INFORMATION
		}
		if ownership.groupSID != nil {
			information |= aclapi.GROUP_SECURITY_INFORMATION
		}

		// Set the information.
		if err := aclapi.SetNamedSecurityInfo(
			path,
			aclapi.SE_FILE_OBJECT,
			information,
			ownership.userSID,
			ownership.groupSID,
			0,
			0,
		); err != nil {
			return errors.Wrap(err, "unable to set ownership information")
		}
	}

	// Set permissions, if specified.
	mode = mode & ModePermissionsMask
	if mode != 0 {
		if err := acl.Chmod(path, os.FileMode(mode)); err != nil {
			return errors.Wrap(err, "unable to set permission bits")
		}
	}

	// Success.
	return nil
}

// openHandle is the underlying open implementation shared by OpenDirectory and
// OpenFile.
func (d *Directory) openHandle(name string, wantDirectory bool) (string, windows.Handle, error) {
	// Verify that the name is valid.
	if err := ensureValidName(name); err != nil {
		return "", 0, err
	}

	// Compute the full path.
	path := filepath.Join(d.file.Name(), name)

	// Convert the path to UTF-16.
	path16, err := windows.UTF16PtrFromString(path)
	if err != nil {
		return "", 0, errors.Wrap(err, "unable to convert path to UTF-16")
	}

	// Open the path in a manner that is suitable for reading, doesn't allow for
	// other threads or processes to delete or rename the file while open,
	// avoids symbolic link traversal (at the path leaf), and has suitable
	// semantics for both files and directories.
	handle, err := windows.CreateFile(
		path16,
		windows.GENERIC_READ,
		windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE,
		nil,
		windows.OPEN_EXISTING,
		windows.FILE_ATTRIBUTE_NORMAL|windows.FILE_FLAG_BACKUP_SEMANTICS|windows.FILE_FLAG_OPEN_REPARSE_POINT,
		0,
	)
	if err != nil {
		return "", 0, errors.Wrap(err, "unable to open path")
	}

	// Query file metadata.
	var metadata windows.ByHandleFileInformation
	if err := windows.GetFileInformationByHandle(handle, &metadata); err != nil {
		windows.CloseHandle(handle)
		return "", 0, errors.Wrap(err, "unable to query file metadata")
	}

	// Verify that the handle does not represent a symbolic link and that the
	// type coincides with what we want. Note that FILE_ATTRIBUTE_REPARSE_POINT
	// can be or'd with FILE_ATTRIBUTE_DIRECTORY (since symbolic links are
	// "typed" on Windows), so we have to explicitly exclude reparse points
	// before checking types.
	//
	// TODO: Are there additional attributes upon which we should reject here?
	// The Go os.File implementation doesn't seem to for normal os.Open
	// operations, so I guess we don't need to either, but we should keep the
	// option in mind.
	if metadata.FileAttributes&windows.FILE_ATTRIBUTE_REPARSE_POINT != 0 {
		windows.CloseHandle(handle)
		return "", 0, errors.New("path pointed to symbolic link")
	} else if wantDirectory && metadata.FileAttributes&windows.FILE_ATTRIBUTE_DIRECTORY == 0 {
		windows.CloseHandle(handle)
		return "", 0, errors.New("path pointed to non-directory location")
	}

	// Success.
	return path, handle, nil
}

// OpenDirectory opens the directory within the directory specified by name.
func (d *Directory) OpenDirectory(name string) (*Directory, error) {
	// Open the directory handle.
	path, handle, err := d.openHandle(name, true)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open directory handle")
	}

	// Open the corresponding file object. Unfortunately we can't force the file
	// to use the base name in order to keep consistency with file os.File
	// objects on Windows and file/directory os.File POSIX, but it's okay since
	// this object (and its Name method) isn't exposed anyway.
	file, err := os.Open(path)
	if err != nil {
		windows.CloseHandle(handle)
		return nil, errors.Wrap(err, "unable to open file object for directory")
	}

	// Success.
	return &Directory{
		handle: handle,
		file:   file,
	}, nil
}

// ReadContentNames queries the directory contents and returns their base names.
// It does not return "." or ".." entries.
func (d *Directory) ReadContentNames() ([]string, error) {
	// Read content names. Fortunately we can use the os.File implementation for
	// this since it operates on the underlying file descriptor directly.
	names, err := d.file.Readdirnames(0)
	if err != nil {
		return nil, err
	}

	// Filter names (without allocating a new slice).
	results := names[:0]
	for _, name := range names {
		// Watch for names that reference the directory itself or the parent
		// directory. The implementation underlying os.File.Readdirnames does
		// filter these out, but that's not guaranteed by its documentation, so
		// it's better to do this explicitly.
		if name == "." || name == ".." {
			continue
		}

		// Store the name.
		results = append(results, name)
	}

	// Success.
	return names, nil
}

// ReadContentMetadata reads metadata for the content within the directory
// specified by name.
func (d *Directory) ReadContentMetadata(name string) (*Metadata, error) {
	// Verify that the name is valid.
	if err := ensureValidName(name); err != nil {
		return nil, err
	}

	// Query metadata.
	metadata, err := os.Lstat(filepath.Join(d.file.Name(), name))
	if err != nil {
		return nil, err
	}

	// Success.
	return &Metadata{
		Name:             name,
		Mode:             Mode(metadata.Mode()),
		Size:             uint64(metadata.Size()),
		ModificationTime: metadata.ModTime(),
	}, nil
}

// ReadContents queries the directory contents and their associated metadata.
// While the results of this function can be computed as a combination of
// ReadContentNames and ReadContentMetadata (and this is indeed the mechanism by
// which this function is implemented on POSIX systems), it may be significantly
// faster than a naïve combination implementation on some platforms
// (specifically Windows, where it relies on FindFirstFile/FindNextFile
// infrastructure). This function doesn't not return metadata for "." or ".."
// entries.
func (d *Directory) ReadContents() ([]*Metadata, error) {
	// Read directory content. On Windows, we use the os.File implementation to
	// read names and (an acceptable amount of metadata) in one fell swoop,
	// rather than using a "read names + loop and query" construct. The reason
	// for this is that Windows file metadata queries are extremely slow,
	// requiring use of either GetFileInformationByHandle (which requires
	// opening the file) or GetFileAttributesEx (which I'm fairly sure uses the
	// first function under the hood). Instead, os.File.Readdir uses
	// FindFirstFile/FindNextFile infrastructure under the hood (in fact os.File
	// is just a search handle for directory objects on Windows), which is much
	// faster and retrieves just enough of the necessary metadata.
	contents, err := d.file.Readdir(0)
	if err != nil {
		return nil, err
	}

	// Allocate the result slice with enough capacity to accommodate all
	// entries.
	results := make([]*Metadata, 0, len(contents))

	// Loop over contents.
	for _, content := range contents {
		// Watch for names that reference the directory itself or the parent
		// directory. The implementation underlying os.File.Readdir does seem to
		// filter these out, but that's not guaranteed by its documentation, so
		// it's better to do this explicitly.
		name := content.Name()
		if name == "." || name == ".." {
			continue
		}

		// Convert and append the metadata. Unfortunately we can't populate
		// FileID and DeviceID because the FindFirstFile/FindNextFile
		// infrastructure used by the os package doesn't provide access to this
		// information. We'd have to open each file and use
		// GetFileInformationByHandle, which is just way too expensive.
		results = append(results, &Metadata{
			Name:             name,
			Mode:             Mode(content.Mode()),
			Size:             uint64(content.Size()),
			ModificationTime: content.ModTime(),
		})
	}

	// Success.
	return results, nil
}

// OpenFile opens the file within the directory specified by name.
func (d *Directory) OpenFile(name string) (ReadableFile, error) {
	// Open the file handle.
	_, handle, err := d.openHandle(name, false)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open file handle")
	}

	// Wrap the file handle in an os.File object. We use the base name for the
	// file since that's the name that was used to "open" the file, which is
	// what os.File.Name is supposed to return (even though we don't expose
	// os.File.Name).
	file := os.NewFile(uintptr(handle), name)

	// Success.
	return file, nil
}

// ReadSymbolicLink reads the target of the symbolic link within the directory
// specified by name.
func (d *Directory) ReadSymbolicLink(name string) (string, error) {
	// Verify that the name is valid.
	if err := ensureValidName(name); err != nil {
		return "", err
	}

	// Read the symbolic link.
	return os.Readlink(filepath.Join(d.file.Name(), name))
}

// RemoveDirectory deletes a directory with the specified name inside the
// directory. The removal target must be empty.
func (d *Directory) RemoveDirectory(name string) error {
	// Verify that the name is valid.
	if err := ensureValidName(name); err != nil {
		return err
	}

	// Compute the full path.
	path := filepath.Join(d.file.Name(), name)

	// Convert the path to UTF-16.
	path16, err := windows.UTF16PtrFromString(path)
	if err != nil {
		return errors.Wrap(err, "unable to convert path to UTF-16")
	}

	// Remove the directory.
	return windows.RemoveDirectory(path16)
}

// RemoveFile deletes a file with the specified name inside the directory.
func (d *Directory) RemoveFile(name string) error {
	// Verify that the name is valid.
	if err := ensureValidName(name); err != nil {
		return err
	}

	// Compute the full path.
	path := filepath.Join(d.file.Name(), name)

	// Convert the path to UTF-16.
	path16, err := windows.UTF16PtrFromString(path)
	if err != nil {
		return errors.Wrap(err, "unable to convert path to UTF-16")
	}

	// Remove the directory.
	return windows.DeleteFile(path16)
}

// RemoveSymbolicLink deletes a symbolic link with the specified name inside the
// directory.
func (d *Directory) RemoveSymbolicLink(name string) error {
	return d.RemoveFile(name)
}

// Rename performs an atomic rename operation from one filesystem location (the
// source) to another (the target). Each location can be specified in one of two
// ways: either by a combination of directory and (non-path) name or by path
// (with corresponding nil Directory object). Different specification mechanisms
// can be used for each location.
//
// This function does not support cross-device renames. To detect whether or not
// an error is due to an attempted cross-device rename, use the
// IsCrossDeviceError function.
func Rename(
	sourceDirectory *Directory, sourceNameOrPath string,
	targetDirectory *Directory, targetNameOrPath string,
) error {
	// Adjust the source path if necessary.
	if sourceDirectory != nil {
		if err := ensureValidName(sourceNameOrPath); err != nil {
			return errors.Wrap(err, "source name invalid")
		}
		sourceNameOrPath = filepath.Join(sourceDirectory.file.Name(), sourceNameOrPath)
	}

	// Adjust the target path if necessary.
	if targetDirectory != nil {
		if err := ensureValidName(targetNameOrPath); err != nil {
			return errors.Wrap(err, "target name invalid")
		}
		targetNameOrPath = filepath.Join(targetDirectory.file.Name(), targetNameOrPath)
	}

	// Perform an atomic rename.
	return os.Rename(sourceNameOrPath, targetNameOrPath)
}

const (
	// _ERROR_NOT_SAME_DEVICE is the error code returned by MoveFileEx on
	// Windows when attempting to move a file across devices. This can actually
	// be avoided on Windows by specifying the MOVEFILE_COPY_ALLOWED flag, but
	// Go's standard library doesn't do this (most likely to keep consistency
	// with POSIX, which has no such facility). Since we don't know the exact
	// mechanism by which MOVEFILE_COPY_ALLOWED (i.e. whether or not it uses an
	// intermediate file), we avoid using it via a direct call to MoveFileEx.
	_ERROR_NOT_SAME_DEVICE = 0x11
)

// IsCrossDeviceError checks whether or not an error returned from rename
// represents a cross-device error.
func IsCrossDeviceError(err error) bool {
	if linkErr, ok := err.(*os.LinkError); !ok {
		return false
	} else if errno, ok := linkErr.Err.(syscall.Errno); !ok {
		return false
	} else {
		return errno == _ERROR_NOT_SAME_DEVICE
	}
}