package semver

// ISemver is ...
type ISemver interface {
	Version() string
	Increment(string) (string, error)
}

// Semver is ...
type Semver struct {
	version string
}

// NewSemver creates
func NewSemver(v string) *Semver {

	return &Semver{
		version: v,
	}
}

// Version returns
func (s *Semver) Version() string {
	return s.version
}

// Increment increments
func (s *Semver) Increment(level string) (string, error) {

	return
}

// Validate validates
func (s *Semver) Validate() error {
	return nil
}
