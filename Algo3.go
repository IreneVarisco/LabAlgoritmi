package main

import (
	"fmt"
)

func main() {
	var a, b string

	fmt.Scan(&a)
	fmt.Scan(&b)

	//fmt.Println(b[0], a[0])
	fmt.Println(b[0] - a[0] + 1)
}
