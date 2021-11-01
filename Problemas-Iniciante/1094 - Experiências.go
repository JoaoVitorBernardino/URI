package main

import "fmt"

func main() {
	var N, Quantia, total, totalC, totalR, totalS int
	var Tipo string

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d %s", &Quantia, &Tipo)

		total += Quantia

		if Tipo == "C" {
			totalC += Quantia
		} else if Tipo == "R" {
			totalR += Quantia
		} else if Tipo == "S" {
			totalS += Quantia
		}
	}

	var percentualCoelho, percentualRato, percentualSapo float64

	fmt.Printf("Total: %d cobaias\n", total)
	fmt.Println("Total de coelhos:", totalC)
	fmt.Println("Total de ratos:", totalR)
	fmt.Println("Total de sapos:", totalS)
	percentualCoelho = float64(totalC*100) / float64(total)
	fmt.Printf("Percentual de coelhos: %.2f %%\n", percentualCoelho)
	percentualRato = float64(totalR*100) / float64(total)
	fmt.Printf("Percentual de ratos: %.2f %%\n", percentualRato)
	percentualSapo = float64(totalS*100) / float64(total)
	fmt.Printf("Percentual de sapos: %.2f %%\n", percentualSapo)
}
