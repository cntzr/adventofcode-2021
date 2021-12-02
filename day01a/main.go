package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("data.txt")

	if err != nil {
		log.Fatalf("failed to open data.txt")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()

	count := 0
	last := text[0]
	for _, line := range text {
		x, _ := strconv.Atoi(line)
		y, _ := strconv.Atoi(last)
		if x > y {
			count++
		}
		last = line
	}

	fmt.Printf("\nZähler: %d\n", count)
}
