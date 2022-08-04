package main

import (
	"math"
	"math/rand"
	"time"
)

func computeVote(candSlice []Coordinate, generator *rand.Rand) int {
	indiv := Coordinate{generator.Float64(), generator.Float64()}
	var curWinner int = 0
	var curMinDist float64 = math.MaxFloat64
	for j := range candSlice {
		candDist := distance(indiv, (candSlice)[j])
		if candDist < curMinDist {
			curWinner = j
			curMinDist = candDist
		}
	}
	return curWinner
}

func worker(id int, jobs <-chan int, results chan []int, candSlice []Coordinate) {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)

	for popNum := range jobs {
		updateSlice := make([]int, CANDIDATES_NUM)
		for i := 0; i < popNum; i += 1 {
			winner := computeVote(candSlice, generator)
			updateSlice[winner] += 1
		}
		results <- updateSlice
	}
}

func pluralitySimulatorConcurrentWorkerPool(candSlice []Coordinate) []int {
	defer timer()()

	numJobs := WORKER_NUM * NUM_JOBS_PER_WORKER
	popPerJob := POPULATION_NUM / numJobs
	jobs := make(chan int, 1000)
	results := make(chan []int, numJobs)
	ret := make([]int, CANDIDATES_NUM)

	for w := 0; w < WORKER_NUM; w += 1 {
		go worker(w, jobs, results, candSlice)
	}

	for j := 1; j <= numJobs; j += 1 {
		if j == numJobs {
			jobs <- popPerJob + (POPULATION_NUM % numJobs)
		} else {
			jobs <- popPerJob
		}
	}
	close(jobs)

	for r := 0; r < numJobs; r += 1 {
		updateSlice := <-results
		for i := 0; i < CANDIDATES_NUM; i += 1 {
			ret[i] += updateSlice[i]
		}
	}
	return ret
}

func pluralitySimulatorSingle(candSlice []Coordinate) []int64 {
	defer timer()()
	countSlice := make([]int64, CANDIDATES_NUM)
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	for i := 0; i < POPULATION_NUM; i += 1 {
		curWinner := computeVote(candSlice, generator)
		countSlice[curWinner] += 1
	}
	return countSlice
}
