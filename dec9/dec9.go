package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	part1()
}

var marbles = list.New()
var scores []int
//478 players; last marble is worth 71240 points

func part1() {
	players := 478
	scores = make([]int, players)
	next := 1
	current := marbles.PushBack(0)
	for	i := next; i < 7124000; i++{
		if i %23 == 0 {
			scores[(i%players)] += i
			el := get7CounterClockwise(current)
			scores[(i%players)] += el.Value.(int)

			if el.Next() != nil {
				current = el.Next()
			}else {
				current = marbles.Front()
			}

			marbles.Remove(el)
			continue
		}
		nextEl := current.Next()
		if nextEl == nil {
			nextEl = marbles.Front()
		}

		current = marbles.InsertAfter(i, nextEl)
	}
	max, player := 0, -1
	for k, v := range scores {
		if v > max {
			max = v
			player = k
		}
	}
	fmt.Print(max, player)

}

func get7CounterClockwise(element *list.Element) *list.Element {
	for i:= 0; i < 7; i++ {
		if element.Prev() == nil {
			element = marbles.Back()
		}else {
			element = element.Prev()
		}
	}
	return element
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