package main

import "fmt"

func main() {
	var num, posicao int
	valores := 100
	maior := 0

	for i := 0; i < valores; i++ {
		fmt.Scan(&num)

		if num > maior {
			maior = num
			posicao = i + 1
		}
	}

	fmt.Println(maior)
	fmt.Println(posicao)
}
