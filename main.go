package main

import (
	"fmt"

	"github.com/fforootd/calc/utils"
)

func main() {
	// Set some Testparams
	head := 112.11111
	poly := 222.22222
	tail := 555.55555

	name1 := "hodor"
	name2 := "blub"

	frag := utils.CallCulateFragements(head, poly, tail)
	total := utils.CalcTotal(head, poly, tail)

	fmt.Printf("%+v\n", frag)
	fmt.Println("total mass: ", total)
}
