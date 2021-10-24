package main

import "fmt"

func main() {
	var salario, novoSalario, ajuste float32
	fmt.Scan(&salario)

	if salario <= 400.00 {
		ajuste = salario * 0.15
		novoSalario = salario + ajuste
		fmt.Printf("Novo salario: %.2f\n", novoSalario)
		fmt.Printf("Reajuste ganho: %.2f\n", ajuste)
		fmt.Println("Em percentual: 15 %")
	} else if salario > 400.00 && salario <= 800.00 {
		ajuste = salario * 0.12
		novoSalario = salario + ajuste
		fmt.Printf("Novo salario: %.2f\n", novoSalario)
		fmt.Printf("Reajuste ganho: %.2f\n", ajuste)
		fmt.Println("Em percentual: 12 %")
	} else if salario > 800.00 && salario <= 1200.00 {
		ajuste = salario * 0.10
		novoSalario = salario + ajuste
		fmt.Printf("Novo salario: %.2f\n", novoSalario)
		fmt.Printf("Reajuste ganho: %.2f\n", ajuste)
		fmt.Println("Em percentual: 10 %")
	} else if salario > 1200.00 && salario <= 2000.00 {
		ajuste = salario * 0.07
		novoSalario = salario + ajuste
		fmt.Printf("Novo salario: %.2f\n", novoSalario)
		fmt.Printf("Reajuste ganho: %.2f\n", ajuste)
		fmt.Println("Em percentual: 7 %")
	} else {
		ajuste = salario * 0.04
		novoSalario = salario + ajuste
		fmt.Printf("Novo salario: %.2f\n", novoSalario)
		fmt.Printf("Reajuste ganho: %.2f\n", ajuste)
		fmt.Println("Em percentual: 4 %")
	}
}
