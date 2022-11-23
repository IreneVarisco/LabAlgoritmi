package main

import (
	"fmt"
)

type Dipendente struct {
	nome        string
	subordinati []*Dipendente
}

func main() {
	var isco *Dipendente
	isco = NuovoAlbero()
	CostruisciAlbero(isco)

	stampaSubordinati(isco)
	fmt.Println(quantiSenzaSubordinati(isco), "dipendenti senza subordinati")

	fmt.Println("supervisore di", isco.subordinati[1].subordinati[1].nome, ":", supervisore(isco.subordinati[1].subordinati[1], isco).nome)

	fmt.Println("impiegati sopra a", isco.subordinati[1].subordinati[1].nome, ":")
	stampaImpiegatiSopra(isco.subordinati[1].subordinati[1], isco)

	fmt.Println()
	stampaImpiegatiConSupervisore(isco)
}

func NuovoAlbero() *Dipendente {
	return CreaNodo("Giovanni")
}

func AggiungiNodo(d *Dipendente, daAggiungere *Dipendente) {
	d.subordinati = append(d.subordinati, daAggiungere)
}

func CreaNodo(nome string) *Dipendente {
	return &Dipendente{nome, nil}
}

// questa funzione costruisce l'albero
func CostruisciAlbero(isco *Dipendente) {
	AggiungiNodo(isco, CreaNodo("Franco"))
	figlio := CreaNodo("Lore")
	AggiungiNodo(isco, figlio)
	AggiungiNodo(isco, CreaNodo("Carlo"))
	AggiungiNodo(figlio, CreaNodo("Juan"))
	AggiungiNodo(figlio, CreaNodo("Ale"))
}

func stampaSubordinati(d *Dipendente) {
	fmt.Println("elenco subordinati di", d.nome)
	for _, v := range d.subordinati {
		fmt.Println(v.nome)
	}
	fmt.Println("")
}

func quantiSenzaSubordinati(d *Dipendente) int {
	n := 0
	if d.subordinati == nil {
		n++
	} else {
		for _, v := range d.subordinati {
			n = n + quantiSenzaSubordinati(v)
		}
	}
	return n
}

// stampa il supervisore del dipendente inserito
func supervisore(d *Dipendente, radice *Dipendente) *Dipendente {
	if radice.subordinati == nil {
		return nil
	}
	for _, v := range radice.subordinati {
		if v == d {
			return radice
		} else {
			x := supervisore(d, v)
			if x != nil {
				return x
			}
		}
	}
	return nil
}

func stampaImpiegatiSopra(d *Dipendente, radice *Dipendente) {

	thasup := supervisore(d, radice)

	if thasup != radice {
		// stampo il supervisore e chiamo stampaImpiegatiSopra sul supervisore
		stampaImpiegatiSopra(thasup, radice)
	}
	fmt.Println(thasup.nome)
}

/*
Stampare l’elenco di tutti i dipendenti –non importa l’ordine–, indicando per ciascuno chi è
il suo supervisore (tranne nel caso di dipendenti di massimo livello).
*/
func stampaImpiegatiConSupervisore(d *Dipendente) {
	for _, v := range d.subordinati {
		fmt.Println("supervisore di", v.nome, ":", supervisore(v, d).nome)
		if v.subordinati != nil {
			stampaImpiegatiConSupervisore(v)
		}
	}
}
