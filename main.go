package main

import (
	"fmt"
	"math/rand"

	"github.com/arnemolland/exinc/interval"
)

func main() {
	include := []interval.Interval{{Start: 1, End: 3}, {Start: 5, End: 8}}
	exclude := []interval.Interval{{Start: 2, End: 4}, {Start: 6, End: 7}}

	result := interval.ProcessIntervals(include, exclude)
	fmt.Println("Resultant Intervals:", result)
}

func generateRandomIntervals(n int) []interval.Interval {
	intervals := make([]interval.Interval, n)
	for i := 0; i < n; i++ {
		r1 := rand.Int()
		r2 := rand.Int()

		intervals[i] = interval.Interval{
			Start: min(r1, r2),
			End:   max(r1, r2),
		}
	}
	return intervals
}
