package naive

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
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
func (s *Semver) Increment(l string) (string, error) {
	switch strings.ToLower(l) {
	case "patch":
		s.patch++
	case "minor":
		s.patch = 0
		s.minor++
	case "major":
		s.patch = 0
		s.minor = 0
		s.major++
	default:
		return "", errors.New("unsupported level")
	}

	return s.Version(), nil
}

func (s *Semver) GetMajor() int {
	return s.major
}

func (s *Semver) GetMinor() int {
	return s.minor
}

func (s *Semver) GetPatch() int {
	return s.patch
}

func validate(v string) (int, int, int, error) {
	re := regexp.MustCompile(`([0-9]|[1-9][0-9]*)\.([0-9]|[1-9][0-9]*)\.([0-9]|[1-9][0-9]*)`)

	matched := re.FindAllStringSubmatch(v, -1)
	if len(matched) != 1 || len(matched[0]) != 4 {
		return 0, 0, 0, errors.New("version is not valid")
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
