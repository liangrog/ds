package utils

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWait(t *testing.T) {
	assert.IsType(t, sync.WaitGroup{}, GetWait(10))
}
