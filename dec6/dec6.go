package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	part1()
}

func part1() {

}

func part2() {

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
	text, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(text), "\n")
	return lines[:len(lines)-1]
}
