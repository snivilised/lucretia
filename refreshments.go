package lucretia

import (
	"fmt"
	"math/rand"
)

type Refreshment struct {
	Aperitif string
	Vintage  string
	Chaser   string
}

func (r *Refreshment) Serve() string {
	min := 2
	max := 10
	treats := rand.Intn(max-min) + min
	return fmt.Sprintf("Aperitif: %v, Vintage: %v, Chaser: %v and :%v surprise treats",
		r.Aperitif, r.Vintage, r.Chaser, treats,
	)
}
