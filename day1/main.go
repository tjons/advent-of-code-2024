package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	left := make([]int, 0, 1000)
	right := make([]int, 0, 1000)

	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("error reading input")
		os.Exit(1)
	}

	s := bufio.NewScanner(input)
	for s.Scan() {
		line := s.Text()
		nums := strings.Split(line, "   ")

		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])

		left = append(left, n1)
		right = append(right, n2)
	}

	slices.Sort(left)
	slices.Sort(right)

	distance := 0.0
	for i := 0; i < len(left); i++ {
		distance += math.Abs(float64(left[i] - right[i]))
	}

	fmt.Printf("The answer to part1 is %v\n", int(distance))

	rightScores := make(map[int]int)
	for _, val := range right {
		if _, ok := rightScores[val]; !ok {
			rightScores[val] = 1
		} else {
			rightScores[val]++
		}
	}

	similarity := 0
	for _, val := range left {
		score, ok := rightScores[val]
		if !ok {
			continue
		}

		similarity += val * score
	}

	fmt.Printf("the answer to part2 is %v\n", similarity)
}
