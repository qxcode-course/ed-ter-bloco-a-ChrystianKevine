package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func resolver(bombas [][]int) int {
	n := len(bombas)
	grafo := make([][]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			dx := bombas[i][0] - bombas[j][0]
			dy := bombas[i][1] - bombas[j][1]
			r := bombas[i][2]

			if (dx*dx)+(dy*dy) <= r*r {
				grafo[i] = append(grafo[i], j)
			}
		}
	}

	maximoDetonadas := 0

	for i := 0; i < n; i++ {
		visitados := make([]bool, n)
		detonadas := 0

		var dfs func(int)
		dfs = func(atual int) {
			visitados[atual] = true
			detonadas++
			for _, vizinho := range grafo[atual] {
				if !visitados[vizinho] {
					dfs(vizinho)
				}
			}
		}

		dfs(i)

		if detonadas > maximoDetonadas {
			maximoDetonadas = detonadas
		}
		
		if maximoDetonadas == n {
			return n
		}
	}

	return maximoDetonadas
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
	n, _ := strconv.Atoi(parts[0])

	bombas := make([][]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		linhaStr := strings.Fields(scanner.Text())
		bomba := make([]int, 3)
		for j := 0; j < 3; j++ {
			bomba[j], _ = strconv.Atoi(linhaStr[j])
		}
		bombas[i] = bomba
	}

	fmt.Println(resolver(bombas))
}