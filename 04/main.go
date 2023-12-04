package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func findCardTotal(line string) float64 {
	cardTotal := 0

	line = strings.ReplaceAll(line, " |", "")
	line = strings.ReplaceAll(line, "  ", " ")

	_, nums, _ := strings.Cut(line, ": ")
	numArr := strings.Split(nums, " ")
	winners := []string{}

	for i := 0; i < len(numArr); i++ {
		if i < 10 {
			winners = append(winners, numArr[i])
		} else if slices.Contains(winners, numArr[i]) {
			cardTotal++
		}
	}

	if cardTotal == 0 {
		return 0
	}

	return math.Pow(2.0, (float64(cardTotal) - 1))
}

func main() {
	sum := 0.0
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sum += findCardTotal(line)
	}

	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
	}
}
