package main

import "fmt"

func main() {
	var N int
	var div, X, Y float32

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&X, &Y)

		if Y == 0 {
			fmt.Println("divisao impossivel")
		} else {
			div = X / Y
			fmt.Printf("%.1f\n", div)
		}

	}
}
