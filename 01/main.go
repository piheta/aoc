package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func buildNumLine(line string) string {
	numLine := ""
	for i := 0; i < len(line); i++ {
		if _, err := strconv.Atoi(string(line[i])); err == nil {
			numLine += string(line[i])
		}
	}
	return numLine
}

func convertLettersToNumber(line string) string { //🍝🍝🍝
	line = strings.ReplaceAll(line, "one", "o1e")
	line = strings.ReplaceAll(line, "two", "t2o")
	line = strings.ReplaceAll(line, "three", "t3e")
	line = strings.ReplaceAll(line, "four", "f4r")
	line = strings.ReplaceAll(line, "five", "f5e")
	line = strings.ReplaceAll(line, "six", "s6x")
	line = strings.ReplaceAll(line, "seven", "s7n")
	line = strings.ReplaceAll(line, "eight", "e8t")
	line = strings.ReplaceAll(line, "nine", "n9e")
	return line
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = convertLettersToNumber(line)

		numLine := buildNumLine(line)
		trunk := string(numLine[0]) + string(numLine[len(numLine)-1])
		fmt.Println(trunk)
		num, err := strconv.ParseInt(trunk, 10, 64)
		if err == nil {
			sum += int(num)
		}
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
	}
}
