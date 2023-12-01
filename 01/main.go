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

func convertLettersToNumber(line string) string {
	convertedLine := line
	convertedLine = strings.ReplaceAll(convertedLine, "one", "one1one")
	convertedLine = strings.ReplaceAll(convertedLine, "two", "two2two")
	convertedLine = strings.ReplaceAll(convertedLine, "three", "three3three")
	convertedLine = strings.ReplaceAll(convertedLine, "four", "four4four")
	convertedLine = strings.ReplaceAll(convertedLine, "five", "five5five")
	convertedLine = strings.ReplaceAll(convertedLine, "six", "six6six")
	convertedLine = strings.ReplaceAll(convertedLine, "seven", "seven7seven")
	convertedLine = strings.ReplaceAll(convertedLine, "eight", "eight8eight")
	convertedLine = strings.ReplaceAll(convertedLine, "nine", "nine9nine")
	fmt.Println(convertedLine)
	return convertedLine //eightwothree eighttwo2twothree eighttwo2twothree3three eight8eighttwo2twothree3three
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
