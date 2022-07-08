package main

import "fmt"

// Não há operador ternário
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}

	return "Reprovado"
}

func main() {
	fmt.Printf(obterResultado(6.2))
}
