package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"log"
)



func main() {
	fast()
}

func slow()  {
	text, _ := ioutil.ReadFile("input.txt")
	sum := 0
	lines := strings.Split(string(text), "\n")
	sums := make([]int, 1)
	for i := 0; true; i++ {
		n, err := strconv.Atoi(lines[i%len(lines)])
		if err != nil {
			log.Fatalf("%v", err)
		}
		sum += n
		for j := 0; j <= i; j++ {
			if sums[j] == sum {
				fmt.Println(sum)
				return
			}
		}
		sums = append(sums, sum)
	}
}

func fast(){
		text, _ := ioutil.ReadFile("input.txt")
		sum := 0
		lines := strings.Split(string(text), "\n")
		sums := make(map[int]bool)
		for i := 0; true; i++ {
			n, err := strconv.Atoi(lines[i%len(lines)])
			if err != nil {
				log.Fatalf("%v", err)
			}
			sum += n
			if _, ok := sums[sum]; ok {
				fmt.Println(sum)
				return
			}
			sums[sum] = true
		}
}