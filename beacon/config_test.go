package beacon_test

import (
	"encoding/json"

	. "github.com/naveego/api/beacon"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {

	var (
		sut *Config
	)

	sut = &Config{}
	const jsonData = `
	{
		"level1": {
			"level2": {
				"string12": "a12",
				"number12": 42
			}
		},
		"string": "a",
		"number": 17
	}	
	`

	BeforeEach(func() {
		Expect(json.Unmarshal([]byte(jsonData), sut)).To(Succeed())
	})

	Describe("GetConfig", func() {
		It("should find string entry", func() {
			Expect(sut.GetConfig("string")).To(Equal("a"))
		})

		It("should find number entry", func() {
			Expect(sut.GetConfig("number")).To(Equal(float64(17)))
		})

		It("should get nested number entry", func() {
			Expect(sut.GetConfig("level1.level2.number12")).To(Equal(float64(42)))
		})

	})

	Describe("GetConfigWithDefault", func() {
		It("should find string entry", func() {
			Expect(sut.GetConfig("string")).To(Equal("a"))
		})

		It("should find number entry", func() {
			Expect(sut.GetConfig("number")).To(Equal(float64(17)))
		})

		It("should get nested number entry", func() {
			Expect(sut.GetConfig("level1.level2.number12")).To(Equal(float64(42)))
		})

	})

	Describe("GetStringConfig", func() {

		It("should find typed string entry", func() {
			actual, ok := sut.GetStringConfig("string")
			Expect(ok).To(BeTrue())
			Expect(actual).To(Equal("a"))
		})

		It("should return false if entry is not a string", func() {
			_, ok := sut.GetStringConfig("number")
			Expect(ok).To(BeFalse())
		})

		It("should return false if entry is missing", func() {
			_, ok := sut.GetStringConfig("invalid")
			Expect(ok).To(BeFalse())
		})

		It("should return false if entry is missing", func() {
			_, ok := sut.GetStringConfig("invalid")
			Expect(ok).To(BeFalse())
		})

		It("should get nested string entry", func() {
			actual, ok := sut.GetStringConfig("level1.level2.string12")
			Expect(ok).To(BeTrue())
			Expect(actual).To(Equal("a12"))
		})
	})

})
