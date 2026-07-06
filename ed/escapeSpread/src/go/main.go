package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func resolver(grid [][]int) int {
	linhas := len(grid)
	colunas := len(grid[0])
	infinito := 2000000000

	tempoFogo := make([][]int, linhas)
	var filaFogo [][]int

	for i := 0; i < linhas; i++ {
		tempoFogo[i] = make([]int, colunas)
		for j := 0; j < colunas; j++ {
			tempoFogo[i][j] = infinito
			if grid[i][j] == 1 {
				tempoFogo[i][j] = 0
				filaFogo = append(filaFogo, []int{i, j})
			}
		}
	}

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(filaFogo) > 0 {
		atual := filaFogo[0]
		filaFogo = filaFogo[1:]
		l, c := atual[0], atual[1]

		for _, d := range dirs {
			nl, nc := l+d[0], c+d[1]
			if nl >= 0 && nl < linhas && nc >= 0 && nc < colunas && grid[nl][nc] == 0 {
				if tempoFogo[l][c]+1 < tempoFogo[nl][nc] {
					tempoFogo[nl][nc] = tempoFogo[l][c] + 1
					filaFogo = append(filaFogo, []int{nl, nc})
				}
			}
		}
	}

	consegueEscapar := func(espera int) bool {
		if tempoFogo[0][0] <= espera {
			return false
		}

		visitados := make([][]bool, linhas)
		for i := range visitados {
			visitados[i] = make([]bool, colunas)
		}
		visitados[0][0] = true

		var filaPessoa [][]int
		filaPessoa = append(filaPessoa, []int{0, 0, espera})

		for len(filaPessoa) > 0 {
			atual := filaPessoa[0]
			filaPessoa = filaPessoa[1:]
			l, c, tempo := atual[0], atual[1], atual[2]

			if l == linhas-1 && c == colunas-1 {
				return true
			}

			for _, d := range dirs {
				nl, nc := l+d[0], c+d[1]
				if nl >= 0 && nl < linhas && nc >= 0 && nc < colunas && grid[nl][nc] == 0 && !visitados[nl][nc] {
					tempoChegada := tempo + 1
					chegouJunto := (nl == linhas-1 && nc == colunas-1 && tempoChegada == tempoFogo[nl][nc])

					if tempoChegada < tempoFogo[nl][nc] || chegouJunto {
						visitados[nl][nc] = true
						filaPessoa = append(filaPessoa, []int{nl, nc, tempoChegada})
					}
				}
			}
		}
		return false
	}

	esq, dir := 0, 1000000000
	melhorEspera := -1

	for esq <= dir {
		meio := esq + (dir-esq)/2
		if consegueEscapar(meio) {
			melhorEspera = meio
			esq = meio + 1
		} else {
			dir = meio - 1
		}
	}

	return melhorEspera
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
	linhas, _ := strconv.Atoi(parts[0])
	colunas, _ := strconv.Atoi(parts[1])

	grid := make([][]int, linhas)
	for i := 0; i < linhas; i++ {
		scanner.Scan()
		linhaStr := strings.Fields(scanner.Text())
		linhaInt := make([]int, colunas)
		for j := 0; j < colunas; j++ {
			linhaInt[j], _ = strconv.Atoi(linhaStr[j])
		}
		grid[i] = linhaInt
	}

	fmt.Println(resolver(grid))
}