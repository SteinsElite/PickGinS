package util

import (
	"math"
	"math/big"
)


const decimal = 18

// Amount2Float convert the amount in big.Int in wei to proper unit in float
func Amount2Float(amount *big.Int) float64{
	bf, _ := new(big.Float).SetString(amount.String())
	f,_ := new(big.Float).Quo(bf, big.NewFloat(math.Pow10(decimal))).Float64()
	return f
}
