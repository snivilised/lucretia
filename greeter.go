package lucretia

import "fmt"

type Greeter struct {
	Name  string
	Title int
}

func (g *Greeter) Hello() string {
	return fmt.Sprintf("---> Hello %v %v", g.Title, g.Name)
}
