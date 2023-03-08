package lucretia

import "fmt"

type Game struct {
	Sport       string
	NoOfPlayers uint
	IsRacket    bool
	IsOlympic   bool
}

func (g *Game) Play() string {
	return fmt.Sprintf("🎯 Sport %v, #️⃣ No Of Players %v, 🎾Is Racket sport? %v, Is Olympic? %v",
		g.Sport, g.NoOfPlayers, g.IsRacket, g.IsOlympic)
}
