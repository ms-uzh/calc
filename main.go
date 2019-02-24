package main

import (
	"fmt"

	"github.com/fforootd/calc/utils"
)

func main() {
	// Set some Testparams
	//h := models.Head{Mass: 111.11111}
	//p := models.PolyamineUnit{Mass: 222.22222}
	//t := models.Tail{Mass: 555.55555}

	head := 111.11111
	poly := 222.22222
	tail := 555.55555

	frag := utils.CallCulateFragements(head, poly, tail)

	fmt.Println(frag)
	/*
		total := utils.CalcTotal()
		fmt.Println("hop: ", total)

		headtotal := utils.CalcHead()
		fmt.Println("the head is: ", headtotal)

		tailtotal := utils.CalcHead()
		fmt.Println("the head is: ", tailtotal)
	*/
}
