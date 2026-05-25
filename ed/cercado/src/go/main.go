package main

import (
	"bufio"
	"fmt"
	"os"
)

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	nrows := len(board)
	ncols := len(board[0])

	visited := make([][]bool, nrows)
	for i := range visited {
		visited[i] = make([]bool, ncols)
	}

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= nrows || c < 0 || c >= ncols {
			return
		}
		if board[r][c] != 'O' || visited[r][c] {
			return
		}
		
		visited[r][c] = true
		
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	for i := 0; i < nrows; i++ {
		if board[i][0] == 'O' && !visited[i][0] {
			dfs(i, 0)
		}
		if board[i][ncols-1] == 'O' && !visited[i][ncols-1] {
			dfs(i, ncols-1)
		}
	}
	
	for j := 0; j < ncols; j++ {
		if board[0][j] == 'O' && !visited[0][j] {
			dfs(0, j)
		}
		if board[nrows-1][j] == 'O' && !visited[nrows-1][j] {
			dfs(nrows-1, j)
		}
	}

	for i := 0; i < nrows; i++ {
		for j := 0; j < ncols; j++ {
			if board[i][j] == 'O' && !visited[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
} 
