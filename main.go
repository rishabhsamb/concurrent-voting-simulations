package main

import (
	"fmt"
	"math/rand"
	"time"
)

var WORKER_NUM = 8
var NUM_JOBS_PER_WORKER = 1
var CANDIDATES_NUM = 2
var POPULATION_NUM = 10000000000

func timer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Execution took %v\n", time.Since(start))
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	candidatesSlice := []Coordinate{
		{rand.Float64(), rand.Float64()},
		{rand.Float64(), rand.Float64()},
	}
	fmt.Println(candidatesSlice)

	fmt.Println(pluralitySimulatorSingle(candidatesSlice))

	fmt.Println(pluralitySimulatorConcurrentWorkerPool(candidatesSlice))

}
