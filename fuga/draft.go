package main
import "fmt"
func main() {
    var H, P, F, D int
    fmt.Scan(&H, &P, &F, &D)

    for {
        F += D

        if F == 16{
            F = 0
        } else if F == -1 {
            F = 15
        }

        if F == H{
            fmt.Println("S")
            break
        }

        if F == P{
            fmt.Println("N")
            break
        }
    }
}        
    

