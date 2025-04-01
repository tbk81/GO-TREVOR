package main

import "fmt"

func main() {
	t1 := travelerProfile{}
	travelerMaker(t1)
	fmt.Println("Your name is", t1.name)
}

type travelerProfile struct {
	name string
	// origin    string
	// homeworld string
	// archtype  string
	// class     string
	// age       int
	// consLevel int
}

func travelerMaker(t travelerProfile) {
	fmt.Print("Type name: ")
	fmt.Scanln(&t.name)
	// fmt.Println("Your name is", t.name)
}

/*
Character sheet for Exodus TTRPG
*/
