package main

import "fmt"

func main() {
	var M, N, soma int
	for {
		fmt.Scan(&M, &N)

		if M <= 0 || N <= 0 {
			break
		}

		if M > N {
			fmt.Printf("%d ", N)
			soma = N
			for N += 1; N <= M; N++ {
				fmt.Printf("%d ", N)
				soma += N
			}
			fmt.Printf("Sum=%d\n", soma)
		} else {
			fmt.Printf("%d ", M)
			soma = M
			for M += 1; M <= N; M++ {
				fmt.Printf("%d ", M)
				soma += M
			}
			fmt.Printf("Sum=%d\n", soma)
		}

	}
}
