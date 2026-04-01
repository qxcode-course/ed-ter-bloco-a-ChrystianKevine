package main
import "fmt"
func resto (num int){
    if num == 0 {
        return
    }
    resto(num/2)
    fmt.Printf("%d %d\n", num / 2, num%2)
}

func main() {
    var a int 
    fmt.Scan(&a)
    num := a
    resto(num)
}
