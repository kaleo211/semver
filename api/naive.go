package api

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// Regular expression for number
var nre = "[0-9]|[1-9][0-9]*"

// Semver is ...
type Semver struct {
	major int
	minor int
	patch int
}

// NewSemver creates
func NewSemver(v string) (*Semver, error) {
	var major, minor, patch int
	var err error

	if v == "" {
		v = "0.0.0"
	} else {
		major, minor, patch, err = validate(v)
		if err != nil {
			return nil, err
		}
	}

	return &Semver{major, minor, patch}, nil
}

// Version returns
func (s *Semver) Version() string {
	return fmt.Sprintf("%d.%d.%d", s.major, s.minor, s.patch)
}

// IncMajor increments
func (s *Semver) IncMajor() string {
	s.patch, s.minor = 0, 0
	s.major++
	return s.Version()
}

// IncMinor increments
func (s *Semver) IncMinor() string {
	s.patch = 0
	s.minor++
	return s.Version()
}

// IncPatch increments
func (s *Semver) IncPatch() string {
	s.patch++
	return s.Version()
}

// Major returns major
func (s *Semver) Major() int {
	return s.major
}

// Minor returns minor
func (s *Semver) Minor() int {
	return s.minor
}

// Patch returns patch
func (s *Semver) Patch() int {
	return s.patch
}

// GT compare whether larger than passed version
func (s *Semver) GT(v *Semver) bool {
	if s.major != v.major {
		return s.major > v.major
	}

	if s.minor != v.minor {
		return s.minor > v.minor
	}

	return s.patch > v.patch
}

// Validate version if it's valid
func Validate(v string) (string, error) {
	s, err := NewSemver(v)
	if err != nil {
		return "", err
	}

	return s.Version(), nil
}

// Clean none needed characters
func Clean(v string) (string, error) {
	versionRegexp := fmt.Sprintf("[^0-9]*((%s)\\.(%s)\\.(%s))[^0-9]*", nre, nre, nre)
	re := regexp.MustCompile(versionRegexp)

	matched := re.FindAllStringSubmatch(v, -1)
	if len(matched) < 1 || len(matched[0]) < 2 {
		return "", errors.New("clean: version is not valid")
	}

	return matched[0][1], nil
}

func validate(v string) (int, int, int, error) {
	versionRegexp := fmt.Sprintf("^(%s)\\.(%s)\\.(%s)$", nre, nre, nre)
	re := regexp.MustCompile(versionRegexp)

	matched := re.FindAllStringSubmatch(v, -1)
	if len(matched) != 1 || len(matched[0]) != 4 {
		return 0, 0, 0, errors.New("version is not valid ")
	}

	major, err := strconv.Atoi(matched[0][1])
	if err != nil {
		return 0, 0, 0, errors.New("version major is not valid")
	}
	minor, _ := strconv.Atoi(matched[0][2])
	if err != nil {
		return 0, 0, 0, errors.New("version minor is not valid")
	}
	patch, _ := strconv.Atoi(matched[0][3])
	if err != nil {
		return 0, 0, 0, errors.New("version patch is not valid")
	}

	return major, minor, patch, nil
}
