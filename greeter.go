package lucretia

import "fmt"

type Greeter struct {
	Forename string
	Surname  string
	Title    string
	Message  string
}

func (g *Greeter) Hello() string {
	return fmt.Sprintf("---> Hello %v. %v %v: %v",
		g.Title, g.Forename, g.Surname, g.Message,
	)
}
