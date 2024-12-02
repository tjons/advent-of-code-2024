package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	day2()
}

func day2() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("error reading input")
		os.Exit(1)
	}

	s := bufio.NewScanner(input)

	safeCount := 0            // track how many lines are safe
	potentiallySafeCount := 0 // track how many lines only have one number making the line unsafe
	// track how many lines have two or more numbers making the line unsafe. these will
	// have to be reanalyzed with removal of each iteam.
	unsafeLines := [][]string{}

	for s.Scan() {
		line := s.Text()
		nums := strings.Split(line, " ")
		safe, safetyScore := lineIsSafe(nums, -1)
		if safe {
			safeCount++
		} else if safetyScore == 1 {
			potentiallySafeCount++
		} else {
			unsafeLines = append(unsafeLines, nums)
		}
	}

	fmt.Printf("%d reports are safe\n", safeCount)

	dangerousButSafeCount := 0
	for _, nums := range unsafeLines {
		for i := 0; i < len(nums); i++ {
			safe, _ := lineIsSafe(nums, i)
			if safe {
				dangerousButSafeCount++
				break
			}
		}
	}
	fmt.Printf("with Problem Dampener, %d reports are safe\n", potentiallySafeCount+safeCount+dangerousButSafeCount)
}

// lineIsSafe takes a slice of strings to represent the line, split on whitespace.
// skipIdx refers to a given number to skip, i.e., process the line as if that item
// was not in nums. To not skip any numbers, use -1.
// the bool return value will represent whether the line is safe, and the int
// value will describe how many numbers made it unsafe.
func lineIsSafe(nums []string, skipIdx int) (bool, int) {
	unsafeScore := 0
	start := 0
	prevNum := 0

	if skipIdx == 0 {
		prevNum, _ = strconv.Atoi(nums[1])
		start = 2
	} else {
		prevNum, _ = strconv.Atoi(nums[0])
		start = 1
	}

	lineDirection := 0

	for i := start; i < len(nums); i++ {
		if i == skipIdx {
			continue
		}

		num, _ := strconv.Atoi(nums[i])
		if lineDirection == 0 {
			if num < prevNum {
				lineDirection = -1
			} else if num > prevNum {
				lineDirection = 1
			} else {
				unsafeScore++
				continue
			}
		}

		if lineDirection == -1 {
			if num > prevNum || num == prevNum || prevNum-num > 3 {
				unsafeScore++
				continue
			}
			prevNum = num
		} else {
			if num < prevNum || num == prevNum || num-prevNum > 3 {
				unsafeScore++
				continue
			}
			prevNum = num
		}
	}

	return unsafeScore == 0, unsafeScore
}
