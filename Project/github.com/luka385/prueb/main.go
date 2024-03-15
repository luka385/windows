package main

import "fmt"

type Circulo struct {
	n1 int
	n2 int
}

func (c *Circulo) suma() int {
	return c.n1 + c.n2
}

func (c *Circulo) resta() int {
	return c.n1 - c.n2
}

func main() {
	circulo1 := Circulo{n1: 20, n2: 50}
	circulo2 := Circulo{n1: 30, n2: 50}
	fmt.Println(circulo1.suma())
	fmt.Println(circulo2.resta())

}
