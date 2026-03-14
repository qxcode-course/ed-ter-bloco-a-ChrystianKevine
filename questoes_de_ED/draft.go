package main
import (
    "fmt"
    "math"
)
func main() {
    var a float64
    var b float64
    var c float64
    fmt.Scan(&a, &b, &c)

    p := (a + b + c) / 2
    area := math.Sqrt(p * (p - a) * (p - b) * (p - c))
    fmt.Printf("%.2f\n", area)
}
