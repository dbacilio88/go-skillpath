package main

import (
	"fmt"
)

//https://github.com/go-chi/chi
//https://gobyexample.com/

func main() {

	medical := Person{
		Name: "David",
		Age:  40,
	}

	fmt.Println(medical)

}

type Person struct {
	Name string
	Age  int
}
