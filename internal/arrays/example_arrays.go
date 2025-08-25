package arrays

import "fmt"

func Startup() {

	var arr [10]int

	for i := range 10 {
		arr[i] = i * 2
	}

	var slice = make([]int, 5)
	fmt.Printf("cap: %d, len %d\n", cap(slice), len(slice))
	for i := range 10 {

		slice = append(slice, i)
	}
	fmt.Println("slice:", slice)
	fmt.Println("slice:", slice[0:5])
	fmt.Println("array:", arr)
	fmt.Println("array:", arr[:5])

}
