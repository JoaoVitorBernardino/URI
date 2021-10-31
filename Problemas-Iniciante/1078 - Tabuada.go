package main

import "fmt"

func main() {
	var N, tabuada int

	fmt.Scan(&N)

	for i := 1; i <= 10; i++ {
		tabuada = i * N
		fmt.Printf("%d x %d = %d\n", i, N, tabuada)
	}
}
