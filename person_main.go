package main

// should probably be package person

import (
	"fmt"
	. "github.com/jayteesf/person"
)

func main() {
	var person = &Person{Name: "JayTeeSF", Email: "jaytee@jayteesf.com"}
	var other_person = &Person{Name: "JayTeeSFOther", Email: "jaytee+other@jayteesf.com"}
	other_person.Save()
	other_person_json, err := other_person.AsJson()
	if err != nil {
		panic(err)
	}
	fmt.Println("saved person: ", string(other_person_json))
	person.Save()
	person_json, err := person.AsJson()
	if err != nil {
		panic(err)
	}
	fmt.Println("saved person: ", string(person_json))
	loaded_person, err := LoadPerson(person.Email)
	if err != nil {
		panic(err)
	}
	fmt.Printf("got person: %+v\n", loaded_person)

	fmt.Println("done")
}
