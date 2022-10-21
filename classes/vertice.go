package classes
// package main

type Vertice struct {

	Rotulo string;
	Distancia int;
	Visitado bool;
	Adjacentes []Adjacente;

}

func NewVertice(rotulo string, distancia int) *Vertice {
	return &Vertice{Rotulo: rotulo, Distancia: distancia}
}

func (vertice *Vertice) AdicionaAdjacentes(adjacente Adjacente) *Vertice {
	return append(vertice.Adjacentes, adjacente)
}