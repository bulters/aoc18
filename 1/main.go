package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	lines, _ := ioutil.ReadFile("input1")
	numbers := strings.Split(string(lines), "\n")
	var ns []int
	for _, e := range numbers {
		if e == "" {
			continue
		}
		n, _ := strconv.Atoi(e)
		ns = append(ns, n)
	}

	sum := 0
	for _, e := range ns {
		sum += e
	}
	fmt.Println(sum)

	// part 2

	var freqs map[int]int = make(map[int]int)
	var freq = 0
Search:
	for true {
		for _, e := range ns {
			freq += e

			if _, ok := freqs[freq]; ok {
				break Search
			}
			freqs[freq] = 1
		}
	}
	fmt.Println(freq)
}
