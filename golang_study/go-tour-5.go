package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Now you have %g problem.",
			math.Nextafter(2, 3))
}
