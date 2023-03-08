package lucretia_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/snivilised/lucretia"
)

var _ = Describe("Game", func() {
	It("should: greet user", func() {
		game := lucretia.Game{
			Sport:       "Tennis",
			NoOfPlayers: 2,
			IsRacket:    true,
			IsOlympic:   true,
		}
		Expect(game.Play()).ToNot(BeNil())
	})
})
