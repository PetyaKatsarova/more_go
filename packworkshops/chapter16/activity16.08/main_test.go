package main

import "testing"

func Test_Sum(t *testing.T) {
	for i := 1; i < 20; i++ {
		res := sum(i, 1, 100)
		if res != 5050 {
			t.Errorf("Expected 5050 with %d workers but got: %d", i, res)
		}
	}
}