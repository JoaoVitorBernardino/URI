package main

import "fmt"

func main() {
	n := 7
	x := 0
	for i := 1; i <= 9; {
		x = n
		for j := 1; j <= 3; j++ {
			fmt.Printf("I=%d J=%d\n", i, n)
			n--
		}
		i += 2
		n = x + 2
	}
}
