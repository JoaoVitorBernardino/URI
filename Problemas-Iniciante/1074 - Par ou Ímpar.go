package main

import "fmt"

func main() {
	var N, X int

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&X)

		if X > 0 {
			if X%2 == 0 {
				fmt.Println("EVEN POSITIVE")
			} else {
				fmt.Println("ODD POSITIVE")
			}
		} else if X < 0 {
			if X%2 == 0 {
				fmt.Println("EVEN NEGATIVE")
			} else {
				fmt.Println("ODD NEGATIVE")
			}
		} else {
			fmt.Println("NULL")
		}
	}
}
