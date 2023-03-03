package lucretia

import "fmt"

type Greeter struct {
	Forename string
	Surname  string
	Title    int
}

func (g *Greeter) Hello() string {
	return fmt.Sprintf("---> Hello %v. %v %v", g.Title, g.Forename, g.Surname)
}
