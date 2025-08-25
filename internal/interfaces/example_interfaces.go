package interfaces

import "fmt"

type Animals interface {
	Sound() string
}

func Startup() {
	animals := []Animals{
		Dog{
			Name: "Pepe",
		},
		Cat{
			Name: "Cata",
		},
	}

	for _, w := range animals {
		fmt.Println(w.Sound())
	}

	cat := NewCat("Catalina")
	//	dog := example.NewDog("Toby")
	fmt.Println(cat.Sound())

	dog1 := Dog{
		Name: "Dalmata",
	}

	fmt.Println("1.", dog1.Sound())
	ch := dog1.ChangeName("A")
	fmt.Println("2 ch.", ch.Sound())
	dog1.ChangeNamePointer("a")
	fmt.Println("3.", dog1.Sound())
	ch.ChangeNamePointer("B")
	fmt.Println("4.", dog1.Sound())

	fmt.Println("final 1 ", dog1)
	fmt.Println("final 2 ", ch)

}
