package main

import "fmt"

func main() {
	var N int

	fmt.Scan(&N)

	for i := 1; i < 10000; i++ {
		if i%N == 2 {
			fmt.Println(i)
		}
	}
}
