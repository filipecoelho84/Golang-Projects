// In this exercise I'm going to write some code
//to help me cook a brilliant lasagna from my favorite cooking book.

package main

//  Expected oven time in minutes
const OvenTime = 40

// Calculation of the remaining oven time in minutes
func RemainingOvenTime(actual int) int {
	return OvenTime - actual
}

// Calculation of the preparation time in minutes
func PreparationTime(numberOfLayers int) int {
	preparation := 2
	return preparation * numberOfLayers
}

// Calculation of the elapsed working time in minutes
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}

func main() {
	RemainingOvenTime(30)

	PreparationTime(2)

	ElapsedTime(3, 20)
}
