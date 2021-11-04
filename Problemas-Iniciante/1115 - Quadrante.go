package main

import "fmt"

func main() {
	var X, Y int

	for {
		fmt.Scan(&X, &Y)

		if X > 0 && Y > 0 {
			fmt.Println("primeiro")
		} else if X < 0 && Y > 0 {
			fmt.Println("segundo")
		} else if X < 0 && Y < 0 {
			fmt.Println("terceiro")
		} else if X > 0 && Y < 0 {
			fmt.Println("quarto")
		} else if X == 0 || Y == 0 {
			break
		}
	}
}
