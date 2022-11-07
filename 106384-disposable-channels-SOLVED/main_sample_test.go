package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSampleSolution(t *testing.T) {
	a := make(chan string, 1)
	b := make(chan string, 1)
	Solution(2, "ali", a, b)
	assert.Equal(t, "ali", <-a)
	assert.Equal(t, "ali", <-b)
}
