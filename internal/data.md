 ¿Qué es interface{} en Go?
 Es el tipo más genérico en Go: representa cualquier tipo de dato.

 	var data interface{}
	data = 1
	data = "123"
	data = true
	fmt.Printf("Valor: %v, Tipo: %T\n", data, data)

¿Qué es any? ¿Es diferente de interface{}?
any fue introducido en Go 1.18 como un alias de interface{}, simplemente para legibilidad.
	type any = interface{}
"Este parámetro puede ser cualquier cosa, no me importa el tipo específico".
```go
package main

import (
	"fmt"
	"reflect"
)

//https://github.com/go-chi/chi
//https://gobyexample.com/

func main() {
	data := []any{"Hola mundo", 123, true, 1.23, 'A', -1}
	for _, v := range data {
		showValues(v)
	}
}

// func showValues(data interface{}) {
func showValues(data any) {
	switch value := data.(type) {
	case int:
		print(value)
	case string:
		print(value)
	case bool:
		print(value)
	case any:
		print(value)
	default:
		print("No")
	}
}

func print(value any) {
	fmt.Println("Reflection: ", reflect.TypeOf(value))
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}

```

En Go, un switch es una estructura de control que permite ejecutar diferentes bloques de código según el valor de una expresión, de forma más limpia que usar múltiples if/else if.

Sintaxis básica de switch
```go
switch expresion {
case valor1:
    // código si expresion == valor1
case valor2:
    // código si expresion == valor2
default:
    // código si ningún caso coincide
}
```

 ¿Qué es un struct en Go?
 Un struct (estructura) en Go es un tipo compuesto que te permite agrupar múltiples campos con nombre, cada uno con su propio tipo.

 ```go
 type NombreStruct struct {
    Campo1 Tipo1
    Campo2 Tipo2
}
 ```