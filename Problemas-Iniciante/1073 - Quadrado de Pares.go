package main

import "fmt"

func main() {
	var N int

	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		if i%2 == 0 {
			quadrado := i * i
			fmt.Printf("%d^2 = %d\n", i, quadrado)
		}
	}
}
