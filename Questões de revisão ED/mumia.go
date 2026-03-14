package main
import "fmt"
func main() {
    var pessoa string
    var idade int
    fmt.Scan(&pessoa, &idade)
    if idade < 12 {
        fmt.Println(pessoa, "eh", "crianca")
    } else if idade >= 12 && idade < 18 {
        fmt.Println(pessoa, "eh", "jovem")
    } else if idade >= 18 && idade < 65 {
        fmt.Println(pessoa, "eh", "adulto")
    } else if idade >= 65 && idade < 1000 {
        fmt.Println(pessoa, "eh", "idoso")
    } else {
        fmt.Println(pessoa, "eh", "mumia")
    }

}

