package main

import "fmt"

func main() {
	var A, B, C float32
	fmt.Scan(&A, &B, &C)

	if (A+B) > C && (A+C) > B && (C+B) > A {
		perimetro := A + B + C
		fmt.Printf("Perimetro = %.1f\n", perimetro)
	} else {
		area := (A + B) * C / 2
		fmt.Printf("Area = %.1f\n", area)
	}

}
