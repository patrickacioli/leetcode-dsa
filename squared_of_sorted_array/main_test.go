package main

import (
	"reflect"
	"testing"
)

func TestSquaredArrays(t *testing.T) {
	testCases := []struct {
		InputEntry []int
		Expected   []int
	}{
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}
	for _, tc := range testCases {
		response := sortedSquares(tc.InputEntry)

		if !reflect.DeepEqual(response, tc.Expected) {
			t.Errorf("got %v want %v", response, tc.Expected)
		}
	}
}
