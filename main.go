package main

import "log"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{"Samson", "German Shephered"}
	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Jock",
		Color:         "Grey",
		NumberOfTeeth: 38,
	}
	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	log.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string {
	return "Woof"
}
func (d *Dog) NumberOfLegs() int {
	return 4
}
func (d *Gorilla) Says() string {
	return "Ugh"
}
func (d *Gorilla) NumberOfLegs() int {
	return 2
}
