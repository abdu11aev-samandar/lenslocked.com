package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Dog  string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name: "John Smith",
		Dog:  "Pug",
	}

	data.Name = "John Calhoun"
	data.Dog = "Beagle"
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
