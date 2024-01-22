package main

import "testing"

func TestMerge(t *testing.T) {
	c		:= merge(asChan(1,2,3), asChan(4,5,6), asChan(7,8,9))
	seen 	:= make(map[int]bool)
	for v := range c {
		if seen[v] {
			t.Errorf("saw %d at least twice", v)
		}
		seen[v] = true
	}

	for i := 1; i<=9; i++ {
		if !seen[i] {
			t.Errorf("didnt see %d", i)
		}
	}
}

func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c		:= merge(asChan(1,2,3), asChan(4,5,6), asChan(7,8,9))
		seen 	:= make(map[int]bool)
		for v := range c {
			if seen[v] {
				b.Errorf("saw %d at least twice", v)
			}
			seen[v] = true
		}

		for i := 1; i<=9; i++ {
			if !seen[i] {
				b.Errorf("didnt see %d", i)
			}
		}
	}
}

/*
pkg: github.com/PetyaKatsarova/more_go/pr8_justforfunc/27and28/main
cpu: 13th Gen Intel(R) Core(TM) i7-13850HX
BenchmarkMerge-28    	  139833	      8300 ns/op	    1010 B/op	      20 allocs/op
PASS
meaning: test was executed 139833 times,
8300 ns/op: on average, each operation in the benchmark test took approximately 8,300 nanoseconds to execute.
1010 B/op:On average, each operation allocated approximately 1,010 bytes of memory.
20 allocs/op: Each operation in the benchmark test resulted in approximately 20 allocations.
*/