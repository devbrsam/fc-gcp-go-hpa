package main

import (
	"fmt"
	"math"
)

func sqrt(value float64, loop int) float64 {
	x := value
	for i := 0 ; i <= loop ; i++ {
		x += math.Sqrt(x)
	}
	fmt.Printf("Code.education Rocks!")
	return x
}