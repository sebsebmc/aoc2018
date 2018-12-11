package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	part1()
}

type Claim struct {
	id int
	x,y int
	width, height int
}

func part2() {
	lines := readFileLines()
	var (
		id int
		x,y int
		width, height int
	)
	claims := make([]Claim, len(lines))
	for i := range lines {
		_, err := fmt.Sscanf(lines[i], "#%d @ %d,%d:%dx%d", &id, &x, &y, &width, &height)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		claims[i] = Claim{id, x, y, width, height}
	}
	fabric := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		fabric[i] = make([]int, 1000)
	}
	for i := range claims {
		hasOverlap := false
		for j := 0; j < len(claims); j++ {
			if i == j {
				continue
			}
			if !(claims[i].x > (claims[j].x+claims[j].width-1) ||
				(claims[i].x+claims[i].width-1) < claims[j].x ||
				claims[i].y > (claims[j].y+claims[j].height-1) ||
				(claims[i].y+claims[i].height-1) < claims[j].y) {
				hasOverlap = true
				break
			}
		}
		if !hasOverlap {
			fmt.Println(claims[i].id)
		}
	}
}

func part1() {
	lines := readFileLines()
	var (
		id int
		x,y int
		width, height int
	)
	claims := make([]Claim, len(lines))
	for i := range lines {
		_, err := fmt.Sscanf(lines[i], "#%d @ %d,%d:%dx%d", &id, &x, &y, &width, &height)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		claims[i] = Claim{id, x, y, width, height}
	}
	fabric := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		fabric[i] = make([]int, 1000)
	}
	for i := range claims {
		claim := claims[i]
		for j := 0; j < claim.width; j++ {
			for k := 0; k < claim.height; k++ {
				fabric[j+claim.x][k+claim.y]++
			}
		}
	}
	overlap := 0
	for j := 0; j < 1000; j++ {
		for k := 0; k < 1000; k++ {
			if fabric[j][k] > 1 {
				overlap++
			}
		}
	}
	fmt.Println(overlap)
}

func readFileLines() []string {
	text, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(text), "\n")
	return lines[:len(lines)-1]
}