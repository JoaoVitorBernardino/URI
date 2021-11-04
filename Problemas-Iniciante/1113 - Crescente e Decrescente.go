package main

import "fmt"

func main() {
	var X, Y int

	for {
		fmt.Scan(&X, &Y)

		if X > Y {
			fmt.Println("Decrescente")
		} else if X < Y {
			fmt.Println("Crescente")
		} else {
			break
		}
	}
}
