package main

import (
	"fmt"
	"runtime"
)

func main() {

	// NumCPU returns the number of logical
	// Cores available for the current process.
	fmt.Println(runtime.NumCPU())
}
