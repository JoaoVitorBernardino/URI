package main

import "fmt"

func main() {
	var salario float32
	fmt.Scan(&salario)

	if salario <= 2000.00 {
		fmt.Println("Isento")
	} else if salario > 2000.00 && salario <= 3000.00 {
		salario -= 2000
		imposto08 := salario * 0.08
		fmt.Printf("R$ %.2f\n", imposto08)
	} else if salario > 3000.00 && salario <= 4500.00 {
		imposto08 := ((salario - 2000.00) - (salario - 3000.00)) * 0.08
		imposto18 := (salario - 3000.00) * 0.18
		impostoTotal := imposto08 + imposto18
		fmt.Printf("R$ %.2f\n", impostoTotal)
	} else {
		imposto28 := (salario - 4500.00) * 0.28
		imposto18 := ((salario - 3000.00) - (salario - 4500.00)) * 0.18
		imposto08 := ((salario - 3500.00) - (salario - 4500.00)) * 0.08
		impostoTotal := imposto08 + imposto18 + imposto28
		fmt.Printf("R$ %.2f\n", impostoTotal)
	}
}
