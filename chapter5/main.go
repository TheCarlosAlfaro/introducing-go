package main

import "fmt"

func main() {
	fmt.Println("Chapter 5")
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	myArray := [5]float64{98, 93, 77, 82, 83}

	var total float64 = 0

	for _, value := range myArray {

		total += value
	}
	fmt.Println(total / float64(len(myArray)))

	myMap := make(map[string]int)
	myMap["key"] = 10
	fmt.Println(myMap["key"])

	excerse := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	fmt.Println("Smallest Number", findSmallestNum(excerse))
}

func findSmallestNum(list []int) int {
	smallest := list[0]
	for _, num := range list {
		if num < smallest {
			smallest = num
		}
	}
	return smallest
}
