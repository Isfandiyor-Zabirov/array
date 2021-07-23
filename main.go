package main

import "fmt"

func main() {
	//string
	//int
	//rune
	// [capacity]type ---> 0, 1, 2, 3, 4, 5 ... n
	//

	var array [4]int
	array[0] = 111
	array[1] = 222
	array[2] = 666
	array[3] = 999


	fmt.Println(array)
	fmt.Println("Len of array =",len(array))

	for i := 0; i < len(array); i++ {
		//array[i] = 0
	}

	fmt.Println(array)

	for index, value := range array {
		fmt.Printf("Array[%d] = %d \n", index, value)
	}

	for _, value := range array {
		fmt.Println(value)
	}

	fmt.Println(array)

	for _, value := range array {
		value++
	}
	fmt.Println(array)

	for index, _ := range array {
		array[index]++
	}
	fmt.Println(array)

}
