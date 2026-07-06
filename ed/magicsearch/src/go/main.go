package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MagicSearch(slice []int, value int) int {
	esq, dir := 0, len(slice)-1

	for esq <= dir {
		meio := esq + (dir-esq)/2

		if slice[meio] == value {
			for meio < len(slice)-1 && slice[meio+1] == value {
				meio++
			}
			return meio
		} else if slice[meio] < value {
			esq = meio + 1
		} else {
			dir = meio - 1
		}
	}

	return esq 
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	
	slice := make([]int, 0, len(parts)-2)
	for _, elem := range parts[1 : len(parts)-1] {
		val, _ := strconv.Atoi(elem)
		slice = append(slice, val)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	
	fmt.Println(MagicSearch(slice, value))
}