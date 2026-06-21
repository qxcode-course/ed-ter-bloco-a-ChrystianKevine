package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordenada struct {
	linha, coluna int
}
 
func resolver(matriz [][]int) int {
	if len(matriz) == 0 || len(matriz[0]) == 0 {
		return 0
	}

	linhas := len(matriz)
	colunas := len(matriz[0])

	var fila []Coordenada
	frescas := 0

	for l := 0; l < linhas; l++ {
		for c := 0; c < colunas; c++ {
			if matriz[l][c] == 2 {
				fila = append(fila, Coordenada{linha: l, coluna: c})
			} else if matriz[l][c] == 1 {
				frescas++
			}
		}
	}

	if frescas == 0 {
		return 0
	}

	minutos := 0
	direcoes := []Coordenada{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	for len(fila) > 0 && frescas > 0 {
		minutos++
		tamanhoNivel := len(fila)

		for i := 0; i < tamanhoNivel; i++ {
			atual := fila[0]
			fila = fila[1:]

			for _, d := range direcoes {
				nl := atual.linha + d.linha
				nc := atual.coluna + d.coluna

				if nl >= 0 && nl < linhas && nc >= 0 && nc < colunas && matriz[nl][nc] == 1 {
					matriz[nl][nc] = 2
					frescas--
					fila = append(fila, Coordenada{linha: nl, coluna: nc})
				}
			}
		}
	}

	if frescas > 0 {
		return -1
	}

	return minutos
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}

	partes := strings.Fields(scanner.Text())
	nl, _ := strconv.Atoi(partes[0])
	nc, _ := strconv.Atoi(partes[1])

	matriz := make([][]int, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		linhaStr := strings.Fields(scanner.Text())
		linha := make([]int, nc)
		for j := 0; j < nc; j++ {
			linha[j], _ = strconv.Atoi(linhaStr[j])
		}
		matriz[i] = linha
	}

	fmt.Println(resolver(matriz))
}