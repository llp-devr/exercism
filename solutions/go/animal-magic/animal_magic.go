package chance

import (
    "math/rand"
    "time"
)

// SeedWithTime seeds math/rand with the current computer time.
func SeedWithTime() {
	rand.Seed(time.Now().UnixNano())
}

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return rand.Intn(20) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return float64(rand.Intn(12)) + rand.Float64()
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	slice := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
    var newSlice []string

    for len(slice) > 0 {
		// Pick a random index
		index := rand.Intn(len(slice))

		// Append the element at the random index to the destination slice
		newSlice = append(newSlice, slice[index])

		// Remove the element from the source slice
		slice = append(slice[:index], slice[index+1:]...)
	}

    return newSlice
}