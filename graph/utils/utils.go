package utils

import (
	"sync"
)

const (
	WAITMAX = 1000
)

// Get wait group for given
// worker munber. Maximum worker
// number is limited by const
// WAITMAX.
func GetWait(num int) sync.WaitGroup {
	var wg sync.WaitGroup

	if num <= WAITMAX {
		wg.Add(num)
	} else {
		wg.Add(WAITMAX)
	}

	return wg
}
