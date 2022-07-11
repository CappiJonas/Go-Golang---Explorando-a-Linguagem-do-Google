package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceitoSemSwitch(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	} else if n.entre(7.0, 8.99) {
		return "B"
	} else if n.entre(5.0, 7.99) {
		return "C"
	} else if n.entre(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

func notaParaConceitoComSwitch(n nota) string {
	switch {
	case n.entre(9.0, 10.0):
		return "A"
	case n.entre(7.0, 8.99):
		return "B"
	case n.entre(5.0, 7.99):
		return "C"
	case n.entre(3.0, 4.99):
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceitoSemSwitch(9.8))
	fmt.Println(notaParaConceitoSemSwitch(6.9))
	fmt.Println(notaParaConceitoSemSwitch(2.1))
	fmt.Println(notaParaConceitoComSwitch(9.8))
	fmt.Println(notaParaConceitoComSwitch(6.9))
	fmt.Println(notaParaConceitoComSwitch(2.1))
}
