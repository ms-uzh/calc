package main

import "fmt"

// TODO Comment
const decay = 1.00782503223

type head struct {
	mass float64
}

type polyamineUnit struct {
	mass            float64
	quaternaryamine bool
}

type tail struct {
	mass            float64
	quaternaryamine bool
}

type result struct {
	fragments []float64
}

func main() {
	// Set some Testparams
	head := head{
		mass: 111.11111,
	}
	poly := polyamineUnit{
		mass: 222.22222,
	}
	tail := tail{
		mass: 555.55555,
	}

	total := calcTotal(head.mass, poly.mass, tail.mass)
	fmt.Println("hop: ", total)

	headtotal := calcHead(tail.mass, total)
	fmt.Println("the head is: ", headtotal)

	tailtotal := calcHead(head.mass, total)
	fmt.Println("the head is: ", tailtotal)

}

func calcTotal(input ...float64) float64 {
	total := 0.0
	for _, value := range input {
		total += value
	}
	return total
}

func calcHead(tail, total float64) float64 {
	x := total - tail
	return x
}

func calcTail(head, total float64) float64 {
	y := total - head
	return y
}
