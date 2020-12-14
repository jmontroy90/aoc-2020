package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type passwordCheck struct {
	min, max int
	token string
	password string
}

func (c *passwordCheck) isValidOne() bool {
	var tokenCount int
	for _, r := range c.password {
		if string(r) == c.token {
			tokenCount += 1
		}
	}
	if c.min <= tokenCount && tokenCount <= c.max {
		return true
	}
	return false
}

// this is an XOR, but don't be fancy about it
func (c *passwordCheck) isValidTwo() bool {
	atOne := string(c.password[c.min-1]) == c.token // off by one
	atTwo := string(c.password[c.max-1]) == c.token // off by one

	switch {
	case (atOne && !atTwo) || (!atOne && atTwo):
		return true
	default:
		return false
	}
}

func main() {
	tests, err := load()
	if err != nil {
		log.Fatalf("fatal error loading file: %v", err)
	}

	var passingOne, passingTwo int
	for _, l := range tests {
		if l.isValidOne() {
			passingOne += 1
		}
		if l.isValidTwo() {
			passingTwo += 1
		}
	}

	fmt.Printf("passing one: %d\npassing two: %d", passingOne, passingTwo)
}

func load() ([]passwordCheck, error) {
	inputFile, err := os.Open("day2/day2.txt")
	if err != nil {
		return nil, err
	}
	input, err := ioutil.ReadAll(inputFile)
	if err != nil {
		return nil, err
	}
	unparsed := strings.Split(string(input), "\n")
	tests := make([]passwordCheck, len(unparsed))
	for i, r := range unparsed {
		s := strings.Split(r, ": ")
		l1 := strings.Split(s[0], " ")
		l2 := strings.Split(l1[0], "-")
		tests[i] = passwordCheck{min: mustInt(l2[0]), max: mustInt(l2[1]), token: l1[1], password: s[1]}
	}
	return tests, nil
}

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("can't convert to int: %s", s)
	}
	return i
}