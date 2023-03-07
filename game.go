package lucretia

import "fmt"

type Game struct {
	Sport       string
	NoOfPlayers uint
	IsRacket    bool
}

func (g *Game) Play() string {
	return fmt.Sprintf("🎯 Sport %v, #️⃣ No Of Players %v and 🎾Is Racket sport? %v",
		g.Sport, g.NoOfPlayers, g.IsRacket)
}
