package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var distance = rand.Intn(401000000) + 56000000
	fmt.Println("Random distance ", distance)
}
