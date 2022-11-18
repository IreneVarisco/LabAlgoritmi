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
	var pila []string
	var post string
	//scanner
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	isco = scanner.Text()
	//split dell'input
	tokens := strings.Split(isco, " ")

	//controllo ogni carattere dell'input per capire se è un numero o meno
	for _, token := range tokens {
		_, err := strconv.Atoi(token)
		if err == nil { //se è un numero viene aggiunto alla alla stringa
			post = post + token + " "
		}
		if err != nil { //se è un'operazione la aggiungo alla pila
			if token != "(" {
				pila, post = valuta(token, pila, post)
			}
		}

	}
	post = post + pila[len(pila)-1]
	fmt.Println(post)
}

func valuta(token string, pila []string, post string) ([]string, string) {
	//se è una parentesi chiusa poppo l'operazione dalla pila e la metto nella stringa
	if token == ")" {
		post = post + pila[len(pila)-1] + " "
		pop(pila, len(pila)-1)
	} else {
		pila = append(pila, token)
	}
	return pila, post
}

func pop(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
