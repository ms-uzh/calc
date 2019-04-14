package calculation

import "github.com/fforootd/calc/models"

func CalculateHDX(head models.Head, tail models.Tail, polyamines ...models.Polyamine) uint {
	hdx := head.HDX
	for _, polyamine := range polyamines {
		hdx += polyamine.HDX
	}
	hdx += tail.HDX

	return hdx
}
