package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var RED_LIMIT = 12
var GREEN_LIMIT = 13
var BLUE_LIMIT = 14

func isLinePossible(line string) bool {
	r, g, b := 0, 0, 0

	for i := 0; i < len(line); i++ {
		if num, err := strconv.Atoi(string(line[i])); err == nil { //if char is a number
			if numleft, err := strconv.Atoi(string(line[i-1])); err == nil { //check for 2digit number
				num += numleft * 10
			}

			if string(line[i+2]) == "r" {
				r += num
			} else if string(line[i+2]) == "g" {
				g += num
			} else if string(line[i+2]) == "b" {
				b += num
			}

		} else if string(line[i]) == ";" && (r <= RED_LIMIT && g <= GREEN_LIMIT && b <= BLUE_LIMIT) {
			r, g, b = 0, 0, 0
		}
	}

	return (r <= RED_LIMIT && g <= GREEN_LIMIT && b <= BLUE_LIMIT)
}

func getLinePower(line string) int {
	r_max, g_max, b_max := 0, 0, 0

	for i := 2; i < len(line)-2; i++ {
		if num, err := strconv.Atoi(string(line[i])); err == nil { //if char is a number
			if numleft, err := strconv.Atoi(string(line[i-1])); err == nil { //check for 2digit number
				num += numleft * 10
			}

			if string(line[i+2]) == "r" && num > r_max {
				r_max = num
			} else if string(line[i+2]) == "g" && num > g_max {
				g_max = num
			} else if string(line[i+2]) == "b" && num > b_max {
				b_max = num
			}
		}
	}

	return r_max * g_max * b_max
}

func main() {
	sum := 0
	powerSum := 0
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if isLinePossible(line) {
			line = strings.TrimPrefix(line, "Game ")
			parts := strings.Split(line, ":")
			value, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				continue
			}
			sum += value
		}
		powerSum += getLinePower(line)
	}

	fmt.Println("part1", sum)
	fmt.Println("part2", powerSum)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
	}
}
