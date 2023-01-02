package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	return sum(birdsPerDay)
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	week = week - 1
	start := (week * 7)
	end := start + 8
	if len(birdsPerDay) <= end {
		end = len(birdsPerDay) - 1
	}
	return sum(birdsPerDay[start:end])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i] = birdsPerDay[i] + 1
		}
	}
	return birdsPerDay
}

// sum takes a slice of integers and returns the sum of them all
func sum(ints []int) int {
	runningTotal := 0
	for i := range ints {
		runningTotal = runningTotal + ints[i]
	}
	return runningTotal
}
