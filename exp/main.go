package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Int   int
	Float float64
	Slice []string
	Map   map[string]string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name:  "John Smith",
		Int:   1,
		Float: 100.21,
		Slice: []string{"a", "b", "c"},
		Map: map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
