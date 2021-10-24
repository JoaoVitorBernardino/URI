package main

import (
	"fmt"
	"sort"
)

func main() {
	var A, B, C float64
	fmt.Scan(&A, &B, &C)

	lados := []float64{A, B, C}

	//ordem decrescente
	sort.Slice(lados, func(i, j int) bool {
		return lados[i] > lados[j]
	})

	A = lados[0]
	B = lados[1]
	C = lados[2]

	if A >= B+C {
		fmt.Println("NAO FORMA TRIANGULO")
	} else {
		if (A * A) == (B*B)+(C*C) {
			fmt.Println("TRIANGULO RETANGULO")
		}

		if (A * A) > (B*B)+(C*C) {
			fmt.Println("TRIANGULO OBTUSANGULO")
		}

		if (A * A) < (B*B)+(C*C) {
			fmt.Println("TRIANGULO ACUTANGULO")
		}

		if A == B && B == C {
			fmt.Println("TRIANGULO EQUILATERO")

		} else if A == B || A == C || B == C {
			fmt.Println("TRIANGULO ISOSCELES")

		}
	}
}
