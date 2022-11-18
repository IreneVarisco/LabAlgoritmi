package main

import (
	"fmt"
	"strconv"
	"strings"
)

type dot struct {
	x int
	y int
}

func main() {
	var isco []dot
	var a string

	for !strings.HasPrefix(a, "fold along") {
		fmt.Scan(&a)

		if !strings.HasPrefix(a, "fold along") {
			x, _ := strconv.Atoi(string(a[0]))
			y, _ := strconv.Atoi(string(a[2]))
			isco = append(isco, dot{x, y})
		} else {
			if string(a[11]) == "x" {
				var mappa [5][14]int
				for _, punto := range isco {
					b := 10
					for b > 1 {
						if punto.x > 5 {
							punto.x = punto.x - b
							b = b - 2
						}
					}
					mappa.append(mappa, punto)

				}

			}

			if string(a[11]) == "y" {

			}
		}

	}

}
