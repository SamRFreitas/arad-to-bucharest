package main

import (
	"arad-to-bucharest/classes"
	"fmt"
)

type Grafo struct {
	Arad, Zerind, Oradea, Sibiu, Timisoara, Lugoj, Mehadia, Dobreta, Craiova, Rimnicu, Fagaras, Pitest, Bucharest, Giurgiu classes.Vertice
}

func NewGrafo(
	arad classes.Vertice,
	zerind classes.Vertice,
	oradea classes.Vertice,
	sibiu classes.Vertice,
	timisoara classes.Vertice,
	lugoj classes.Vertice,
	mehadia classes.Vertice,
	dobreta classes.Vertice,
	craiova classes.Vertice,
	rimnicu classes.Vertice,
	fagaras classes.Vertice,
	pitest classes.Vertice,
	bucharest classes.Vertice,
	giurgiu classes.Vertice,
) *Grafo {

	grafo := &Grafo{
		Arad:      arad,
		Zerind:    zerind,
		Oradea:    oradea,
		Sibiu:     sibiu,
		Timisoara: timisoara,
		Lugoj:     lugoj,
		Mehadia:   mehadia,
		Dobreta:   dobreta,
		Craiova:   craiova,
		Rimnicu:   rimnicu,
		Fagaras:   fagaras,
		Pitest:    pitest,
		Bucharest: bucharest,
		Giurgiu:   giurgiu,
	}

	var adjacentes []classes.Adjacente
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Zerind, 75))
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Sibiu, 140))
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Timisoara, 118))

	grafo.Arad.Adjacentes = adjacentes

	adjacentes = nil
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Arad, 75))
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Oradea, 71))

	grafo.Zerind.Adjacentes = adjacentes

	adjacentes = nil
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Oradea, 71))
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Sibiu, 151))

	grafo.Oradea.Adjacentes = adjacentes

	adjacentes = nil
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Oradea, 151))
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Arad, 140))
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Fagaras, 99))
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Rimnicu, 80))

	grafo.Sibiu.Adjacentes = adjacentes

	adjacentes = nil
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Arad, 118))
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Lugoj, 111))

	grafo.Timisoara.Adjacentes = adjacentes

	adjacentes = nil
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Timisoara, 111))
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Mehadia, 70))

	grafo.Lugoj.Adjacentes = adjacentes

	adjacentes = nil
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Lugoj, 70))
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Dobreta, 75))

	grafo.Mehadia.Adjacentes = adjacentes

	adjacentes = nil
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Mehadia, 75))
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Craiova, 120))

	grafo.Dobreta.Adjacentes = adjacentes

	adjacentes = nil
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Dobreta, 120))
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Rimnicu, 146))
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Pitest, 138))

	grafo.Craiova.Adjacentes = adjacentes

	adjacentes = nil
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Craiova, 138))
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Bucharest, 101))
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Rimnicu, 97))

	grafo.Pitest.Adjacentes = adjacentes

	adjacentes = nil
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Pitest, 97))
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Sibiu, 80))
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Craiova, 146))

	grafo.Rimnicu.Adjacentes = adjacentes

	adjacentes = nil
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Sibiu, 99))
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Bucharest, 211))

	grafo.Fagaras.Adjacentes = adjacentes

	adjacentes = nil
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Fagaras, 211))
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Giurgiu, 90))
	adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Pitest, 101))

	grafo.Bucharest.Adjacentes = adjacentes

	adjacentes = nil
	adjacentes = append(giurgiu.Adjacentes, *classes.NewAdjacente(&grafo.Bucharest, 90))

	grafo.Giurgiu.Adjacentes = adjacentes

	return grafo
}

func main() {
	arad := classes.NewVertice("Arad", 366)
	zerind := classes.NewVertice("Zerind", 374)
	oradea := classes.NewVertice("Oradea", 380)
	sibiu := classes.NewVertice("Sibiu", 253)
	timisoara := classes.NewVertice("Timisoara", 329)
	lugoj := classes.NewVertice("Lugoj", 244)
	mehadia := classes.NewVertice("Mehadia", 241)
	dobreta := classes.NewVertice("Dobreta", 242)
	craiova := classes.NewVertice("Craiova", 160)
	rimnicu := classes.NewVertice("Rimnicu", 193)
	fagaras := classes.NewVertice("Fagaras", 178)
	pitest := classes.NewVertice("Pitest", 98)
	bucharest := classes.NewVertice("Bucharest", 0)
	giurgiu := classes.NewVertice("Giurgiu", 77)

	grafo := NewGrafo(*arad, *zerind, *oradea, *sibiu, *timisoara, *lugoj, *mehadia, *dobreta, *craiova, *rimnicu, *fagaras, *pitest, *bucharest, *giurgiu)

	// fmt.Println(grafo.Arad.Rotulo)

	
	// gulosa := classes.NewGulosa(&grafo.Arad, &grago.bucharest)


	// fmt.Println(gulosa.Destino.Rotulo)

	// busca(&grafo.Arad, &grafo.Bucharest)

	// adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Zerind, 75))
	// adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Sibiu, 140))
	// adjacentes = append(adjacentes, *classes.NewAdjacente(&grafo.Timisoara, 118))

	// a := buscaIndiceDeInsercao(grafo.Arad.Adjacentes, 2)

	
	var teste []classes.Adjacente
	teste = append(teste, *classes.NewAdjacente(&grafo.Fagaras, 99))
	teste = append(teste, *classes.NewAdjacente(&grafo.Rimnicu, 80))
	teste = append(teste, *classes.NewAdjacente(&grafo.Arad, 140))
	teste = append(teste, *classes.NewAdjacente(&grafo.Oradea, 151))

	exibeVetor(teste)

	fmt.Println("CHAMA")
	
	a := inserir(teste, *mehadia)

	fmt.Println(a)
}

func busca(origem *classes.Vertice, destino *classes.Vertice) {

	status := false


	if (origem.Rotulo == destino.Rotulo) {
		status = true
	} else {
		capacidade := len(origem.Adjacentes)
		fmt.Println(capacidade)
		ultimaPosicao := len(origem.Adjacentes) - 1
		fmt.Println(ultimaPosicao)
		vetorOrdenado := classes.NewVetorOrdenado(capacidade, ultimaPosicao)
		fmt.Println(vetorOrdenado)
	}
	
	fmt.Println(status)
	
}
// func iniciarVetor(vetor []int) []int{

// 	for index, _ := range vetor {
// 		vetor[index] = -1
// 	}

// }

func percorrerVetor(listaAdjacentes []classes.Vertice) []classes.Vertice{
	for _, adjacente := range listaAdjacentes {
		if(!adjacente.Visitado) {
			adjacente.Visitado = true
			// inserir()
		}
	}
	
	return listaAdjacentes
}
func exibeVetor(vetorOrdenado []classes.Adjacente) {
	for _, adjacente := range vetorOrdenado {
		fmt.Println(adjacente.Vertice.Rotulo, adjacente.Vertice.Distancia )
	}
}


func buscaIndiceDeInsercao(vetorOrdenado []classes.Adjacente, vertice classes.Vertice) int {

	for index, elemento := range vetorOrdenado {
		if vertice.Distancia < elemento.Vertice.Distancia || elemento.Custo == -1 {

			return index
		}
	}

	return -1
}

func inserir(vetorOrdenado []classes.Adjacente, numero classes.Vertice) []classes.Adjacente {

	var indiceDeInsercao int = buscaIndiceDeInsercao(vetorOrdenado, numero)

	var i int = len(vetorOrdenado) - 1

	for i > 0 {

		if i >= indiceDeInsercao {
			// fmt.Println(vetorOrdenado[i])
			vetorOrdenado[i] = vetorOrdenado[i-1]
		}

		i--
	}

	vetorOrdenado[indiceDeInsercao].Vertice = &numero

	exibeVetor(vetorOrdenado)

	return vetorOrdenado
}

