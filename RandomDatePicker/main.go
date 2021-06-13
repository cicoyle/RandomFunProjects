package main

import (
	"fmt"
	"math/rand"
)

func main() {
	dates := []string{
		"Hopscotch in SA",
		"Wonderspaces",
		"Round Rock Donuts",
		"Hot Yoga",
		"Tantrum Yoga",
		"Race COTA Karts",
		"Yard House",
		"Mavericks Dancehall",
		"Rainforest Cafe in SA",
		"Kayak",
	}

	randomIndex := rand.Intn(len(dates))
	pick := dates[randomIndex]
	fmt.Println(pick)
}
