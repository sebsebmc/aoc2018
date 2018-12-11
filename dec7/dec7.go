package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	part1()
}

func part1() {

}

func part2() {

}

func readFileLines() []string {
	text, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(text), "\n")
	return lines[:len(lines)-1]
}