package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Point struct {
	x, y       int
	velX, velY int
}

func main() {
	part1()
}
//position=<-42346,  10806> velocity=< 4, -1>
func part1() {
	lines := readFileLines()
	points := make([]Point, len(lines))
	var (
		x, y       int
		velX, velY int
	)
	for i := 0; i < len(lines); i++ {
		fmt.Sscanf(lines[i],"position=<%d, %d> velocity=<%d, %d>", &x, &y, &velX, &velY)
		points[i] = Point{x,y, velX, velY}
	}
	minDist := 10000000
	var maxX, maxY, minX, minY, min int
	for i := 0; i < 10639; i++{
		step(points)
		min, maxX, maxY, minX, minY = calcDist(points)
		if min < minDist {
			minDist = min
			//iter = i
		}
	}
	fmt.Println(maxX, maxY, minX, minY)
	render(points, maxX, maxY, minX, minY)
}

func render(points []Point, maxX, maxY, minX, minY int){
	lines := make([][]int, maxY - minY+1)
	for i := minY; i <= maxY; i++ {
		lines[i-minY] = make([]int, maxX - minX+1)
	}
	for i := 0; i < len(points); i++ {
		lines[points[i].y-minY][points[i].x-minX] = 8
	}
	for i := 0; i < len(lines); i++  {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] != 8 {
				fmt.Print(" ")
			}else {
				fmt.Print("8")
			}
		}
		fmt.Println()
	}
}

func step(points []Point){
	for i := 0; i<len(points);	i++  {
		points[i].x += points[i].velX
		points[i].y += points[i].velY
	}
}

func calcDist(points []Point) (int, int, int, int, int) {
	minX, minY := 100000, 100000
	maxX, maxY := -100000, -100000
	for i := 0; i < len(points); i++ {
		if points[i].x < minX {
			minX = points[i].x
		}
		if points[i].y < minY {
			minY = points[i].y
		}
		if points[i].x > maxX {
			maxX = points[i].x
		}
		if points[i].y > maxY {
			maxY = points[i].y
		}
	}
	totalDist := (maxX - minX) + (maxY - minY)
	//fmt.Println(minX, minY, maxX, maxY)
	return totalDist, maxX, maxY, minX, minY
}

func readFileLines() []string {
	text, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("Cannot read file: %v", err)
	}
	lines := strings.Split(string(text), "\n")
	return lines[:len(lines)-1]
}
