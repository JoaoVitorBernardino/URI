package main

import "fmt"

func main() {
	var N int
	var nota1, nota2, nota3 float32

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&nota1, &nota2, &nota3)

		media := ((nota1 * 2) + (nota2 * 3) + (nota3 * 5)) / 10

		fmt.Printf("%.1f\n", media)
	}

}
