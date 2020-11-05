/*
Single Test case
Statement
Create a Testing function to check the behavior of the following function. If the function returns a different value from the expected one, return an error specifying the test case. 
package main

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

Topics to Practice: 
testing, function, conditions, error

Multiple Test cases
Statement
Create a Testing function to check the behavior of the following function. Create a struct with the function parameter and the ‘want’ (expected) value, then iterate all the test cases and validate the behavior of the function against each case.
package main

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

Topics to Practice: 
testing, multiple cases
*/

package main

import (
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestSingleCase(t *testing.T) {
	got := IntMin(-2, 2)
	want := -2
	if got != want {
		t.Errorf("IntMin(%v, %v) = %v, want %v", -2, 2, got, want)
	}
}

func TestMultipleCases(t *testing.T) {
	type SampleTest struct {
		a, b, want int
	}
	tests := []SampleTest {
		{a: 2, b: -2, want: -2},
		{a: 1, b: -2, want: -2},
		{a: 1, b: 2, want: 1},
		{a: 0, b: 0, want: 0},
		{a: 100, b: 1000, want: 100},
		{a: 300, b: 200, want: 200},						
	}
	for _, tt := range tests {
		if got := IntMin(tt.a, tt.b); got != tt.want {
			t.Errorf("IntMin(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
		}
	}
}
