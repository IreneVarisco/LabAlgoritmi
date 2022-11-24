package main

import (
	"fmt"
	"strconv"
	"strings"
)

type grafo struct {
	n         int // numero di vertici
	adiacenti []*linkedList
}

type linkedList struct {
	val  int
	next *linkedList
}

func main() {
	g := readGrafo(nuovoGrafo())
	printGrafo(g)
}

func nuovoGrafo() *grafo {
	g := new(grafo)
	return g
}

func readGrafo(g *grafo) *grafo {
	var isco string
	// formato n 0:adiacente1,adiacente2,adiacente3 ... n-1:adiacente1,adiacente2,adiacente3
	// 3 0:1,2 1:2 3:1
	fmt.Println("inserire il grafo nel formato n 0:adiacente1,adiacente2,adiacente3 1:adiacente1,adiacente2,adiacente3 ... n-1:adiacente1,adiacente2,adiacente3")
	fmt.Scan(&isco)

	// prendo il valore n
	a := strings.Split(isco, " ")
	n, _ := strconv.Atoi(a[0])
	// analizzo l'input carattere per carattere
	a = strings.Split(isco, "")
	// per ogni vertice creo una linked list
	for i := 0; i < n; i++ {
		g.adiacenti = append(g.adiacenti, new(linkedList))
	}
	// creazione della linked list per ogni vertice
	for i := 0; i < len(a); i++ {
		_, err := strconv.Atoi(a[i])

		if i != len(a)-1 {
			if err == nil && string(a[i+1]) == ":" {
				var ll *linkedList
				for j := i + 2; string(a[j+1]) != " "; j += 2 {
					ll.val, _ = strconv.Atoi(a[j])
					ll.next.val, _ = strconv.Atoi(a[j+2])
				}
				g.adiacenti[i] = ll
			}
		}
	}
	return g
}

func printGrafo(g *grafo) {
	c := 0
	for _, v := range g.adiacenti {
		fmt.Print(c, ":")
		for i := 0; i < len(g.adiacenti); i++ {
			if v.next != nil {
				fmt.Print(v.val, ",")
			} else {
				fmt.Println(v.val)
			}
		}
		c++
	}

}

func existsArch(x int, y int) bool {
	return true

}
