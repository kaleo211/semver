package api

// ISemver is ...
type ISemver interface {
	Version() string

	IncMajor() string
	IncMinor() string
	IncPatch() string
	Major() int
	Minor() int
	Patch() int

	GT(ISemver) bool
	LT(ISemver) bool
	EQ(ISemver) bool
	Compare(ISemver) int
}
