package main

// should probably be package person

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name  string
	Email string
}

func Filename(email string) string {
	return email + ".txt"
}

func (p *Person) filename() string {
	return Filename(p.Email)
}

func (p *Person) AsJson() ([]byte, error) {
	return json.Marshal(p)
}

func personFromJson(jsonBlob []byte) (*Person, error) {
	person := &Person{}
	err := json.Unmarshal(jsonBlob, person)
	if err != nil {
		return &Person{}, err
	}
	return person, err
}

// at some point this will be part of an interface
func ReadJSONFrom(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func LoadPerson(email string) (*Person, error) {
	jsonBlob, err := ReadJSONFrom(Filename(email))
	if err != nil {
		return &Person{}, err
	}
	fmt.Printf("got JSON: %+v\n", string(jsonBlob))

	person, err := personFromJson(jsonBlob)
	//return &Person{Name: loaded_name, Email: loaded_email}, nil
	fmt.Printf("loaded: %+v", person)
	return person, err
}

func (p *Person) Save() error {
	json_bytes, err := p.AsJson()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(p.filename(), json_bytes, 0600)
}

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
