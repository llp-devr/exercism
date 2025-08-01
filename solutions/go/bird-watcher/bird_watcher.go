package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	totalBirds := 0
    for _, numOfBirds := range birdsPerDay {
        totalBirds += numOfBirds
    }
    return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    weekSlice := birdsPerDay[7*(week-1):(7*week)]
    
    totalBirds := 0
    for _, numOfBirds := range weekSlice {    
        totalBirds += numOfBirds
    }
    return totalBirds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i := 0; i < len(birdsPerDay); i++ {
		if i % 2 == 0 {
			birdsPerDay[i]++
		}
	}
	return birdsPerDay
}
