package semver_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSemver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Semver Suite")
}
