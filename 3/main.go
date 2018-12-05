package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	raw, _ := ioutil.ReadFile("input3")
	split:= strings.Split(string(raw), "\n")
	lines := split[:len(split)-1]

	var field map[int]map[int][]int = make(map[int]map[int][]int)
	var claims map[int]bool = make(map[int]bool)

	for _, l := range lines {
		id, x, y, w, h := parseLine(l)
		addClaim(field, claims, id, x, y, w, h)
	}

	s1 := countOverclaim(field)
	fmt.Println(s1)
	s2 := findNoOverlap(field, claims)
	fmt.Println(s2)
}

func addClaim(field map[int]map[int][]int, claims map[int]bool, id int, x int, y int, w int, h int) {
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			if _, ok := field[j]; !ok {
				field[j] = make(map[int][]int)
			}
			field[j][i] = append(field[j][i], id)
			claims[id] = true
		}
	}
}

func countOverclaim(field map[int]map[int][]int) int {
	c := 0
	for _, row := range field {
		for _, val := range row {
			if len(val) > 1 {
				c += 1
			}
		}
	}

	return c
}

func findNoOverlap(field map[int]map[int][]int, claims map[int]bool) int {
	for _, row := range field {
    		for _, claim := range row {
        		for _, c := range claim {
        			claims[c] = false
        		}
    		}
	}

	c := -1
	for claim, claimed := range claims {
    		if claimed {
        		c = claim
    		}
	}

	return c
}

func parseLine(s string) (int, int, int, int, int) {
	parts := strings.Split(s, " ")
	id, _ := strconv.Atoi(parts[0][1:len(parts[0])])
	coords := strings.Split(parts[2][0:len(parts[2])-1], ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])

	size := strings.Split(parts[3], "x")
	w, _ := strconv.Atoi(size[0])
	h, _ := strconv.Atoi(size[1])

	return id, x, y, w, h
}

