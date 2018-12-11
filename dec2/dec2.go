package main

import (
	"bitbucket.org/sebsebmc/cryptopals/utils"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
//	part1()
	part2()
}


func part1(){
	text, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(text), "\n")
	twos, threes := 0, 0
	for line := range lines {
		chars := make(map[rune]int)
		for _, char := range lines[line] {
			chars[char] += 1
		}
		two, three := 0, 0
		for _, v := range chars {
			if v == 2 {
				two++
			}else if v == 3 {
				three++
			}
		}
		if two > 0 {
			fmt.Println(twos)
			twos++
		}
		if three > 0{
			fmt.Println(threes)
			threes++
		}
	}
	fmt.Println(twos*threes)
}

func part2(){
	text, _ := ioutil.ReadFile("input.txt")
	lines := bytes.Split(text, []byte("\n"))
	for line := range lines {
		for line2 := line; line2 < len(lines); line2++ {
			diff := utils.XorBytes(lines[line], lines[line2])
			num := 0
			idx := 0
			for j := range diff {
				if diff[j] != 0 {
					num++
					idx = j
				}
			}
			if num == 1 {
				fmt.Println(string(lines[line][:idx])+string(lines[line][idx+1:]))
			}
		}
	}
}