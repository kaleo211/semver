package semver

// ISemver is ...
type ISemver interface {
	Version() string
	IncMajor() string
	IncMinor() string
	IncPatch() string
	Major() int
	Minor() int
	Patch() int
}
