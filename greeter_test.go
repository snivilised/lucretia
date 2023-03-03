package lucretia_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/snivilised/lucretia"
)

var _ = Describe("Greeter", func() {
	Context("Greet", func() {
		It("should: greet user", func() {
			greeter := lucretia.Greeter{
				Title:    "Mr",
				Forename: "Scooby",
				Surname:  "Do",
			}
			Expect(greeter.Hello()).ToNot(BeNil())
		})
	})
})
