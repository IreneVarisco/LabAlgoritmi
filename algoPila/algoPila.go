package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var isco string
	var pila []int

	//scanner
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	isco = scanner.Text()

	//split dell'input
	tokens := strings.Split(isco, " ")

	//controllo ogni carattere dell'input per capire se è un numero o meno
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil { //se è un numero viene aggiunto alla pila
			pila = append(pila, val)
		}
		if err != nil { //se è un'operazione la faccio tra gli ultimi 2 valori inseriti nella pila
			pila = valuta(token, pila)
			fmt.Println(pila)
		}

	}

}

func pop(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func valuta(token string, pila []int) []int {
	var result int
	switch token {
	case "+":
		result = pila[len(pila)-2] + pila[len(pila)-1]
	case "-":
		result = pila[len(pila)-2] - pila[len(pila)-1]
	case "*":
		result = pila[len(pila)-2] * pila[len(pila)-1]
	case "/":
		result = pila[len(pila)-2] / pila[len(pila)-1]
	}
	pila = pop(pila, len(pila)-1) //rimuovo gli ultimi 2 elementi della pila
	pila = pop(pila, len(pila)-1)
	pila = append(pila, result) //inserisco il risultato nella pila

	return pila
}
