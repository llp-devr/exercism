package lasagna

// OvenTime is the standard duration in minutes that a lasagna should be baked in the oven.
const OvenTime int = 40

// RemainingOvenTime takes the actual minutes the lasagna has been in the oven
// and returns how many minutes are left for the lasagna to be fully baked.
// This function assumes that the standard oven time for a lasagna is defined by the OvenTime constant.
func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

// PreparationTime calculates the time in minutes needed to prepare the lasagna.
// It does this based on the number of layers since each layer takes 2 minutes to prepare.
func PreparationTime(numberOfLayers int) int {
    var MinutesPerLayer int = 2
	return numberOfLayers * MinutesPerLayer
}

// ElapsedTime calculates the total time in minutes that has been spent so far cooking the lasagna.
// This includes the time spent preparing the lasagna layers and the actual time the lasagna has been baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}
