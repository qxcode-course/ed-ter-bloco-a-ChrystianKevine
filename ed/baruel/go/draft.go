package main

import (
	"fmt"
	"strings"
)

func main() {
	var tfalbum, tfbaruel int
	fmt.Scan(&tfalbum, &tfbaruel)

	fbaruel := make([]int, tfbaruel)
	for i := 0; i < tfbaruel; i++ {
		fmt.Scan(&fbaruel[i])
	}

	var repetidas []int
	for i := 1; i < tfbaruel; i++ {
		if fbaruel[i] == fbaruel[i-1] {
			repetidas = append(repetidas, fbaruel[i])
		}
	}
	
	if len(repetidas) == 0 {
		fmt.Println("N")
	} else {
		fmt.Println(strings.Trim(fmt.Sprint(repetidas), "[]"))
	}

	album := make([]bool, tfalbum+1)
	for i := 0; i < tfbaruel; i++ {
		album[fbaruel[i]] = true
	}

	var faltantes []int
	for i := 1; i <= tfalbum; i++ {
		if !album[i] {
			faltantes = append(faltantes, i)
		}
	}
	
	if len(faltantes) == 0 {
		fmt.Println("N")
	} else {
		fmt.Println(strings.Trim(fmt.Sprint(faltantes), "[]"))
	}
}