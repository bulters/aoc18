package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	raw, _ := ioutil.ReadFile("input2")
	split := strings.Split(string(raw), "\n")
	lines := split[:len(split)-1]
	twos := 0
	threes := 0

	for _, e := range lines {
		histo := histo(e)
		if hasCount(histo, 2) {
			twos += 1
		}

		if hasCount(histo, 3) {
			threes += 1
		}
	}

	checksum := twos * threes
	fmt.Println(checksum)

	// part 2
	var match1 string
	var match2 string
Search:
	for i, e := range lines {
		for _, o := range lines[i+1:] {
			if d := levenshtein(e, o); d == 1 {
				match1 = e
				match2 = o
				break Search
			}

		}
	}

	o := overlap(match1, match2)
	fmt.Println(match1)
	fmt.Println(match2)
	fmt.Println(o)
}

func histo(w string) map[rune]int {
	histo := make(map[rune]int)
	for _, e := range w {
		if _, ok := histo[e]; ok {
			histo[e] += 1
		} else {
			histo[e] = 1
		}
	}

	return histo

}

func hasCount(hs map[rune]int, n int) bool {
	for _, e := range hs {
		if e == n {
			return true
		}
	}

	return false
}

func minimum3(a, b, c int) int {
	min := a
	if min > b {
		min = b
	}
	if min > c {
		min = c
	}
	return min
}

func levenshtein(s1, s2 string) int {
	s1r := []rune(s1)
	s1l := len(s1r) + 1
	s2r := []rune(s2)
	s2l := len(s2r) + 1

	d := make([][]int, s1l)
	for i := range d {
		d[i] = make([]int, s2l)
	}

	for i := 0; i < s1l; i++ {
		d[i][0] = i
	}

	for j := 0; j < s2l; j++ {
		d[0][j] = j
	}

	for j := 1; j < s2l; j++ {
		for i := 1; i < s1l; i++ {
			sc := 0
			if s1r[i-1] != s2r[j-1] {
				sc = 1
			}
			d[i][j] = minimum3(
				d[i-1][j]+1,
				d[i][j-1]+1,
				d[i-1][j-1]+sc)
		}
	}

	return d[s1l-1][s2l-1]
}

func overlap(s1, s2 string) string {
	var s string
	for i, c := range s1 {
		if c == rune(s2[i]) {
			s = s + string(c)
		}
	}
	return s
}
