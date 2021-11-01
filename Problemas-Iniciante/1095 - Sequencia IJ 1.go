package main

import "fmt"

func main() {
	J := 60
	I := 1
	for i := 1; i <= 12; i++ {
		fmt.Printf("I=%d J=%d\n", I, J)
		I += 3
		J -= 5
	}
}
