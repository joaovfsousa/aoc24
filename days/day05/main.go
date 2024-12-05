package day05

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/joaovfsousa/advent_of_code/core/parsing"
)

func isOrdered(nums []int, rules map[int][]int) bool {
	for i, n := range nums {
		after := rules[n]
		for _, p := range nums[:i] {
			if slices.Index(after, p) != -1 {
				return false
			}
		}
	}

	return true
}

func orderNums(nums []int, rules map[int][]int) {
	n := len(nums)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			before := nums[j]
			after := nums[j+1]
			r := rules[after]
			if slices.Index(r, before) != -1 {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func solve1() int {
	total := 0

	file, err := os.Open("days/day05/input01.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	isScanningRules := true

	rules := make(map[int][]int)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" { // rules/updates separator
			isScanningRules = false
			continue
		}

		if isScanningRules {
			nums := strings.Split(line, "|")

			before, errB := strconv.Atoi(nums[0])
			after, errA := strconv.Atoi(nums[1])

			if errB != nil || errA != nil {
				panic(fmt.Sprintf("Failed to parse => %s", line))
			}

			rules[before] = append(rules[before], after)
		} else {
			parts := strings.Split(line, ",")

			nums := parsing.StrSliceToIntSlice(parts)

			if isOrdered(nums, rules) {
				total += nums[(len(nums)-1)/2]
			}
		}
	}

	return total
}

func solve2() int {
	total := 0

	file, err := os.Open("days/day05/input01.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	rules := make(map[int][]int)
	isScanningRules := true

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" { // rules/updates separator
			isScanningRules = false
			continue
		}

		if isScanningRules {
			nums := strings.Split(line, "|")

			before, errB := strconv.Atoi(nums[0])
			after, errA := strconv.Atoi(nums[1])

			if errB != nil || errA != nil {
				panic(fmt.Sprintf("Failed to parse => %s", line))
			}

			rules[before] = append(rules[before], after)
		} else {
			parts := strings.Split(line, ",")

			nums := parsing.StrSliceToIntSlice(parts)

			if !isOrdered(nums, rules) {
				orderNums(nums, rules)
				total += nums[(len(nums)-1)/2]
			}
		}
	}

	return total
}

func Run() {
	res1 := solve1()
	res2 := solve2()

	fmt.Printf("Result 1: %d\n", res1)
	fmt.Printf("Result 2: %d\n", res2)
}
