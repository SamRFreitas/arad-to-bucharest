package classes

type Adjacente struct {

	Vertice *Vertice;
	Custo int
	Adjacentes []*Adjacente;

}

func NewAdjacente(vertice *Vertice, custo int,) *Adjacente {
	return &Adjacente{Vertice: vertice, Custo: custo}
}