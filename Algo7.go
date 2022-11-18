package main

import "fmt"

type nome struct {
	key  int
	name string
}

func main() {
	var isco []nome = []nome{{3, "Francesco"}, {4, "Tommaso"}, {5, "Lorenzo"}, {2, "Elisa"}, {1, "Edoardo"}}

	for i := 0; i < len(isco)-1; i++ {
		for j := 0; j < len(isco)-i-1; j++ {
			if isco[j].key < isco[j+1].key {
				isco[j], isco[j+1] = isco[j+1], isco[j]
			}
		}
	}

	fmt.Println(isco)

}
