package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	part2()
}

type Event struct {
	day, month, year int
	hour, minute     int
	action           string
}

type Events []Event

func (e Events) Len() int      { return len(e) }
func (e Events) Swap(i, j int) { e[i], e[j] = e[j], e[i] }
func (e Events) Less(i, j int) bool {
	if e[i].year != e[j].year {
		return e[i].year < e[j].year
	}
	if e[i].month != e[j].month {
		return e[i].month < e[j].month
	}
	if e[i].day != e[j].day {
		return e[i].day < e[j].day
	}
	if e[i].hour != e[j].hour {
		return e[i].hour < e[j].hour
	}
	if e[i].minute != e[j].minute {
		return e[i].minute < e[j].minute
	}
	return false
}

func part1() {
	lines := readFileLines()
	evs := make(Events, len(lines))
	for i := range lines {
		var (
			day, month, year int
			hour, minute     int
			action, action2  string
		)
		fmt.Sscanf(lines[i], "[%d-%d-%d %d:%d] %s%s", &year, &month, &day, &hour, &minute, &action, &action2)
		evs[i] = Event{day, month, year, hour, minute, action + action2}
	}
	sort.Sort(evs)
	sleepTime := make(map[int]int)
	cur := -1
	start := 0
	for i := 0; i < len(evs); i++ {
		fmt.Println(evs[i])
		switch evs[i].action[0] {
		case 'G':
			fmt.Sscanf(evs[i].action, "Guard#%d", &cur)
		case 'f':
			start = evs[i].minute
		case 'w':
			total := evs[i].minute - start
			sleepTime[cur] += total
		}
	}
	max, guard := 0, -1
	for k, v := range sleepTime {
		if v > max {
			max = v
			guard = k
		}
	}
	fmt.Println(guard)
	minutes := make(map[int]int)
	for i := 0; i < len(evs); i++ {
		switch evs[i].action[0] {
		case 'G':
			fmt.Sscanf(evs[i].action, "Guard#%d", &cur)
		case 'f':
			start = evs[i].minute
		case 'w':
			if cur == guard {
				for j := start; j < evs[i].minute; j++ {
					minutes[j]++
				}
			}
		}
	}
	max2, minute := 0, -1
	for k, v := range minutes {
		if v > max2 {
			max2 = v
			minute = k
		}
	}
	fmt.Println(minute, guard, minute*guard)
}

func part2()  {
	lines := readFileLines()
	evs := make(Events, len(lines))
	for i := range lines {
		var (
			day, month, year int
			hour, minute     int
			action, action2  string
		)
		fmt.Sscanf(lines[i], "[%d-%d-%d %d:%d] %s%s", &year, &month, &day, &hour, &minute, &action, &action2)
		evs[i] = Event{day, month, year, hour, minute, action + action2}
	}
	sort.Sort(evs)
	sleepTime := make(map[int]map[int]int)
	cur := -1
	start := 0
	for i := 0; i < len(evs); i++ {
		switch evs[i].action[0] {
		case 'G':
			fmt.Sscanf(evs[i].action, "Guard#%d", &cur)

			if sleepTime[cur] == nil {
				sleepTime[cur] = make(map[int]int)
			}
		case 'f':
			start = evs[i].minute
		case 'w':
			for j := start; j < evs[i].minute; j++ {
				sleepTime[cur][j]++
			}
		}
	}
	max, guard, minute := 0, -1, -1
	for g, v := range sleepTime {
		for m, c := range v {
			if c > max {
				max = c
				guard = g
				minute = m
			}
		}
	}
	fmt.Println(guard)

	fmt.Println(minute, guard, minute*guard)
}

func readFileLines() []string {
	text, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(text), "\n")
	return lines[:len(lines)-1]
}
