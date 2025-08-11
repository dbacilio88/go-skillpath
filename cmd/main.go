package main

import (
	"fmt"
	example "github.com/dbacilio88/go-skillpath/internal/interfaces"
)

type PersonAdapter interface {
	doctores()
}

//https://github.com/go-chi/chi
//https://gobyexample.com/

func main() {

	animals := []example.Animals{
		example.Dog{
			Name: "Pepe",
		},
		example.Cat{
			Name: "Cata",
		},
	}

	for _, w := range animals {
		fmt.Println(w.Sound())
	}

	cat := example.NewCat("Catalina")
	//	dog := example.NewDog("Toby")
	fmt.Println(cat.Sound())

	dog1 := example.Dog{
		Name: "Dalmata",
	}

	fmt.Println("1.", dog1.Sound())
	ch := dog1.ChangeName("A")
	fmt.Println("2 ch.", ch.Sound())
	dog1.ChangeNamePointer("a")
	fmt.Println("3.", dog1.Sound())
	ch.ChangeNamePointer("B")
	fmt.Println("4.", dog1.Sound())

	fmt.Println("final 1 ",dog1)
	fmt.Println("final 2 ",ch)

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
