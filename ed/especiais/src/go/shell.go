package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

type Pair struct {
	One int
	Two int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func occurr(vet []int) []Pair {
	contagens := make(map[int]int)

	for _, num := range vet {
		contagens [abs(num)]++
	}

	var chaves []int
	for chave := range contagens {
		chaves = append(chaves, chave)
	}

	sort.Ints(chaves)

	var resultado []Pair
	for _, chave := range chaves {
		resultado = append(resultado, Pair{One: chave, Two: contagens[chave]})
	}

	if resultado == nil {
		return []Pair{}
	}
	return resultado 
}	

func teams(vet []int) []Pair {
	if len(vet) == 0 {
		return []Pair {}
	}

	var resultado []Pair
	stressAtual := abs(vet[0])
	quantidade := 1

	for i := 1; i < len(vet); i++ {
		stress := abs(vet[i])

		if stress == stressAtual {
			quantidade++
		} else {
			resultado = append(resultado, Pair{One: stressAtual, Two: quantidade})
			stressAtual = stress
			quantidade = 1
		} 

	}

	resultado = append(resultado, Pair{One: stressAtual, Two: quantidade})

	return resultado
}

func mnext(vet []int) []int {
	resultado := make ([]int, len(vet))

	for i, v := range vet {
		if v > 0 {
			temMulherLado := false

			if i > 0 && vet[i-1] < 0 {
				temMulherLado = true
			}

			if i < len(vet)-1 && vet[i+1] < 0 {
				temMulherLado = true
			}

			if temMulherLado {
				resultado[i] = 1
			}
		}
	}
	return resultado
}

func alone(vet []int) []int {
	resultado := make([]int, len(vet))

	for i, v := range vet {
		if v > 0 {
			temMulherLado := false

			if i > 0 && vet[i-1] < 0 {
				temMulherLado = true
			}

			if i < len(vet)-1 && vet[i+1] < 0 {
				temMulherLado = true
			}

			if !temMulherLado { // Se NÃO tem mulher ao lado
				resultado[i] = 1
			}
		}
	}
	return resultado
}

func couple(vet []int) int {
	homens := make(map[int]int)
	mulheres := make(map[int]int)

	for _, v := range vet {
		if v > 0 {
			homens[v]++
		} else if v < 0 {
			mulheres[-v]++ 
		}
	}

	casais := 0

	for stress, qtdHomens := range homens {
		qtdMulheres := mulheres[stress]
		
		if qtdHomens < qtdMulheres {
			casais += qtdHomens
		} else {
			casais += qtdMulheres
		}
	}
	
	return casais
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	if pos+len(seq) > len(vet) {
		return false
	}
	
	for i := 0; i < len(seq); i++ {
		if vet[pos+i] != seq[i] {
			return false
		}
	} 
	
	return true
}

func subseq(vet []int, seq []int) int {
	if len(seq) == 0 {
		return 0
	}
	
	limite := len(vet) - len(seq)
	for i := 0; i <= limite; i++ {
		if hasSubseq(vet, seq, i) {
			return i 
		}
	}
	
	return -1
}

func erase(vet []int, posList []int) []int {
	// Cria um "caderninho" (mapa) anotando quais índices devem sumir
	remover := make(map[int]bool)
	for _, pos := range posList {
		remover[pos] = true
	}

	var resultado []int
	for i, v := range vet {
		if !remover[i] {
			resultado = append(resultado, v)
		}
	}
	
	if resultado == nil {
		return []int{}
	}
	return resultado
}

func clear(vet []int, value int) []int {
	var resultado []int
	for _, v := range vet {
		if v != value {
			resultado = append(resultado, v)
		}
	}

	if resultado == nil {
		return []int{}
	}
	return resultado
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
