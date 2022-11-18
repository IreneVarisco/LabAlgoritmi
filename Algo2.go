package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	var isco1 float64
	var isco2 float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	isco1 = (-b + math.Sqrt(math.Pow(b, 2)-(4*a*c))) / (2 * a)
	isco2 = (-b - math.Sqrt(math.Pow(b, 2)-(4*a*c))) / (2 * a)

	fmt.Println(isco1, isco2)
}
