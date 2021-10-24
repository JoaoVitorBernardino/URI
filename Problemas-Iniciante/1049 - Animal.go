package main

import "fmt"

func main() {
	var palavra1, palavra2, palavra3 string
	fmt.Scan(&palavra1, &palavra2, &palavra3)

	if palavra1 == "vertebrado" {
		if palavra2 == "ave" {
			if palavra3 == "carnivoro" {
				fmt.Println("aguia")
			} else {
				fmt.Println("pomba")
			}
		} else {
			if palavra3 == "onivoro" {
				fmt.Println("homem")
			} else {
				fmt.Println("vaca")
			}
		}
	} else {
		if palavra2 == "inseto" {
			if palavra3 == "hematofago" {
				fmt.Println("pulga")
			} else {
				fmt.Println("lagarta")
			}
		} else {
			if palavra3 == "hematofago" {
				fmt.Println("sanguessuga")
			} else {
				fmt.Println("minhoca")
			}
		}
	}
}
