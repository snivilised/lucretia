package lucretia

import "fmt"

type Refreshment struct {
	Aperitif string
	Vintage  string
	Chaser   string
}

func (r *Refreshment) Serve() string {
	return fmt.Sprintf("Aperitif: %v, Vintage: %v and Chaser: %v",
		r.Aperitif, r.Vintage, r.Chaser,
	)
}
