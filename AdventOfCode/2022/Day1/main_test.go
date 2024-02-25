package main

import (
	"fmt"
	"testing"
)

var test1String string = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

var testCases = []struct {
	given    string
	top3     int
	greatest int
}{
	{
		test1String,
		45000,
		24000,
	},
	{
		input(),
		201524,
		69281,
	},
}

func TestHelloWorld(t *testing.T) {
	for i, e := range testCases {
		testName := fmt.Sprintf("Test Case #%d", i)
		t.Run(testName, func(t *testing.T) {
			if res := SumTopCalorieCounts(e.given); res != e.top3 {
				t.Fatalf(`Oi! We got %d, but we wanted %d`, res, e.top3)
			}
			if greatest := GetGreatestCalories(e.given); greatest != e.greatest {
				t.Fatalf(`Oi! We got %d, but we wanted %d`, greatest, e.greatest)
			}
			greatestCal, top3Cal := GetGreatestCaloriesAndOther(e.given)
			if top3Cal != e.top3 {
				t.Fatalf(`Oi! We got %d, but we wanted %d`, top3Cal, e.top3)
			}
			if greatestCal != e.greatest {
				t.Fatalf(`Oi! We got %d, but we wanted %d`, greatestCal, e.greatest)
			}
		})
	}
}

func BenchmarkMine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetGreatestCaloriesAndOther(test1String)
	}
}
