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

// Increment increments
func (s *Semver) Increment(l Level) (string, error) {
	switch l {
	case Patch:
		s.patch++
	case Minor:
		s.patch = 0
		s.minor++
	case Major:
		s.patch = 0
		s.minor = 0
		s.major++
	default:
		return "", errors.New("unsupported level")
	}

	return s.Version(), nil
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
