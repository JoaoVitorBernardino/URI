package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Scan(&n1, &n2)

	if n1%n2 == 0 || n2%n1 == 0 {
		fmt.Println("Sao Multiplos")
	} else {
		fmt.Println("Nao sao Multiplos")
	}
}
