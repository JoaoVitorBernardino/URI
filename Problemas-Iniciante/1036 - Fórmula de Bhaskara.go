package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B, C float64
	fmt.Scanf("%f %f %f", &A, &B, &C)

	delta := math.Pow(B, 2) - (4 * A * C)
	raiz := math.Sqrt(delta)

	R1 := (-B + raiz) / (2 * A)
	R2 := (-B - raiz) / (2 * A)

	if R1 > 0 && R2 > 0 || R1 < 0 && R2 < 0 {
		fmt.Printf("R1 = %.5f\n", R1)
		fmt.Printf("R2 = %.5f\n", R2)
	} else {
		fmt.Println("Impossivel calcular")
	}
}
