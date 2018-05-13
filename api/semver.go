package semver

// ISemver is ...
type ISemver interface {
	Version() string
	Increment(string) (string, error)
	GetMajor() int
	GetMinor() int
	GetPatch() int
}
