package main
import "fmt"
func main() {
    var n, e int
    fmt.Scan(&n, &e)

    pessoas := make([]int, n)
    for i := 0; i < n; i++ {
        pessoas[i] = i + 1
    }
    var espada int = e - 1

    for len(pessoas) > 1 {
        linha := "["
        for i, p := range pessoas {
            linha += " "
            if i == espada {
                linha += fmt.Sprintf("%d>", p)
            } else {
                linha += fmt.Sprintf("%d", p)
            }
        }
        linha += " ]"
        fmt.Println(linha)

        pos_vitima := (espada + 1) % len(pessoas)
        pessoas = append(pessoas[:pos_vitima], pessoas[pos_vitima+1:]...)
        espada = pos_vitima % len(pessoas)
    }

    linha := "["
    for i, p := range pessoas {
        linha += " "
        if i == espada {
            linha += fmt.Sprintf("%d>", p)
        } else {
            linha += fmt.Sprintf("%d", p)
        }
    }
    linha += " ]"
    fmt.Println(linha)
}
