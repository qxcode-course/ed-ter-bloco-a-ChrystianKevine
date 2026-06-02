package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	linhas := len(grid)
	colunas := len(grid[0])

	var dfs func(l, c, posicao int) bool
	dfs = func (l, c, posicao int) bool {
		if posicao == len(word) {
			return true
		}

		if l < 0 || l >= linhas || c < 0 || c >= colunas || grid[l][c] != word[posicao] {
			return false
		}

		temp := grid[l][c]
		grid[l][c] = '#'

		encontrou := dfs(l+1, c, posicao+1) ||
			dfs(l-1, c, posicao+1) ||
			dfs(l, c+1, posicao+1) ||
			dfs(l, c-1, posicao+1)

		grid[l][c] = temp

		return encontrou

	}

	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {
			if grid[i][j] == word[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
