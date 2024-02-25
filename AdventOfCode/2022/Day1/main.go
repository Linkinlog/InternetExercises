package main

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
)

// GetGreatestCalories returns the greatest amount of calories in the string
// The string is a multiline string that contains amounts of calories separated by newlines
// And grouped by blank lines
func GetGreatestCaloriesAndOther(calorieList string) (int, int) {
	// Make a temp var to hold the current/temp totals
	var tempTotal int
	totals := make([]int, 0, 3)

	// Ingest the string and convert to ints
	scanner := bufio.NewScanner(strings.NewReader(calorieList))
	for scanner.Scan() {
		if scanner.Text() != "" {
			scannerInt, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			// Add every line to the temp total
			tempTotal += scannerInt
		} else {
			// If we hit a blank line, append total to totals and clear temp total
			totals = append(totals, tempTotal)
			tempTotal = 0
		}
	}

	// Append the final temp total
	totals = append(totals, tempTotal)

	// Sort ascending
	sort.Ints(totals)

	// Define return values and return them
	greatestCalorieCount := totals[len(totals)-1]
	top3Greatest := totals[len(totals)-1] + totals[len(totals)-2] + totals[len(totals)-3]
	return greatestCalorieCount, top3Greatest
}

// GetGreatestCalories returns the greatest amount of calories in the string
// The string is a multiline string that contains amounts of calories separated by newlines
// And grouped by blank lines
func GetGreatestCalories(calorieList string) int {
	// Make a temp var to hold the current/temp totals
	var greatestCalories int
	var tempTotal int
	// Ingest the string
	scanner := bufio.NewScanner(strings.NewReader(calorieList))
	for scanner.Scan() {
		if scanner.Text() != "" {
			scannerInt, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			// Add every line to the temp total
			tempTotal += scannerInt
		} else {
			// If we hit a blank line,
			// Check if the temp amount is greater than the current
			if greatestCalories < tempTotal {
				greatestCalories = tempTotal
			}
			tempTotal = 0
		}
	}
	return greatestCalories
}

// SumTopCalorieCounts returns the top 3 greatest amount of calories in the string
func SumTopCalorieCounts(calorieList string) int {
	// Make a temp var to hold the temp total and totals of all the elves
	totals := make([]int, 0, 3)
	tempTotal := 0
	// Ingest the string
	scanner := bufio.NewScanner(strings.NewReader(calorieList))
	for scanner.Scan() {
		if scanner.Text() != "" {
			// Convert the string to an integer
			scannerInt, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			// Add every line to the previous total
			tempTotal += scannerInt
		} else {
			// If we hit a blank line,
			// Append total to totals and clear temp total
			totals = append(totals, tempTotal)
			tempTotal = 0
		}
	}
	// Append the final temp total
	totals = append(totals, tempTotal)
	// Sort ascending, grab the last 3 and return their sum
	sort.Ints(totals)
	return totals[len(totals)-1] + totals[len(totals)-2] + totals[len(totals)-3]
}

func main() {
	// I'm doing my part!
}
