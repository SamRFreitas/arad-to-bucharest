package classes

type VetorOrdenado struct {
	Capacidade int
	UltimaPosicao int
	Vertices []*Vertice
}

func NewVetorOrdenado(capacidade int) *VetorOrdenado {
	return &VetorOrdenado{Capacidade: capacidade, UltimaPosicao: -1}
}

