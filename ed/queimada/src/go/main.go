package main

import (
	"bufio"
	"fmt"
	"os"
)

func burnTrees(grid [][]rune, l, c int) {
	if l < 0 || l >= len(grid) || c < 0 || c >= len(grid[0]) {
		return
	}

	if grid[l][c] != '#' {
		return
	}
	
	grid[l][c] = 'o' 
	
	burnTrees(grid, l-1, c) // cima
	burnTrees(grid, l+1, c) // baixo
	burnTrees(grid, l, c-1) // esquerda
	burnTrees(grid, l, c+1) // direita	
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(grid [][]rune) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
