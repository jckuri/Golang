/*
Create many test cases as needed to validate the right behavior of the Age Filter exercise from Level 1
package main
import (
	“fmt”
)
func filterByAgeRange(fromAge, toAge int, ages []int) []int{
	response := make([]int, 0)
	for _,ageToCheck := range ages {
		if fromAge <= ageToCheck && ageToCheck <= toAge  {
			response = append(response, ageToCheck)
		}
	}
	return response 
}
*/

package main

import (
	"testing"
)

func filterByAgeRange(fromAge, toAge int, ages []int) []int{
	response := make([]int, 0)
	for _,ageToCheck := range ages {
		if fromAge <= ageToCheck && ageToCheck <= toAge  {
			response = append(response, ageToCheck)
		}
	}
	return response 
}

type Test struct {
	fromAge, toAge int
	ages []int
	wanted []int
}

func AreSlicesEqual(l1, l2 []int) bool {
	if len(l1) != len(l2) {return false}
	for i := range l1 {
		if l1[i] != l2[i] {return false}
	}
	return true
}

func TestFilterByAgeRange(t *testing.T) {
	tests := []Test {
		Test{20, 30, []int {19, 20, 30, 31}, []int {20, 30}}, 
		Test{5, 20, []int{1,5,20, 10, 17, 18, 65,  66, 100, 50, 20}, []int {5, 20, 10, 17, 18, 20}},
	}
	for _, test := range tests {
		//fmt.Println(test)
		result := filterByAgeRange(test.fromAge, test.toAge, test.ages)
		if !AreSlicesEqual(test.wanted, result) {
			t.Errorf("ERROR: Result: %v, Wanted: %v\n", result, test.wanted)
		}
		//fmt.Println(test.wanted, result)
	}
}

/*
OUTPUT:

=== RUN   TestFilterByAgeRange
--- PASS: TestFilterByAgeRange (0.00s)
PASS
*/
