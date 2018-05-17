package api_test

import (
	"github.com/kaleo211/semver/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Semver", func() {
	Describe("NewSemver", func() {
		Context("when version is empty", func() {
			It("should return default value", func() {
				s, err := api.NewSemver("")
				Expect(err).ToNot(HaveOccurred())
				Expect(s.Version()).To(Equal("0.0.0"))
			})
		})

		Context("when version is invalid", func() {
			It("should return error", func() {
				_, err := api.NewSemver("not valid")
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when version is not in correct format", func() {
			It("should return error", func() {
				_, err := api.NewSemver("1.2.3.4")
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when version is valid", func() {
			It("should return semver without error", func() {
				validSemver := "1.2.3"
				semver, err := api.NewSemver(validSemver)
				Expect(err).ToNot(HaveOccurred())
				Expect(semver).ToNot(BeNil())
				Expect(semver.Version()).To(Equal(validSemver))
			})
		})
	})

	Describe("Increment", func() {
		Context("when increment by patch", func() {
			It("should level patch", func() {
				semver, _ := api.NewSemver("1.2.3")
				v := semver.IncPatch()
				Expect(v).To(Equal("1.2.4"))
			})
		})

		Context("when increment by minor", func() {
			It("should level minor", func() {
				semver, _ := api.NewSemver("1.2.3")
				v := semver.IncMinor()
				Expect(v).To(Equal("1.3.0"))
			})
		})

		Context("when increment by major", func() {
			It("should level major", func() {
				semver, _ := api.NewSemver("1.2.3")
				v := semver.IncMajor()
				Expect(v).To(Equal("2.0.0"))
			})
		})
	})

	Describe("Clean", func() {
		Context("when string has valid version inside", func() {
			It("should return valid version", func() {
				version, err := api.Clean(" *^%v2.11.4*&^  ")
				Expect(err).ToNot(HaveOccurred())
				Expect(version).To(Equal("2.11.4"))
			})
		})

		Context("when string is valid version", func() {
			It("should return valid version", func() {
				version, err := api.Clean("2.11.4")
				Expect(err).ToNot(HaveOccurred())
				Expect(version).To(Equal("2.11.4"))
			})
		})

		Context("when string had no valid version", func() {
			It("should return error", func() {
				_, err := api.Clean("2.11.d4")
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Validate", func() {
		Context("when version is not valid", func() {
			It("should return error", func() {
				_, err := api.Validate("a2.1b.1c")
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when version is valid", func() {
			It("should return error", func() {
				version, err := api.Validate("2.1.1")
				Expect(err).ToNot(HaveOccurred())
				Expect(version).To(Equal("2.1.1"))
			})
		})
	})
})
