// how long it take to get to Mars?
package main

import "fmt"

func main() {
	const hoursPerDay = 24  // hours per day
	var distance = 96300000 // km
	var speed = 100800      // km/h

	fmt.Printf("%v days to reach Mars.", (distance / speed / hoursPerDay))

}
