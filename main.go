package main

import "fmt"

const decay = 1.00782503223

var head = 111.11111
var pol1 = 222.22222
var tail = 333.33333

func main() {
	allInputs := []float64{head, pol1, tail}
	total := calcTotal(allInputs)
	fmt.Println("the total is: ", total)
	headtotal := calcHead(total, tail)
	fmt.Println("the head is: ", headtotal)
	tailtotal := calcTail(total, head)
	fmt.Println("the head is: ", tailtotal)
}

func calcTotal(allInputs []float64) float64 {
	total := 0.0
	for _, value := range allInputs {
		total += value
	}
	return total
}

func calcHead(total, tail float64) float64 {
	x := total - tail
	return x
}

func calcTail(total, head float64) float64 {
	y := total - head
	return y
}
