package main

import (
	"fmt"
	"arad-to-bucharest/classes"
)

func main ()  {

	v := classes.NewVertice("Ita", 100)

	fmt.Println(v.Name)

	a := classes.NewAdjacente(v, 100)

	fmt.Println(a.Vertice)
}