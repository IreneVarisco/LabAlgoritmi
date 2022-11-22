package main

import (
	"fmt"
	"math/rand"
)

type bitreeNode struct {
	left  *bitreeNode
	right *bitreeNode
	val   int
}
type bitree struct {
	root *bitreeNode
}

func main() {
	var isco []int
	//creo un array co numeri random
	for i := 0; i < 5; i++ {
		isco = append(isco, rand.Intn(100))
	}
	//creo dall'array un albero
	root := arr2tree(isco, 0)
	//stampo l'albero
	stampaAlberoASommario(root, 0)

}
func arr2tree(isco []int, i int) (root *bitreeNode) {

	if i >= len(isco) {
		return nil
	}
	// assegna a root l'indirizzo di una nuova struttura di tipo bitreeNode, con all'interno i valori: puntatore al figlio di sinistra, puntatore al figlio di destra e valore attuale
	// NB: senza assegnare un nuovo indirizzo al puntatore root all'inizio di questa funzione prende il valore nil.
	root = &bitreeNode{arr2tree(isco, 2*i+1), arr2tree(isco, 2*i+2), isco[i]}

	return root

}
func stampaAlberoASommario(node *bitreeNode, spaces int) {
	//stampo numero di spazi e in seguito il valore preceduto da asterisco
	for i := spaces; i > 0; i-- {
		fmt.Print(" ")
	}
	fmt.Print("*")
	if node == nil {
		return
	}
	fmt.Println(node.val)
	//controllo che il nodo abbia almeno un figlio e chiamo ricorsivamente la funzione
	if node.left != nil || node.right != nil {
		stampaAlberoASommario(node.left, spaces+1)
		stampaAlberoASommario(node.right, spaces+1)
	}

}

// modi strani per stampare l'albero lol
func inorder(node *bitreeNode) {
	if node != nil {
		inorder(node.left)
		fmt.Println(node.val)
		inorder(node.right)
	}
}
func preorder(node *bitreeNode) {
	if node != nil {
		fmt.Println(node.val)
		inorder(node.left)
		inorder(node.right)
	}
}
func postorder(node *bitreeNode) {
	if node != nil {
		inorder(node.left)
		inorder(node.right)
		fmt.Println(node.val)
	}
}
