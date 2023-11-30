#!/bin/bash

base_dir="aoc2023"

# Create the base directory
mkdir -p "$base_dir"

# Create directories for each day
for ((day=1; day<=23; day++))
do
    day_dir="$base_dir/day$day"
    mkdir -p "$day_dir"
    touch "$day_dir/input.txt"
    
    # Create main.go and add code to read input.txt
    main_file="$day_dir/main.go"
    touch "$main_file"
    echo "package main" >> "$main_file"
    echo >> "$main_file"
    echo "import (" >> "$main_file"
    echo "    \"bufio\"" >> "$main_file"
    echo "    \"fmt\"" >> "$main_file"
    echo "    \"os\"" >> "$main_file"
    echo ")" >> "$main_file"
    echo >> "$main_file"
    echo "func main() {" >> "$main_file"
    echo "    file, err := os.Open(\"input.txt\")" >> "$main_file"
    echo "    if err != nil {" >> "$main_file"
    echo "        fmt.Println(\"Error opening file\", err)" >> "$main_file"
    echo "        return" >> "$main_file"
    echo "    }" >> "$main_file"
    echo "    defer file.Close()" >> "$main_file"
    echo >> "$main_file"
    echo "    scanner := bufio.NewScanner(file)" >> "$main_file"
    echo "    for scanner.Scan() {" >> "$main_file"
    echo "        line := scanner.Text()" >> "$main_file"
    echo "        fmt.Println(line)" >> "$main_file"
    echo "    }" >> "$main_file"
    echo >> "$main_file"
    echo "    if err := scanner.Err(); err != nil {" >> "$main_file"
    echo "        fmt.Println(\"Error reading file\", err)" >> "$main_file"
    echo "    }" >> "$main_file"
    echo "}" >> "$main_file"
done
