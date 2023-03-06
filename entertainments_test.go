package lucretia_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/snivilised/lucretia"
)

var _ = Describe("Entertainments", func() {
	It("should: entertain", func() {
		entertain := lucretia.Entertainments{
			StandUp: "Seinfeld",
			Musical: "Moulin Rouge",
			Poetry:  "Kae Tempest",
		}
		Expect(entertain.Enact()).ToNot(BeNil())
	})
})
