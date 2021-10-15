package main

import (
	"fmt"

	"github.com/cod3rcursos/goarea"
)

type pessoa struct {
	name string
	age  int
}

type animal struct {
	name string
	age  int
}

// interfaces são implementadas implicitamente
type imprimivel interface {
	toString(string) string
}

func (p pessoa) toString(v string) string {
	return fmt.Sprintf("Corre pessoa, %v - %v", v, p.age)
}

func (a animal) toString(v string) string {
	return fmt.Sprintf("Corre animal, %v - %v", v, a.name)
}

func imprimir(x imprimivel, v string) {
	fmt.Println(x.toString(v))
}

func main() {

	p := pessoa{"Carl", 22}
	a := animal{"Pepe", 3}

	imprimir(p, "Test1")
	imprimir(a, "Test2")

	fmt.Println(goarea.Circ(3.0))
}
