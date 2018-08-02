package session

import (
	"crypto/sha1"
	"fmt"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/havoc-io/mutagen/pkg/filesystem"
)

const (
	// sessionsDirectoryName is the name of the sessions subdirectory within the
	// Mutagen directory.
	sessionsDirectoryName = "sessions"
	// archivesDirectoryName is the name of the archives subdirectory within the
	// Mutagen directory.
	archivesDirectoryName = "archives"
	// cachesDirectoryName is the name of the caches subdirectory within the
	// Mutagen directory.
	cachesDirectoryName = "caches"
	// stagingDirectoryName is the name of the staging subdirectory within the
	// Mutagen directory.
	stagingDirectoryName = "staging"

	// alphaName is the name to use for alpha when distinguishing endpoints.
	alphaName = "alpha"
	// betaName is the name to use for beta when distinguishing endpoints.
	betaName = "beta"

	// stagingPrefixLength is the byte length to use for prefix directories when
	// load-balancing staged files.
	stagingPrefixLength = 1
)

// pathForSession computes the path to the serialized session for the given
// session identifier. An empty session identifier will return the sessions
// directory path.
func pathForSession(sessionIdentifier string) (string, error) {
	// Compute/create the sessions directory.
	sessionsDirectoryPath, err := filesystem.Mutagen(true, sessionsDirectoryName)
	if err != nil {
		return "", errors.Wrap(err, "unable to compute/create sessions directory")
	}

	// Success.
	return filepath.Join(sessionsDirectoryPath, sessionIdentifier), nil
}

// pathForArchive computes the path to the serialized archive for the given
// session identifier.
func pathForArchive(session string) (string, error) {
	// Compute/create the archives directory.
	archivesDirectoryPath, err := filesystem.Mutagen(true, archivesDirectoryName)
	if err != nil {
		return "", errors.Wrap(err, "unable to compute/create archives directory")
	}

	// Success.
	return filepath.Join(archivesDirectoryPath, session), nil
}

// pathForCache computes the path to the serialized cache for the given
// session identifier and endpoint.
func pathForCache(session string, alpha bool) (string, error) {
	// Compute/create the caches directory.
	cachesDirectoryPath, err := filesystem.Mutagen(true, cachesDirectoryName)
	if err != nil {
		return "", errors.Wrap(err, "unable to compute/create caches directory")
	}

	// Compute the endpoint name.
	endpointName := alphaName
	if !alpha {
		endpointName = betaName
	}

	// Compute the cache name.
	cacheName := fmt.Sprintf("%s_%s", session, endpointName)

	// Success.
	return filepath.Join(cachesDirectoryPath, cacheName), nil
}

// pathForStagingRoot computes and creates the path to the staging root for the
// given session identifier and endpoint.
func pathForStagingRoot(session string, alpha bool) (string, error) {
	// Compute the endpoint name.
	endpointName := alphaName
	if !alpha {
		endpointName = betaName
	}

	// Compute the staging root name.
	stagingRootName := fmt.Sprintf("%s_%s", session, endpointName)

	// Compute the staging root, but don't create it.
	return filesystem.Mutagen(false, stagingDirectoryName, stagingRootName)
}

// pathForStaging computes the staging path for the specified path/digest. It
// returns the prefix directory name but does not ensure that it's been created.
func pathForStaging(root, path string, digest []byte) (string, string, error) {
	// Compute the prefix for the entry.
	if len(digest) == 0 {
		return "", "", errors.New("entry digest too short")
	}
	prefix := fmt.Sprintf("%x", digest[:1])

	// Compute the staging name.
	stagingName := fmt.Sprintf("%x_%x", sha1.Sum([]byte(path)), digest)

	// Success.
	return filepath.Join(root, prefix, stagingName), prefix, nil
}
