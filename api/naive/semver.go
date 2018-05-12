package naive

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Level int

const (
	Patch Level = iota
	Minor
	Major
)

// Semver is ...
type Semver struct {
	version string
	major   int
	minor   int
	patch   int
}

// NewSemver creates
func NewSemver(v string) (*Semver, error) {
	var version string
	if v == "" {
		version = "0.0.0"
	}

	major, minor, patch, err := validate(v)
	if err != nil {
		return nil, err
	}

	fmt.Println(major, minor, patch)

	return &Semver{version, major, minor, patch}, nil
}

// Version returns
func (s *Semver) Version() string {
	return s.version
}

// Increment increments
func (s *Semver) Increment(l Level) (string, error) {
	return s.version, nil
}

// Validate validates
func validate(v string) (int, int, int, error) {
	re := regexp.MustCompile(`([0-9]|[1-9][0-9]*)\.([0-9]|[1-9][0-9]*)\.([0-9]|[1-9][0-9]*)`)

	matched := re.FindAllStringSubmatch(v, -1)
	if len(matched) != 1 || len(matched[0]) != 4 {
		return 0, 0, 0, errors.New("version is not valid")
	}

	major, _ := strconv.Atoi(matched[0][1])
	minor, _ := strconv.Atoi(matched[0][2])
	patch, _ := strconv.Atoi(matched[0][3])
	return major, minor, patch, nil
}
