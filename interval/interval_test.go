package interval

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals []Interval
		want      []Interval
	}{
		{"No intervals", []Interval{}, []Interval{}},
		{"Single interval", []Interval{{1, 3}}, []Interval{{1, 3}}},
		{"Non-overlapping intervals", []Interval{{1, 3}, {4, 6}}, []Interval{{1, 3}, {4, 6}}},
		{"Overlapping intervals", []Interval{{1, 3}, {2, 6}}, []Interval{{1, 6}}},
		{"Multiple overlaps", []Interval{{1, 3}, {2, 6}, {5, 8}}, []Interval{{1, 8}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeIntervals(tt.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExcludeIntervals(t *testing.T) {
	tests := []struct {
		name    string
		include []Interval
		exclude []Interval
		want    []Interval
	}{
		{"No intervals", []Interval{}, []Interval{}, []Interval{}},
		{"No exclude intervals", []Interval{{1, 5}}, []Interval{}, []Interval{{1, 5}}},
		{"Non-overlapping", []Interval{{1, 5}}, []Interval{{6, 7}}, []Interval{{1, 5}}},
		{"Exclude within include", []Interval{{1, 5}}, []Interval{{2, 3}}, []Interval{{1, 1}, {4, 5}}},
		{"Exclude overlapping include", []Interval{{1, 5}}, []Interval{{0, 3}}, []Interval{{4, 5}}},
		{"Multiple intervals", []Interval{{1, 5}, {6, 10}}, []Interval{{2, 3}, {7, 8}}, []Interval{{1, 1}, {4, 5}, {6, 6}, {9, 10}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExcludeIntervals(tt.include, tt.exclude); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExcludeIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessIntervals(t *testing.T) {
	tests := []struct {
		name     string
		includes []Interval
		excludes []Interval
		want     []Interval
	}{
		{
			name:     "Example 1",
			includes: []Interval{{10, 100}},
			excludes: []Interval{{20, 30}},
			want:     []Interval{{10, 19}, {31, 100}},
		},
		{
			name:     "Example 2",
			includes: []Interval{{50, 5000}, {10, 100}},
			excludes: []Interval{},
			want:     []Interval{{10, 5000}},
		},
		{
			name:     "Example 3",
			includes: []Interval{{200, 300}, {50, 150}},
			excludes: []Interval{{95, 205}},
			want:     []Interval{{50, 94}, {206, 300}},
		},
		{
			name:     "Example 4",
			includes: []Interval{{200, 300}, {10, 100}, {400, 500}},
			excludes: []Interval{{410, 420}, {95, 205}, {100, 150}},
			want:     []Interval{{10, 94}, {206, 300}, {400, 409}, {421, 500}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProcessIntervals(tt.includes, tt.excludes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMergeIntervals(b *testing.B) {
	intervals := []Interval{{1, 3}, {2, 6}, {5, 8}}

	for i := 0; i < b.N; i++ {
		MergeIntervals(intervals)
	}
}

func BenchmarkExcludeIntervals(b *testing.B) {
	include := []Interval{{200, 300}, {10, 100}, {400, 500}}
	exclude := []Interval{{10, 94}, {206, 300}, {400, 409}, {421, 500}}

	for i := 0; i < b.N; i++ {
		ExcludeIntervals(include, exclude)
	}
}

func BenchmarkProcessIntervals(b *testing.B) {
	include := []Interval{{200, 300}, {10, 100}, {400, 500}}
	exclude := []Interval{{10, 94}, {206, 300}, {400, 409}, {421, 500}}

	for i := 0; i < b.N; i++ {
		ProcessIntervals(include, exclude)
	}
}
