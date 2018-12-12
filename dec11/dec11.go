package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func main() {
	part1()

}

func part1() {
	input := 8868
	grid := make([][]int, 300)
	for i := 0; i < 300; i++ {
		grid[i] = make([]int, 300)
		for j := 0; j < 300; j++ {
			rackId := (j + 1) + 10
			powerLevel := rackId * (i + 1)
			powerLevel += input
			powerLevel *= rackId
			powerLevel = (powerLevel % 1000) / 100
			powerLevel -= 5
			grid[i][j] = powerLevel
		}
	}
	sums := make([]int, 300)
	s := 3
	fmt.Println(s)
	for i := 0; i < 300-s; i++ {
		for j := 0; j < 300-s; j++ {
			for x := 0; x < s; x++ {
				for y := 0; y < s; y++ {
					sums[i*300+j] += grid[i+x][j+y]
				}
			}
		}
	}
	max, idx := findMax(sums)
	fmt.Println(max, idx, (idx%300)+1, (idx/300)+1)
}

func part2() {
	input := 8868
	grid := make([][]int, 300)
	for i := 0; i < 300; i++ {
		grid[i] = make([]int, 300)
		for j := 0; j < 300; j++ {
			rackId := (j + 1) + 10
			powerLevel := rackId * (i + 1)
			powerLevel += input
			powerLevel *= rackId
			powerLevel = (powerLevel % 1000) / 100
			powerLevel -= 5
			grid[i][j] = powerLevel
		}
	}
	sums := make([][]int, 300)

	for s := 1; s < 300; s++ {
		sums[s] = make([]int, 300*300)
		fmt.Println(s)
		for i := 0; i < 300-s; i++ {
			for j := 0; j < 300-s; j++ {
				for x := 0; x < s; x++ {
					for y := 0; y < s; y++ {
						sums[s][i*300+j] += grid[i+x][j+y]
					}
				}
			}
		}
	}
	maxes, idxs := make([]int, 300), make([]int, 300)
	for i := 0; i < len(sums); i++ {
		maxes[i], idxs[i] = findMax(sums[i])
	}
	max, idx := findMax(maxes)
	fmt.Println(idx)
	idx = idxs[idx]
	fmt.Println(max, idx, (idx%300)+1, (idx/300)+1)
}

func findMin(nums []int) (int, int) {
	min, idx := math.MaxInt32, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
			idx = i
		}
	}
	return min, idx
}

func findMax(nums []int) (int, int) {
	max, idx := math.MinInt32, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			idx = i
		}
	}
	return max, idx
}

func readFileLines() []string {
	text, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("Cannot read file: %v", err)
	}
	lines := strings.Split(string(text), "\n")
	return lines[:len(lines)-1]
}
