package goUtils

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"time"
)

func GetTimestampNano() int64 {
	return time.Now().UnixNano()
}

// getDT
func GetDT() string {
	// time date formatting...
	// https://golang.org/src/time/format.go
	return time.Now().Format("2006-01-02 15:04:05.0000")
}

// printMemUsage
func PrintMemUsage(context string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("MiB usage malloc %v  tot malloc %v  Sys %v  GC %v  (%s)\n",
		b2KiB(m.Alloc), b2KiB(m.TotalAlloc), b2KiB(m.Sys), m.NumGC, context)
}

func GetMemAlloc() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc
}

func GetMemAllocTotal() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.TotalAlloc
}

// b2KiB
func Bit2KiB(b uint64) uint64 {
	return b / 1024
}

// readFileContent .
func ReadFileContent(filename string) string {
	var content []byte
	var err error
	if FileExists(filename) {
		content, err = ioutil.ReadFile(filename)
		if err != nil {
			fmt.Printf("Unable to find file to load %s\n", filename)
			return "{}"
		}
	}
	return string(content)
}

// fileExists
func FileExists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	}

	return false
}
