package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}
	var tostr_rec func([]int) string
	tostr_rec = func(v []int) string {
		if len(v) == 1 {
			return strconv.Itoa(v[0])
		}
		return strconv.Itoa(v[0]) + ", " + tostr_rec(v[1:])
	}
	return "[" + tostr_rec(vet) + "]"
}

func tostrrev(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}
	var tostrrev_rec func([]int) string
	tostrrev_rec = func(v []int) string {
		if len(v) == 1 {
			return strconv.Itoa(v[0])
		}
	
		return tostrrev_rec(v[1:]) + ", " + strconv.Itoa(v[0])
	}
	return "[" + tostrrev_rec(vet) + "]"
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	var reverse_rec func(i, j int)
	reverse_rec = func(i, j int) {
		if i >= j {
			return
		}
		vet[i], vet[j] = vet[j], vet[i]
		reverse_rec(i+1, j-1)
	}
	reverse_rec(0, len(vet)-1)
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	if len(vet) == 0 {
		return 0
	}
	return vet[0] + sum(vet[1:])
}


// mult: produto dos elementos do slice
func mult(vet []int) int {
	if len(vet) == 0 {
		return 1
	}
	return vet[0] * mult(vet[1:])
}

// min: retorna o índice do menor valor
func min(vet []int) int {
	if len(vet) == 0 {
		return -1
	}
	// a função min_rec retorna (valor, indice)
	var min_rec func(v []int) (int, int)
	min_rec = func(v []int) (int, int) {
		if len(v) == 1 {
			return v[0], 0
		}
		
		min_val, min_idx := min_rec(v[1:])

		if v[0] <= min_val {
			return v[0], 0
		}
		
		return min_val, min_idx + 1
	}
	_, indice := min_rec(vet)
	return indice
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
