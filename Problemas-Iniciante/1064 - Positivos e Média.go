package main

import "fmt"

func main() {
	var valor float64
	var i int
	positivo := 0
	soma := 0.0

	for i = 0; i < 6; i++ {
		fmt.Scan(&valor)

		if valor > 0 {
			positivo++
			soma += valor
		}
	}

	media := soma / float64(positivo)

	fmt.Println(positivo, "valores positivos")
	fmt.Printf("%.1f\n", media)
}
