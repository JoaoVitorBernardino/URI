package main

import (
	"fmt"
	"sort"
)

func main() {
	var v1, v2, v3 int
	fmt.Scan(&v1, &v2, &v3)

	valores := []int{v1, v2, v3}

	sort.Ints(valores)
	for _, i := range valores {
		fmt.Println(i)
		i++
	}

	fmt.Println()
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
}
