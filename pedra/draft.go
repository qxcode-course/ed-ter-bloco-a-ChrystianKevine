package main
import "fmt"
func main() {
   var N int 
    fmt.Scan(&N)
    var melhorPontuacao int = 101   
    var vencedor int = -1

    for i := 0; i < N; i++ {
        var a, b int
        fmt.Scan(&a, &b)

        if a >= 10 && b >= 10 {
            pontuacao := a - b
            if pontuacao < 0 {
                pontuacao = -pontuacao
            }   

            if vencedor == -1 || pontuacao < melhorPontuacao {
                melhorPontuacao = pontuacao
                vencedor = i
            }
            
        } 

    }

    if vencedor == -1 {
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(vencedor)
    }
    // mudando o commit para testar o envio
}
