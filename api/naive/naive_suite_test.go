package naive_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNaive(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Naive Suite")
}
