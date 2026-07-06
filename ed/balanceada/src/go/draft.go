package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	texto := strings.TrimSpace(scanner.Text())

	pares := map[rune]rune{
		')': '(',
		']': '[',
	}

	var pilha []rune

	for _, char := range texto {
		if char == '(' || char == '[' {
			pilha = append(pilha, char)
		} else if abertura, ehFechamento := pares[char]; ehFechamento {
			if len(pilha) == 0 || pilha[len(pilha)-1] != abertura {
				fmt.Println("nao balanceado")
				return
			}
			pilha = pilha[:len(pilha)-1]
		}
	}

	if len(pilha) == 0 {
		fmt.Println("balanceado")
	} else {
		fmt.Println("nao balanceado")
	}
}