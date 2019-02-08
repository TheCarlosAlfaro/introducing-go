package main

import "fmt"

func main() {

	fmt.Println(toCelsius(40))
	fmt.Println(toMeters(4))
}

func toCelsius(fa int) int {
	return fa - 32*5/9
}

func toMeters(feet float32) float32 {
	return feet / 3.2808
}
