package adventofcode

import (
	"fmt"
	"runtime"

	"github.com/dustin/go-humanize"
)

func PrintMemoryUsage() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	fmt.Printf("Allocated memory: %v\n", humanize.Bytes(mem.Alloc))
	fmt.Printf("Total memory allocated (since start): %v\n", humanize.Bytes(mem.TotalAlloc))
}
