package interfaces

import "fmt"

type Cat struct {
	Name string
}

func NewCat(name string) Cat {
	return Cat{
		Name: name,
	}
}
func (c Cat) Sound() string {
	return fmt.Sprintf("My name is %s and sound is %s", c.Name, "Miau")
}
