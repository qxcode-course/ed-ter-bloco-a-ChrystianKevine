package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	homens := []int{}
	for _, v := range vet {
		if v > 0 {
			homens = append(homens, v)
		}
	}
	return homens
}

func getCalmWomen(vet []int) []int {
	calmas := []int{}
	for _, v := range vet {
		if v < 0 && v > -10 {
			calmas = append(calmas, v)
		}
	}
	return calmas
}

func sortVet(vet []int) []int {
	n := len(vet)
	novo := make([]int, n)
	copy(novo, vet)

	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if novo[j] > novo[j+1] {
				aux := novo[j]
				novo[j] = novo[j+1]
				novo[j+1] = aux
			}
		}
	}
	return novo
}

func sortStress(vet []int) []int {
	n := len(vet)
	novo := make([]int, n)
	copy(novo, vet)

	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			v1 := novo[j]
			if v1 < 0 {
				v1 = -v1
			}
			v2 := novo[j+1]
			if v2 < 0 {
				v2 = -v2
			}

			if v1 > v2 {
				aux := novo[j]
				novo[j] = novo[j+1]
				novo[j+1] = aux
			}
		}
	}
	return novo
}

func reverse(vet []int) []int {
	n := len(vet)
	invertido := make([]int, n)
	for i := 0; i < n; i++ {
		invertido[i] = vet[n-1-i]
	}
	return invertido
}

func unique(vet []int) []int {
	un := []int{}
	for _, v := range vet {
		existe := false
		for _, jaTem := range un {
			if v == jaTem {
				existe = true
				break
			}
		}
		if !existe {
			un = append(un, v)
		}
	}
	return un
}

func repeated(vet []int) []int {
	soUnicos := unique(vet)
	
	// Ordena os unicos (Bubble Sort simples)
	for i := 0; i < len(soUnicos); i++ {
		for j := 0; j < len(soUnicos)-1-i; j++ {
			if soUnicos[j] > soUnicos[j+1] {
				aux := soUnicos[j]
				soUnicos[j] = soUnicos[j+1]
				soUnicos[j+1] = aux
			}
		}
	}

	rep := []int{}
	for _, v := range soUnicos {
		cont := 0
		for _, original := range vet {
			if original == v {
				cont++
			}
		}
		
		if cont > 1 {
			for k := 0; k < cont-1; k++ {
				rep = append(rep, v)
			}
		}
	}
	return rep
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

