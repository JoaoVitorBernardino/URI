package main

import "fmt"

func main() {
	var x, y int

	fmt.Scan(&x)
	fmt.Scan(&y)

	positivo := 0
	negativo := 0

	if x > 0 {
		for i := y; i < x; i++ {
			if i%2 != 0 && i != y {
				if i > 0 {
					positivo += i
				} else {
					negativo -= i
				}
			}
		}
	}

	soma := positivo - negativo

	fmt.Println(soma)
}
