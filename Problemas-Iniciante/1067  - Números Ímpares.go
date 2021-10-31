package main

import "fmt"

func main() {
	var valor int

	fmt.Scan(&valor)

	for i := 1; i <= valor; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
