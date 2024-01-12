package main

import (
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	tests := map[string]struct{
		image [][]int
		sr int
		sc int
		color int
		want [][]int
	}{
		"case 1": {
			image: [][]int{{1,1,1},{1,1,0},{1,0,1}},
			sr: 1,
			sc: 1,
			color: 2,
			want: [][]int{{2,2,2},{2,2,0},{2,0,1}},
		},
		"case 2": {
			image: [][]int{{0,0,0},{0,0,0}},
			sr: 0,
			sc: 0,
			color: 0,
			want: [][]int{{0,0,0},{0,0,0}},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := floodFill(tc.image, tc.sr, tc.sc, tc.color)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("floodFill(%v, %v, %v, %v) = %v, want %v", tc.image, tc.sr, tc.sc, tc.color, got, tc.want)
			}
		})
	}
}
