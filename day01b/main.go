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
	sums := make([]int, 0)
	i := 0
	j := 1
	k := 2
	max := len(text) - 1

	for {
		a, _ := strconv.Atoi(text[i])
		b, _ := strconv.Atoi(text[j])
		c, _ := strconv.Atoi(text[k])
		sum := a + b + c
		sums = append(sums, sum)
		i++
		j++
		k++
		if i > max || j > max || k > max {
			break
		}
	}

	last := sums[0]
	for _, s := range sums {
		if s > last {
			count++
		}
		last = s
	}

	fmt.Printf("\nZähler: %d\n", count)
}
