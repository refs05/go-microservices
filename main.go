package main

import "time"

func main() {
	start := time.Now()
	println("Hello, World!")
	elapsed := time.Since(start).Nanoseconds()
	println("Execution time:", elapsed)
}