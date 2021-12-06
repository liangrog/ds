package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsOddNumber(t *testing.T) {
	odd := []int{223, 777, 99}
	for _, n := range odd {
		assert.True(t, IsOddNumber(n))
	}

	even := []int{224, 0, 1000}
	for _, n := range even {
		assert.False(t, IsOddNumber(n))
	}

}
