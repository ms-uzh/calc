package main

import "github.com/fforootd/calc/models"

func calculateQuaternary(head models.Head, tail models.Tail, polyamines ...models.Polyamine) (quaternary int) {
	quaternary = head.Quaternary
	quaternary += tail.Quaternary
	for _, polyamine := range polyamines {
		quaternary += polyamine.Quaternary
	}
	return quaternary
}
