package main

import "fmt"

var x int = 42
var y string = "James bond"
var z bool = true

func main() {
	// Usamos %d, %s e %t para imprimir os valores das variáveis x, y e z
	s := fmt.Sprintf("%d %s %t", x, y, z)
	fmt.Println(s)
}
