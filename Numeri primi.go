package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	//check se è primo brutalee
	for i := n - 1; i > 1; i-- {
		if n%i == 0 {
			fmt.Println("non è primo")
			return
		}
	}
	fmt.Println("we è primo")
}
