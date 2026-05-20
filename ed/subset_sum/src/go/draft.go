package main
import "fmt"

func somaSubConjuntos(numeros[] int, alvo int) bool {
    if alvo == 0 {
        return true
    } 

    if len(numeros) == 0 || alvo < 0 {
        return false
    }

    if somaSubConjuntos(numeros[1:], alvo) {
		return true
	}

	return somaSubConjuntos(numeros[1:], alvo-numeros[0])
}

    

func main() {
    var n, k int
	fmt.Scan(&n, &k)

	numeros := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numeros[i])
	}

    if somaSubConjuntos(numeros, k) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
