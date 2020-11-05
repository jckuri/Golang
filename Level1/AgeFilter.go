/*
Age Filter
Statement
Create a function that will filter a Slice of ages that are between the range. The function will receive two numbers and a slice of ages as parameters. It should return the ages between the range

Topics to Practice:
slice, append, function, for loop, control flow, return value
*/

package main

import (
	"fmt"
)

func AgeFilter(minAge, maxAge int, ages []int) []int {
	var filteredAges []int
	for _, age := range ages {
		if age >= minAge && age <= maxAge {
			filteredAges = append(filteredAges, age)
		}
	}
	return filteredAges
}

func main() {
	ages := []int {10, 12, 13, 6, 5, 4, 90, 91, 93, 81, 73, 62, 54, 45, 32, 28, 22, 18, 19}
	fmt.Println("Ages:", ages)
	filteredAges := AgeFilter(18, 60, ages)
	fmt.Println("Filtered Ages:", filteredAges)
}

/*
OUTPUT:

Ages: [10 12 13 6 5 4 90 91 93 81 73 62 54 45 32 28 22 18 19]
Filtered Ages: [54 45 32 28 22 18 19]
*/
