package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleSolution(t *testing.T) {
	f := func(inp uint8) uint8 { return inp }
	inp := uint32(1 + 2<<8 + 3<<16 + 4<<24)
	res := Solution(f, inp)
	assert.Equal(t, inp, res)
}
