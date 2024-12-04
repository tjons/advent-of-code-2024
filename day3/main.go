package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	part1()
	part2()
}

const MulToken = "mul("
const EnabledMulToken = "do()"
const DisableMulToken = "don't()"

func part1() {
	multiplications := make([]*multiplication, 0)

	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("error reading input")
		os.Exit(1)
	}

	s := bufio.NewScanner(input)
	s.Split(bufio.ScanRunes)

	last4Char := ""
	inMul := false
	mulClosed := false
	for s.Scan() {
		if len(last4Char) < 4 {
			last4Char += s.Text()
		} else {
			last4Char = string([]byte(last4Char)[1:4]) + s.Text()
		}

		if last4Char == MulToken {
			inMul = true

			var firstNum, secondNum []rune
			curNum := &firstNum
			for s.Scan() {
				if isDigit(rune(s.Text()[0])) {
					n := append(*curNum, rune(s.Text()[0]))
					curNum = &n
				} else if s.Text() == "," {
					firstNum = make([]rune, len(*curNum))
					copy(firstNum, *curNum)
					curNum = &secondNum
				} else if s.Text() == ")" {
					mulClosed = true
					break
				} else {
					break
				}
			}
			secondNum = make([]rune, len(*curNum))
			copy(secondNum, *curNum)

			if inMul && mulClosed {
				n, err := strconv.Atoi(string(firstNum))
				if err != nil {
					continue
				}

				n2, err := strconv.Atoi(string(secondNum))
				if err != nil {
					continue
				}

				multiplications = append(multiplications, &multiplication{
					right: n,
					left:  n2,
				})
			}
			inMul = false
			mulClosed = false
		}
	}

	total := 0
	for _, calc := range multiplications {
		total += calc.calculate()
	}

	fmt.Printf("the result is %d\n", total)
}

func part2() {
	mulEnabled := true
	multiplications := make([]*multiplication, 0)

	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("error reading input")
		os.Exit(1)
	}

	s := bufio.NewScanner(input)
	s.Split(bufio.ScanRunes)

	last4Char := ""
	last7Char := ""
	inMul := false
	mulClosed := false
	for s.Scan() {
		if len(last4Char) < 4 {
			last4Char += s.Text()
		} else {
			last4Char = string([]byte(last4Char)[1:4]) + s.Text()
		}

		if len(last7Char) < 7 {
			last7Char += s.Text()
		} else {
			last7Char = string([]byte(last7Char)[1:7]) + s.Text()
		}

		if last7Char == DisableMulToken {
			mulEnabled = false
		}

		if last4Char == EnabledMulToken {
			mulEnabled = true
		}

		if mulEnabled && last4Char == MulToken {
			inMul = true

			var firstNum, secondNum []rune
			curNum := &firstNum
			for s.Scan() {
				if len(last4Char) < 4 {
					last4Char += s.Text()
				} else {
					last4Char = string([]byte(last4Char)[1:4]) + s.Text()
				}

				if len(last7Char) < 7 {
					last7Char += s.Text()
				} else {
					last7Char = string([]byte(last7Char)[1:7]) + s.Text()
				}

				if isDigit(rune(s.Text()[0])) {
					n := append(*curNum, rune(s.Text()[0]))
					curNum = &n
				} else if s.Text() == "," {
					firstNum = make([]rune, len(*curNum))
					copy(firstNum, *curNum)
					curNum = &secondNum
				} else if s.Text() == ")" {
					mulClosed = true
					break
				} else {
					break
				}
			}
			secondNum = make([]rune, len(*curNum))
			copy(secondNum, *curNum)

			if inMul && mulClosed {
				n, err := strconv.Atoi(string(firstNum))
				if err != nil {
					continue
				}

				n2, err := strconv.Atoi(string(secondNum))
				if err != nil {
					continue
				}

				multiplications = append(multiplications, &multiplication{
					right: n,
					left:  n2,
				})
			}
			inMul = false
			mulClosed = false
		}
	}

	total := 0
	for _, calc := range multiplications {
		total += calc.calculate()
	}

	fmt.Printf("the result is %d\n", total)
}

func isDigit(s rune) bool {
	return s == '1' ||
		s == '2' ||
		s == '3' ||
		s == '4' ||
		s == '5' ||
		s == '6' ||
		s == '7' ||
		s == '8' ||
		s == '9' ||
		s == '0'
}

type multiplication struct {
	right, left int
}

func (m multiplication) calculate() int {
	return m.right * m.left
}
