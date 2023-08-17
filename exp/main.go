package main

import (
	"html/template"
	"os"
)

type Dog struct {
	Name string
}

type User struct {
	Dog   Dog
	Slice []string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Dog: Dog{
			Name: "Mock",
		},
		Slice: []string{"a", "b", "c"},
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
