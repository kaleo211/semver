package semver

// ISemver is ...
type ISemver interface {
	Version() string
	Increment(int) (string, error)
}
