package versions

var (
	// commitFromGit is a constant representing the source versions that
	// generated this build. It should be set during build via -ldflags.
	CommitFromGit string
	// versionFromGit is a constant representing the versions tag that
	// generated this build. It should be set during build via -ldflags.
	VersionFromGit = "unknown"
	// major versions
	MajorFromGit string
	// minor versions
	MinorFromGit string
	// build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
	BuildDate string
	// state of git tree, either "clean" or "dirty"
	GitTreeState string
)
