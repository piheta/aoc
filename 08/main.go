package main

import (
	"bufio"
	"fmt"
	"os"
)

var nav = ""
var directions = map[string][]string{}
var a_directions = map[string][]string{}

func fillMap(line string) {
	a := line[0]
	b := line[1]
	c := line[2]

	d := line[7]
	e := line[8]
	f := line[9]

	g := line[12]
	h := line[13]
	i := line[14]

	directions[string(a)+string(b)+string(c)] = []string{string(d) + string(e) + string(f), string(g) + string(h) + string(i)}
}

var ans = 0

func countSteps(start string, stop string, step int) {
	if start == stop {
		return
	}

	ans++

	if step == len(nav) {
		step = 0
	}

	switch nav[step] {
	case 'R':
		countSteps(directions[start][1], stop, step+1)
	case 'L':
		countSteps(directions[start][0], stop, step+1)
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 20 {
			nav = line
		} else if len(line) > 5 {
			fillMap(line)
		}

	}
	countSteps("AAA", "ZZZ", 0)
	fmt.Println(ans)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
	}
}
