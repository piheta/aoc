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
	convertedLine = "one" + strings.ReplaceAll(convertedLine, "one", "1") + "one"
	convertedLine = "two" + strings.ReplaceAll(convertedLine, "two", "2") + "two"
	convertedLine = "three" + strings.ReplaceAll(convertedLine, "three", "3") + "three"
	convertedLine = "four" + strings.ReplaceAll(convertedLine, "four", "4") + "four"
	convertedLine = "five" + strings.ReplaceAll(convertedLine, "five", "5") + "five"
	convertedLine = "six" + strings.ReplaceAll(convertedLine, "six", "6") + "six"
	convertedLine = "seven" + strings.ReplaceAll(convertedLine, "seven", "7") + "seven"
	convertedLine = "eight" + strings.ReplaceAll(convertedLine, "eight", "8") + "eight"
	convertedLine = "nine" + strings.ReplaceAll(convertedLine, "nine", "9") + "nine"
	return convertedLine
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
