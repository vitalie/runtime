package setup

import (
	"math/rand"
	"runtime"
	"time"
)

func init() {
	// Seed randomines.
	rand.Seed(time.Now().UnixNano())

	// Adjust SO threads.
	runtime.GOMAXPROCS(runtime.NumCPU())
}
