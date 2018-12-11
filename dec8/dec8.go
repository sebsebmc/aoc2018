package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type node struct {
	childCount, metaEntries int
	children []*node
	metadata []int
}

func main() {
	part1()
}

var nums []int
var metaSum = 0

func part1() {
	lines := readFileLines()
	line := lines[0]
	strs := strings.Split(line, " ")
	nums = make([]int, len(strs))
	for str := range strs {
		num, err := strconv.Atoi(strs[str])
		if err != nil {
			fmt.Println(err)
		}
		nums[str] = num
	}
	root := node{}
	cur := &root
	idx := 0
	buildNodes(cur, idx)
	part2(&root)
}

func buildNodes(cur *node, idx int) int {
	cur.childCount = nums[idx]
	cur.children = make([]*node, cur.childCount)
	cur.metaEntries = nums[idx+1]
	cur.metadata = make([]int, cur.metaEntries)
	if cur.childCount > 0 {
		idx += 2
		for i := 0; i < cur.childCount; i++ {
			cur.children[i] = new(node)
			idx = buildNodes(cur.children[i], idx)
		}
	}else {
		idx += 2
	}
	copy(cur.metadata, nums[idx:idx+cur.metaEntries])
	for i := 0; i < cur.metaEntries; i++ {
		metaSum += nums[idx+i]
	}
	return idx+cur.metaEntries
}

func part2(root *node) {
	fmt.Println(evalNode(root))
}

func evalNode(cur *node) int {
	val := 0
	if cur.childCount == 0 {
		for i := 0; i < cur.metaEntries; i++ {
			val += cur.metadata[i]
		}
		return val
	}
	for i := 0; i < cur.metaEntries; i++ {
		if cur.metadata[i]-1 < cur.childCount {
			val += evalNode(cur.children[cur.metadata[i]-1])
		}
	}
	return val
}

func readFileLines() []string {
	text, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(text), "\n")
	return lines[:len(lines)-1]
}
