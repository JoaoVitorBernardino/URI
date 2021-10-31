package main

import "fmt"

func main() {
	var valor int
	var par, impar, positivo, negativo int
	for i := 0; i < 5; i++ {
		fmt.Scan(&valor)

		if valor%2 == 0 {
			par++
		} else {
			impar++
		}

		if valor > 0 {
			positivo++
		} else if valor < 0 {
			negativo++
		}
	}

	fmt.Printf("%d valor(es) par(es)\n", par)
	fmt.Printf("%d valor(es) impar(es)\n", impar)
	fmt.Printf("%d valor(es) positivo(s)\n", positivo)
	fmt.Printf("%d valor(es) negativo(s)\n", negativo)

}
