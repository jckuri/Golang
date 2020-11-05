/*
Different Asserts
Statement
Using Testify/assert pkg, execute different functions and validate the expected result. You can create as many cases as you want
Equal | NotEqual | Nil | NotNil 

Topics to Practice: 
testing, multiple cases, testify
*/

/*
Installation:
sudo apt install git
go get github.com/stretchr/testify
*/

package my_package

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
  assert := assert.New(t)

  // assert equality
  type SquareTest struct {
  	x, want float64
  }
  tests := []SquareTest {
  	SquareTest {2, 4},
  	SquareTest {3, 9}, 
  	SquareTest {4, 16},  
  	SquareTest {-2, 4},
  	SquareTest {-3, 9}, 
  	SquareTest {-4, 16},  
  }
  for _, test := range tests {
  	assert.Equal(Square(test.x), test.want, "they should be equal")
  }
  

  // assert inequality
  tests2 := []float64 {2, 3, 4, -2, -3, -4,}
  for _, test := range tests2 {
  	assert.NotEqual(test, Add64(test), "they should not be equal")
  }

  tests3 := []float64 {2, 3, 0, 4, -2, -3, 0, -4,}
  for _, test := range tests3 {
  	result := HundredDividedBy(test)
  	if test == 0 {
  		// assert for nil (good for errors)
		  assert.Nil(result)
  	} else {
		// assert for not nil (good when you expect something)
		if assert.NotNil(result) {
			assert.Equal(result.number, 100. / test)
		}
	}
  }

}

/*
OUTPUT:

$ go test
PASS
ok  	_/media/jckuri/1.9TB/TRABAJOS/BairesDev/Golang_Development_Program/Level3	0.005s

*/
