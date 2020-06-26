package fragment

import (
	"github.com/ms-uzh/calc/calculation"
	"github.com/ms-uzh/calc/models"
)

func calculateTA(previousMass float64, current, following models.Polyamine, tail models.Tail, isFirst, isLast bool) float64 {
	if isFirst && isLast {
		return calculateLastTA(previousMass, current, tail)
	}
	if isFirst {
		return calculateFirstTA(previousMass, current, following)
	}
	if isLast {
		return calculateLastTA(previousMass, current, tail)
	}
	return calculateMiddleTAs(previousMass, current, following)
}

// mass(Head) + mass(PA_n) + Sub_mass(PA_n+1) –  Quat(PA_n+1) * C1 + C4
func calculateFirstTA(headMass float64, polyamine, following models.Polyamine) float64 {
	return headMass + polyamine.Mass + following.Sub.Mass - float64(following.Quaternary)*calculation.H + calculation.NH3 - calculation.MassElectron
}

//ta(n-1) + mass(PA_n) – Sub_mass(PA_n) + Sub_mass(PA_n+1) –  Quat(PA_n+1) * C1
func calculateMiddleTAs(previous float64, polyamine, following models.Polyamine) float64 {
	return previous + polyamine.Mass - polyamine.Sub.Mass + following.Sub.Mass - float64(following.Quaternary)*calculation.H
}

//ta(m–1) + mass(PA_m) – Sub_mass(PA_m)+ Sub_mass(Tail) –  Quat(Tail) * C1
func calculateLastTA(previous float64, polyamine models.Polyamine, tail models.Tail) float64 {
	return previous + polyamine.Mass - polyamine.Sub.Mass + tail.Sub.Mass - float64(tail.Quaternary)*calculation.H
}
