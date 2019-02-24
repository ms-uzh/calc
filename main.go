package main

import (
	"fmt"

	"github.com/fforootd/calc/utils"
)

func main() {
	// Set some Testparams
	head := 111.11111
	poly := 222.22222
	tail := 555.55555

	frag := utils.CallCulateFragements(head, poly, tail)
	total := utils.CalcTotal(head, poly, tail)

	fmt.Printf("%+v\n", frag)
	fmt.Println("total mass: ", total)
	/*
		total := utils.CalcTotal()
		fmt.Println("hop: ", total)

		headtotal := utils.CalcHead()
		fmt.Println("the head is: ", headtotal)

		tailtotal := utils.CalcHead()
		fmt.Println("the head is: ", tailtotal)
	*/
}
