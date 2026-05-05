package main

import (
	"fmt"
)

func main() {
	var palavra string
	fmt.Scan(&palavra)

	var esquerda []rune
	var direita []rune

	for _, char := range palavra {
		switch char {
		case 'R':
			esquerda = append(esquerda, '\n')
		case '<':
			if len(esquerda) > 0 {
				direita = append([]rune{esquerda[len(esquerda)-1]}, direita...)
				esquerda = esquerda[:len(esquerda)-1]
			}
		case '>':
			if len(direita) > 0 {
				esquerda = append(esquerda, direita[0])
				direita = direita[1:]
			}
		case 'B':
			if len(esquerda) > 0 {
				esquerda = esquerda[:len(esquerda)-1] 
			}
		case 'D':
			if len(direita) > 0 {
				direita = direita[1:] 
			}
		default:
			if (char >= 'a' && char <= 'z') || char == '-' {
				esquerda = append(esquerda, char)
			}
		}
	}

	fmt.Printf("%s|%s\n", string(esquerda), string(direita))
}