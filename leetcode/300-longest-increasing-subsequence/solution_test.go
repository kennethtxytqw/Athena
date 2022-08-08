package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in  []int
	out int
}{

	{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
	{[]int{0, 1, 0, 3, 2, 3}, 4},
	{[]int{7, 7, 7, 7, 7, 7, 7}, 1},
}

func TestSolution1(t *testing.T) {
	assert := assert.New(t)

	for _, test := range tests {
		assert.Equal(test.out, solution1(test.in), fmt.Sprintf("in: %v\n", test.in))
	}
}