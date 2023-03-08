package lucretia

import "fmt"

type Greeter struct {
	Forename string
	Surname  string
	Title    string
	Message  string
	Age      uint
}

func (g *Greeter) Hello() string {
	prompt := "What's your poison"
	if g.Age <= 18 {
		prompt = "go home whipper snapper, no 🥣 soup for you"
	}
	return fmt.Sprintf("---> 👋👋 Hello %v. %v %v: %v 🧿 %v",
		g.Title, g.Forename, g.Surname, g.Message, prompt,
	)
}
