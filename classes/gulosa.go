package classes

type Gulosa struct {
	Vertice *Vertice
	Status bool
}

func NewGulosa(vertice *Vertice) *Gulosa {
	return &Gulosa{Vertice: vertice, Status: false}
}