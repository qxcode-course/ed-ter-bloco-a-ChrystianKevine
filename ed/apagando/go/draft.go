package main

import "fmt"

func main() {
	var n, m, aux int

	fmt.Scan(&n)

	fila := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&fila[i])
	}

	fmt.Scan(&m)

	saiu := make(map[int]bool)
	for i := 0; i < m; i++ {
		fmt.Scan(&aux)
		saiu[aux] = true
	}

	var res []int
	for i := 0; i < n; i++ {
		if !saiu[fila[i]] {
			res = append(res, fila[i])
		}
	}

	for i := 0; i < len(res); i++ {
		fmt.Print(res[i], " ")
	}
	fmt.Print("\n")
}