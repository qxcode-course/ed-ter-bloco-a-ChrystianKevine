package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	var Q int
	var D rune
	fmt.Scanf("%d %c\n", &Q, &D)

	gomas := make([]Point, Q)
	for i := 0; i < Q; i++ {
		fmt.Scanf("%d %d\n", &gomas[i].X, &gomas[i].Y)
	}

	for i := Q - 1; i > 0; i-- {
		gomas[i] = gomas[i-1]
	}

	switch D {
	case 'L':
		gomas[0].X--
	case 'R':
		gomas[0].X++
	case 'U':
		gomas[0].Y--
	case 'D':
		gomas[0].Y++
	}

	for _, gomo := range gomas {
		fmt.Printf("%d %d\n", gomo.X, gomo.Y)
	}
}
