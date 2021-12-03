package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	oxyBin := filter(text, 0, true)
	oxyNum, err := strconv.ParseInt(oxyBin, 2, 64)
	if err != nil {
		fmt.Printf("oxygen failed: %s\n", err.Error())
	}
	coBin := filter(text, 0, false)
	coNum, err := strconv.ParseInt(coBin, 2, 64)
	if err != nil {
		fmt.Printf("co2 failed: %s\n", err.Error())
	}
	fmt.Printf("Oxygen: bin %s ... int: %d\n", oxyBin, oxyNum)
	fmt.Printf("CO2:    bin %s ... int: %d\n", coBin, coNum)
	fmt.Printf("Result: %d\n", oxyNum*coNum)
}

// maxOneZero ... delivers the string "0" or "1" for the largest number
func maxOneZero(one, zero int) string {
	if one >= zero {
		return "1"
	}
	return "0"
}

// minOneZero ... delivers the string "0" or "1" for the lowest number
func minOneZero(one, zero int) string {
	if one < zero {
		return "1"
	}
	return "0"
}

// countOneZero ... counts 0 and 1 at a specific string position
func countOneZero(list []string, pos int) (int, int) {
	countOne := 0
	countZero := 0
	for i := range list {
		splits := strings.Split(list[i], "")
		if splits[pos] == "0" {
			countZero++
		} else {
			countOne++
		}
	}
	return countOne, countZero
}

// filter ... lazy version of recursive filter
func filter(list []string, pos int, oxygen bool) string {
	one, zero := countOneZero(list, pos)
	pattern := ""
	if oxygen {
		pattern = maxOneZero(one, zero)
	} else {
		pattern = minOneZero(one, zero)
	}
	sublist := make([]string, 0)
	for i := range list {
		if list[i][pos:pos+1] == pattern {
			sublist = append(sublist, list[i])
		}
	}

	if len(sublist) == 1 { // we got it!
		return sublist[0]
	}
	// dive further into recursion
	return filter(sublist, pos+1, oxygen)
}
