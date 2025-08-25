package main

import (
	"fmt"

	"github.com/dbacilio88/go-skillpath/internal/arrays"
)

type PersonAdapter interface {
	doctores()
}

//https://github.com/go-chi/chi
//https://gobyexample.com/

func main() {

	//example.Startup()
	arrays.Startup()

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
