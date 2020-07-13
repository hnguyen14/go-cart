package main_test

import (
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/hnguyen14/go-cart"
)

var _ = Describe("Parsers", func() {
	Describe("HTMLParser", func() {
		Context("When parsing a wiki page", func() {
			It("should get all the links", func() {
				content, err := ioutil.ReadFile("./testdata/ma_chao.html")
				Expect(err).To(BeNil())
				expected := Page{
					Links: []*Link{
						{
							Tag: "anchor1",
							URL: "http://test.test",
						},
						{
							Tag: "anchor1",
							URL: "http://test.test",
						},
					},
				}
				actual, err := ParseHTML(content)

				Expect(err).To(BeNil())
				Expect(*actual).To(Equal(expected))
			})
		})
	})
})
