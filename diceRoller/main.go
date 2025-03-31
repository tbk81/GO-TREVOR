package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("This is a dice roller")
	dice := die{
		d4:  4,
		d6:  6,
		d10: 10,
		d12: 12,
		d20: 20,
	}

	for range 100 {
		fmt.Println(roll(dice.d10))
	}
}

type die struct {
	d4  int
	d6  int
	d10 int
	d12 int
	d20 int
}

func roll(d int) int {
	return rand.Intn(d + 1)
}
