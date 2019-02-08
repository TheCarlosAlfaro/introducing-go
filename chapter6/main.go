package main

import "fmt"

func average(num []float64) float64 {
	total := 0.0

	for _, v := range num {
		total += v
	}
	return total / float64(len(num))
}

func main() {
	fmt.Println("Chapter 6")

	xs := []float64{98, 93, 77, 82, 83}

	fmt.Println(average(xs))

}
