package types

import "math/rand/v2"

var (
	Tiers = map[int]string{
		0: "grey/white",
		1: "light-blue",
		2: "blue",
		3: "purple",
		4: "pink",
	}
	Wears = [5]string{"FN", "MW", "FT", "WW", "BS"}
)

func RandomTier() string {
	i := rand.IntN(len(Tiers))
	return Tiers[i]
}

func RandomWear() string {
	i := rand.IntN(len(Wears))
	return Wears[i]
}
