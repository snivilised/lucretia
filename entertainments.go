package lucretia

import "fmt"

type Entertainments struct {
	StandUp string
	Musical string
	Poetry  string
}

func (e *Entertainments) Enact() string {
	return fmt.Sprintf("Giggle: %v, Dance: %v and Think: %v",
		e.StandUp, e.Musical, e.Poetry,
	)
}
