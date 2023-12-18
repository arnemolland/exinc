package interval

import (
	"sort"
)

// Interval represents an interval with a start and end.
type Interval struct {
	Start, End int
}

// Merge and sort the include intervals, then exclude the sorted exclude intervals.
func ProcessIntervals(includes, excludes []Interval) []Interval {
	sort.Slice(includes, func(i, j int) bool {
		return includes[i].Start < includes[j].Start
	})

	sort.Slice(excludes, func(i, j int) bool {
		return excludes[i].Start < excludes[j].Start
	})

	merged := MergeIntervals(includes)

	result := ExcludeIntervals(merged, excludes)

	return result
}

// Merge overlapping intervals.
func MergeIntervals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	var merged []Interval
	current := intervals[0]

	for _, interval := range intervals {
		if interval.Start <= current.End {
			if interval.End > current.End {
				current.End = interval.End
			}
		} else {
			merged = append(merged, current)
			current = interval
		}
	}
	merged = append(merged, current)

	return merged
}

// Exclude intervals from a set of include intervals.
func ExcludeIntervals(includes, excludes []Interval) []Interval {
	result := []Interval{}
	i, j := 0, 0

	for i < len(includes) && j < len(excludes) {
		inc, exc := includes[i], excludes[j]

		// If exclude interval is past than the entire include interval, add the include interval as is.
		if exc.Start > inc.End {
			result = append(result, inc)
			i++
		} else if exc.End < inc.Start {
			// If exclude interval is before the entire include interval, ignore the exclude interval.
			j++
		} else {
			// Partial or complete overlap
			if exc.Start > inc.Start {
				result = append(result, Interval{Start: inc.Start, End: exc.Start - 1})
			}
			if exc.End < inc.End {
				includes[i].Start = exc.End + 1
				j++
			} else {
				i++
			}
		}
	}

	// Append any remaining include intervals
	if i < len(includes) {
		result = append(result, includes[i:]...)
	}

	return result
}
