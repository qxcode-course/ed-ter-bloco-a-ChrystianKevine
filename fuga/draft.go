package main
import "fmt"
func main() {
    var H, P, F, D int
    fmt.Scan(&H, &P, &F, &D)

    if D == -1{
        fmt.Println("S")
    } else if D == 1{
        fmt.Println("N")
    }
}
