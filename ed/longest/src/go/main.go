package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	linhas := len(matrix)
	colunas := len(matrix[0])
	maiorCaminhoGeral := 0

	cache := make([][]int, linhas)
	for i := range cache {
		cache[i] = make([]int, colunas)
	}

	dl := []int{-1, 1, 0, 0}
	dc := []int{0, 0, -1, 1}

	var explorar func(l, c int) int
	explorar = func(l, c int) int {
		if cache[l][c] != 0 {
			return cache[l][c]
		}

		maximo := 1

		for i := 0; i < 4; i++ {
			nl := l + dl[i]
			nc := c + dc[i]

			if nl >= 0 && nl < linhas && nc >= 0 && nc < colunas && matrix[nl][nc] > matrix[l][c] {
				caminho := 1 + explorar(nl, nc)
				if caminho > maximo {
					maximo = caminho
				}
			}
		}

		cache[l][c] = maximo
		return maximo
	}

	for l := 0; l < linhas; l++ {
		for c := 0; c < colunas; c++ {
			caminhoAtual := explorar(l, c)
			if caminhoAtual > maiorCaminhoGeral {
				maiorCaminhoGeral = caminhoAtual
			}
		}
	}

	return maiorCaminhoGeral
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return
	}
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	matrix := make([][]int, nl)
	for i := 0; i < nl; i++ {
		if !scanner.Scan() {
			return
		}
		tokens := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc && j < len(tokens); j++ {
			v, _ := strconv.Atoi(tokens[j])
			row[j] = v
		}
		matrix[i] = row
	}

	fmt.Println(longestIncreasingPath(matrix))
}