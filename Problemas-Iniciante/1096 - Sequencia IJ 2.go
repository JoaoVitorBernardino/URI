package main

import "fmt"

func main() {
	J := 7
	for i := 1; i <= 9; {
		for j := 0; j < 3; j++ {
			fmt.Printf("I=%d J=%d\n", i, J)
			J--
		}
		J = 7
		i += 2
	}
}
