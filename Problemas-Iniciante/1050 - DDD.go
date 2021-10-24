package main

import "fmt"

func main() {
	var DDD int
	fmt.Scan(&DDD)

	switch DDD {
	case 61:
		fmt.Println("Brasilia")
		break
	case 71:
		fmt.Println("Salvador")
		break
	case 11:
		fmt.Println("Sao Paulo")
		break
	case 21:
		fmt.Println("Rio de Janeiro")
		break
	case 32:
		fmt.Println("Juiz de Fora")
		break
	case 19:
		fmt.Println("Campinas")
		break
	case 27:
		fmt.Println("Vitoria")
		break
	case 31:
		fmt.Println("Belo Horizonte")
		break
	default:
		fmt.Println("DDD nao cadastrado")
		break
	}
}
