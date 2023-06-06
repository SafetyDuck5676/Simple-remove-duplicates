package main

import (
	"fmt"
)

func main() {
	input := []int{3, 4, 4, 3, 6, 3}

	fmt.Println(input)
	allKeys := make(map[int]bool)
	list := []int{}
	finallist := []int{}
	iteration := 1

	for i := len(input) - 1; i >= 0; i-- {
		fmt.Printf("Iteration: %d \n", iteration)
		fmt.Printf("allKeys value is: %v \n", allKeys[input[i]])
		fmt.Printf("at position : %v \n", input[i])

		_, value := allKeys[input[i]]

		if !value {
			allKeys[input[i]] = true
			list = append(list, input[i])
		}

		iteration++
		fmt.Println(" ")
	}
	for i := len(list) - 1; i >= 0; i-- {
		finallist = append(finallist, list[i])
	}
	fmt.Println(finallist)
}
