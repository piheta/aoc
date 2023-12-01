package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
		line := string(scanner.Text())

		numLine := ""

		for i := 0; i < len(line); i++ {
			if _, err := strconv.Atoi(string(line[i])); err == nil {
				numLine += string(line[i])
			}
		}

		lineSum := string(numLine[0]) + string(numLine[len(numLine)-1])
		num, err := strconv.ParseInt(lineSum, 10, 64)
		if err == nil {
			sum += int(num)
		}

	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
	}
}
