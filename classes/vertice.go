package classes

import "fmt"

// package main

type Vertice struct {
	Rotulo     string
	Distancia  int
	Visitado   bool
	Adjacentes []Adjacente
}

func NewVertice(rotulo string, distancia int) *Vertice {
	return &Vertice{Rotulo: rotulo, Distancia: distancia}
}

func Exibe(vertice *Vertice) {

	for _, adjacente := range vertice.Adjacentes {
		fmt.Println(adjacente.Vertice.Rotulo)
	}

}
