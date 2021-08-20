package main

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	// Git SHA1 commit hash of the release (set via linker flags).
	gitCommit = ""
	gitDate   = ""
)

// version of the current release
func version() string {
	params.VersionMajor = 0
	params.VersionMinor = 1
	params.VersionPatch = 0
	params.VersionMeta = "rc.1"
	return params.VersionWithCommit(gitCommit, gitDate)
}
