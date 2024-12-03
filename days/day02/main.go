package day02

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/joaovfsousa/advent_of_code/core/nums"
)

func isSafe1(parts []string) bool {
	isIncreasing := false
	lastNum := 0

	for i, p := range parts {
		num, err := strconv.Atoi(p)
		if err != nil {
			panic("Failed to parse numbers")
		}

		if i == 0 {
			lastNum = num
			continue
		}

		diff := num - lastNum
		absDiff := nums.IntAbs(diff)

		if diff == 0 { // Should increase or decrease
			return false
		}

		if absDiff > 3 || absDiff < 1 { // diff should be at least one and at most three
			return false
		}

		if i == 1 && diff < 0 { // set tendency
			isIncreasing = true
		}

		if isIncreasing && diff > 0 { // current diff should follow the tendency
			return false
		}

		if !isIncreasing && diff < 0 { // current diff should follow the tendency
			return false
		}

		lastNum = num
	}

	return true
}

func solve1() int {
	total := 0

	file, err := os.Open("days/day02/input01.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)

		if isSafe1(parts) {
			total++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return total
}

func solve2() int {
	total := 0

	file, err := os.Open("days/day02/input01.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)

		if isSafe1(parts) {
			total++
		} else {
			for i := 0; i < len(parts); i++ {
				cpParts := make([]string, len(parts))

				copy(cpParts, parts)

				cpParts = slices.Delete(cpParts, i, i+1)

				if isSafe1(cpParts) {
					total++
					break
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return total
}

func Run() {
	res1 := solve1()
	res2 := solve2()

	fmt.Printf("Result 1: %d\n", res1)
	fmt.Printf("Result 2: %d\n", res2)
}
