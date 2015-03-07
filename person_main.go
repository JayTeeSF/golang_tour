package main

// should probably be package person

import (
	"fmt"
	"github.com/jayteesf/person"
)

func main() {
	var p1 = &person.Person{Name: "JayTeeSF", Email: "jaytee@jayteesf.com"}
	var p2 = &person.Person{Name: "JayTeeSFOther", Email: "jaytee+other@jayteesf.com"}
	p2.Save()
	p2_json, err := p2.AsJson()
	if err != nil {
		panic(err)
	}
	fmt.Println("saved person: ", string(p2_json))
	p1.Save()
	p1_json, err := p1.AsJson()
	if err != nil {
		panic(err)
	}
	fmt.Println("saved person: ", string(p1_json))
	reloaded_p1, err := person.LoadPerson(p1.Email)
	if err != nil {
		panic(err)
	}
	fmt.Printf("got person: %+v\n", reloaded_p1)

	fmt.Println("done")
}
