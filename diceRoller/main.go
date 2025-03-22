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

// type roller interface{
// 	roll()
// }

/*
package main

import (
	"fmt"
	"math"
)

func main() {
	sq1 := square{3, 4}
	ci1 := circle{12}

	fmt.Println("Area of the square is", info(sq1))
	fmt.Println("Area of the circle is", info(ci1))

}

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}
*/
