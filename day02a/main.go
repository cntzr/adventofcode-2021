package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type (
	move struct {
		dir  string
		dist int
	}
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

	moves := make([]move, 0, len(text))
	for _, line := range text {
		splits := strings.Split(line, " ")
		distance, _ := strconv.Atoi(splits[1])
		m := move{dir: splits[0], dist: distance}
		moves = append(moves, m)
	}

	depth := 0
	distance := 0
	for _, m := range moves {
		switch m.dir {
		case "forward":
			distance += m.dist
		case "up":
			depth -= m.dist
		case "down":
			depth += m.dist
		}
	}

	fmt.Printf("Result: %d\n", depth*distance)
}
