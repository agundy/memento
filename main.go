package main

import (
	"./originator"
	"fmt"
	"strconv"
)

func simpleExample() {
	origin := originator.OriginatorType{}

	origin.Set("State One")
	fmt.Println("Originator: ", origin.Get())

	memento := origin.CreateMemento()

	origin.Set("State Two")

	fmt.Println("Originator: ", origin.Get())

	origin.SetFromMemento(memento)

	fmt.Println("Originator: ", origin.Get())

	// Shows that outside the object we can't get state of memento
	// Won't compile
	//fmt.Println("MementoType: ", memento.getState())
}

func complexExample() {
	var history []originator.MementoType
	origin := originator.OriginatorType{}

	for i := 0; i < 10; i++ {
		origin.Set("state: " + strconv.Itoa(i))
		fmt.Println(origin.Get())
		m := origin.CreateMemento()

		history = append(history, m)
	}

	fmt.Println(history)
	origin.SetFromMemento(history[3])
	fmt.Println("Originator: ", origin.Get())
}

func main() {
	simpleExample()
	complexExample()
}
