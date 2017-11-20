package beacon_test

import (
	"fmt"
	"path/filepath"

	. "github.com/naveego/api/beacon"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {

	Describe("Integration", func() {

		var (
			sut Client
			err error
		)

		BeforeEach(func() {
			sut, err = NewClient("http://localhost:9005", "")
			Expect(err).To(Succeed())
		})

		It("when no path is provided it should get a default config", func() {
			actual, err := sut.GetConfig("")
			Expect(err).ToNot(HaveOccurred())
			data, _ := actual.GetStringConfig("endpoints.graylog.http")
			Expect(data).To(Equal("http://graylog.n5o.green/gelf"))
		})
	})
})

var _ = Describe("Package", func() {

	Describe("Integration", func() {

		Describe("GetConfig", func() {
			It("should get config from file path", func() {
				path := "./data/test.json"
				normalizedPath, err := filepath.Abs(path)
				Expect(err).ToNot(HaveOccurred())
				fmt.Println("normalizedPath", normalizedPath)
				urlPath := "file://" + normalizedPath
				config, _ := GetConfig(urlPath, "")
				actual := config.GetConfig("endpoints.graylog.http")
				Expect(actual).To(Equal("http://graylog.n5o.green/gelf"))
			})
		})
	})
})
