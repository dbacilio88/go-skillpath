package interfaces

import "fmt"

type Dog struct {
	Name string
}

func NewDog(name string) Dog {
	return Dog{
		Name: name,
	}
}

func (d Dog) Sound() string {
	return fmt.Sprintf("My name is %s and sound is %s", d.Name, "Guau")
}

func (d Dog) ChangeName(name string) Dog {
	d.Name = name
	return d
}

func (d *Dog) ChangeNamePointer(name string) {
	d.Name = name
}
