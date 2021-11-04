package main

import "fmt"

func main() {
	var N, X, Y, soma int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&X, &Y)

		if X > Y && X != Y {
			Y++
			for ; X > Y; Y++ {
				if Y%2 != 0 {
					soma += Y
				}
			}
			fmt.Println(soma)
		} else if X < Y && X != Y {
			X++
			for ; X < Y; X++ {
				if X%2 != 0 {
					soma += X
				}
			}
			fmt.Println(soma)
		} else {
			soma = 0
			fmt.Println(soma)
		}
		soma = 0
	}
}
