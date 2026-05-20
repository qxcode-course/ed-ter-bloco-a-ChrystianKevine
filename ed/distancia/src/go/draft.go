package main

import (
	"fmt"
)

func podeColocar(seq []rune, indice int, digito rune, limite int) bool {
	esq := indice - limite
	if esq < 0 {
		esq = 0
	}
	for i := esq; i < indice; i++ {
		if seq[i] == digito {
			return false
		}
	}

	dir := indice + limite
	if dir >= len(seq) {
		dir = len(seq) - 1
	}
	for i := indice + 1; i <= dir; i++ {
		if seq[i] == digito {
			return false
		}
	}

	return true
}

func resolver(seq []rune, indice int, limite int) bool {
	if indice == len(seq) {
		return true
	}

	if seq[indice] != '.' {
		return resolver(seq, indice+1, limite)
	}

	for d := 0; d <= limite; d++ {
		digito := rune('0' + d)
		if podeColocar(seq, indice, digito, limite) {
			seq[indice] = digito
			if resolver(seq, indice+1, limite) {
				return true
			}
			seq[indice] = '.'
		}
	}

	return false
}

func main() {
	var sequencia string
	var limite int

	_, err := fmt.Scan(&sequencia, &limite)
	if err != nil {
		return
	}

	seqRunes := []rune(sequencia)
	resolver(seqRunes, 0, limite)
	fmt.Println(string(seqRunes))
}