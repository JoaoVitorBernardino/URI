package main

import "fmt"

func main() {
	var horaInicial, horaFinal, duracaoTotal int
	fmt.Scan(&horaInicial, &horaFinal)

	if horaInicial == horaFinal {
		fmt.Printf("O JOGO DUROU 24 HORA(S)")
	} else if horaInicial > horaFinal {
		duracaoTotal = 24 - horaInicial + horaFinal
		fmt.Printf("O JOGO DUROU %d HORA(S)", duracaoTotal)
	} else {
		duracaoTotal = horaFinal - horaInicial
		fmt.Printf("O JOGO DUROU %d HORA(S)", duracaoTotal)
	}
}
