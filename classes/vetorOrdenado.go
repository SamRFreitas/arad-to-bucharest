package classes

type VetorOrdenado struct {
	Capacidade int
	UltimaPosicao int
	Vertices []*Vertice
}

func NewVetorOrdenado(capacidade int, ultimaPosicao int) *VetorOrdenado {
	return &VetorOrdenado{Capacidade: capacidade, UltimaPosicao: ultimaPosicao}
}

