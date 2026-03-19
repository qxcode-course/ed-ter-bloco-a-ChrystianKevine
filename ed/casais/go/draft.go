package main

import "fmt"

func main() {
    var tvetor int
    fmt.Scan(&tvetor)

    animais := make([]int, tvetor)
    var solteiros []int
    pares := 0 

    for i := 0; i < tvetor; i++ {
        fmt.Scan(&animais[i])
    }

    for i := 0; i < tvetor; i++ {
        if animais[i] == 0 {
            continue
        }

        achoupar := false

        for j := i + 1; j < tvetor; j++ {
            if animais[i] == -animais[j] {
                animais[i] = 0 
                animais[j] = 0 
                achoupar = true
                pares++ 
                break 
            }
        }

        if achoupar == false {
            solteiros = append(solteiros, animais[i])
        }
    }

    fmt.Println(pares)
}