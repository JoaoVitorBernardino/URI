package main

import "fmt"

func main() {
	var num, positivos float32
	for i := 0; i < 6; i++ {
		fmt.Scan(&num)
		if num > 0 {
			positivos++
		}
	}
	fmt.Printf("%.0f valores positivos\n", positivos)
}
