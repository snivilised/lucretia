package lucretia_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/snivilised/lucretia"
)

var _ = Describe("Refreshment", func() {
	It("should: greet user", func() {
		refresh := lucretia.Refreshment{
			Aperitif: "Sparkling",
			Vintage:  "Rioja",
			Chaser:   "Goldslagger",
		}
		Expect(refresh.Serve()).ToNot(BeNil())
	})
})
