package main

const OvenTime = 40

func RemainingOvenTime(actual int) int {
	return OvenTime - actual
}

func PreparationTime(numberOfLayers int) int {
	preparation := 2
	return preparation * numberOfLayers
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}

func main() {
	RemainingOvenTime(30)

	PreparationTime(2)

	ElapsedTime(3, 20)
}
