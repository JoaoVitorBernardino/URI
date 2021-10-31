package main

import "fmt"

func main() {
	var N, X int
	var in, out int

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&X)

		if X >= 10 && X <= 20 {
			in++
		} else {
			out++
		}
	}

	fmt.Println(in, "in")
	fmt.Println(out, "out")
}
