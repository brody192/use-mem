package main

import (
	"fmt"
	"math"
	"os"
	"runtime/debug"
	"strconv"
	"time"
)

func init() {
	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(math.MaxInt64)
}

func main() {
	var useMem = os.Getenv("USE_MEM")
	if useMem == "" {
		fmt.Println("environment variable USE_MEM is missing")
		os.Exit(0)
	}

	mb, err := strconv.Atoi(useMem)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	var offset = 15 // corrects for reported ram usage on railway

	var array = make([]byte, (mb-offset)*1000000)

	fmt.Println("allocating array of " + useMem + " megabytes")

	for i := range array {
		array[i] = 0
	}

	fmt.Println("sleeping for infinite years")

	for {
		time.Sleep(time.Duration(1<<63 - 1))
	}
}
