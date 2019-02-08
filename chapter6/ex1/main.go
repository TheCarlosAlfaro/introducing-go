package main

import "fmt"

func main() {
	myNumbers := []int{1, 2, 3, 4, 5, 6}
	total := sum(myNumbers)
	fmt.Println(total)
}

func sum(n []int) int {
	total := 0
	for i := 0; i < len(n); i++ {
		total += n[i]; 
	}
	return total
}