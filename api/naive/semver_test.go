package naive_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/kaleo211/semver/api/naive"
)

var _ = Describe("Semver", func() {

	Describe("NewSemver", func() {
		Context("when version is empty", func() {
			It("should return default value", func() {
				s, err := naive.NewSemver("")
				Expect(err).ToNot(HaveOccurred())
				Expect(s.Version()).To(Equal("0.0.0"))
			})
		})

		Context("when version is invalid", func() {
			It("should return error", func() {
				_, err := naive.NewSemver("not valid")
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when version is valid", func() {
			It("should return semver without error", func() {
				validSemver := "1.2.3"
				semver, err := naive.NewSemver(validSemver)
				Expect(err).ToNot(HaveOccurred())
				Expect(semver).ToNot(BeNil())
				Expect(semver.Version()).To(Equal(validSemver))
			})
		})
	})

	Describe("Increment", func() {
		Context("when increment by patch", func() {
			It("should level patch", func() {
				semver, _ := naive.NewSemver("1.2.3")
				v, err := semver.Increment(naive.Patch)
				Expect(err).ToNot(HaveOccurred())
				Expect(v).To(Equal("1.2.4"))
			})
		})

		Context("when increment by minor", func() {
			It("should level minor", func() {
				semver, _ := naive.NewSemver("1.2.3")
				v, err := semver.Increment(naive.Minor)
				Expect(err).ToNot(HaveOccurred())
				Expect(v).To(Equal("1.3.0"))
			})
		})

		Context("when increment by major", func() {
			It("should level major", func() {
				semver, _ := naive.NewSemver("1.2.3")
				v, err := semver.Increment(naive.Major)
				Expect(err).ToNot(HaveOccurred())
				Expect(v).To(Equal("2.0.0"))
			})
		})
	})
})
