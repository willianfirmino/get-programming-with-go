package main

import "fmt"

func main() {
	const days = 28
	const hoursPerDay = 24
	var distance = 56000000 //km
	speed := distance / (days * hoursPerDay)
	fmt.Printf("%v km/h a ship needed to be to reach malacandra.", speed)
}
