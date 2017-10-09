package main

import "fmt"

func main() {
	nthRow, nthCol := GetPosition()
	fmt.Printf("%d,%d\n", nthRow+1, nthCol+1)
}
