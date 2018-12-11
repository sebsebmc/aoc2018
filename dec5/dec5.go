package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

func main() {
	//part1(readFileLines()[0])
	part2()
}

func part1(line string) int {
	//fmt.Println(len(line))
	for {
		//fmt.Println(line)
		changed := false
		var temp strings.Builder
		i := 0
		for ; i < len(line)-1; i++ {
			//fmt.Println(i)
			if line[i] != line[i+1] && unicode.ToLower(rune(line[i])) == unicode.ToLower(rune(line[i+1])) {
				changed = true
				i++
			} else {
				temp.WriteByte(line[i])
			}
		}
		if i == len(line)-1 {
			temp.WriteByte(line[i])
		}
		line = temp.String()
		//fmt.Println(len(line))
		if !changed {
			break
		}
	}
	return len(line)
}

func test() {
	line := `aAbBcCdDeE`
	fmt.Println(len(line))
	for {
		//fmt.Println(line)
		changed := false
		var temp strings.Builder
		for i := 0; i < len(line)-1; i++ {
			//fmt.Println(i)
			if line[i] != line[i+1] && unicode.ToLower(rune(line[i])) == unicode.ToLower(rune(line[i+1])) {
				changed = true
				i++
			} else {
				temp.WriteByte(line[i])
			}
		}
		temp.WriteByte(line[len(line)-1])
		line = temp.String()
		//fmt.Println(len(line))
		if !changed {
			break
		}
	}
	fmt.Println(len(line))
}

func part2(){
	line := readFileLines()[0]
	min := len(line)
	best := 0
	for i := 0; i < 26; i++ {
		var sb strings.Builder
		sb.Reset()
		for j := 0; j < len(line); j++ {
			if line[j] == byte(65+i) || line[j] == byte(97+i) {
				continue
			}
			sb.WriteByte(line[j])
		}
		res := part1(sb.String())
		fmt.Println(i, res)
		if res < min {
			min = res
			best = i
		}
	}
	fmt.Println(best, min)
}

func readFileLines() []string {
	text, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(text), "\n")
	return lines[:len(lines)-1]
}
