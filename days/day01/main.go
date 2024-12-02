package day01

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parse() ([]int, []int) {
	list01, list02 := []int{}, []int{}

	file, err := os.Open("days/day01/input01.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)

		num01, err1 := strconv.Atoi(parts[0])
		num02, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			panic("Failed to parse numbers")
		}

		list01 = append(list01, num01)
		list02 = append(list02, num02)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return list01, list02
}

func solve1() int {
	dist := 0
	list1, list2 := parse()

	slices.Sort(list1)
	slices.Sort(list2)

	for i := 0; i < len(list1); i++ {
		diff := list1[i] - list2[i]

		if diff < 0 {
			dist += -diff
		} else {
			dist += diff
		}
	}

	return dist
}

func solve2() int {
	total := 0

	list1, list2 := parse()

	countsList2 := make(map[int]int)
	visited := make(map[int]bool)

	for _, v := range list2 {
		countsList2[v]++
	}

	fmt.Printf("%v\n", countsList2)

	for _, k := range list1 {
		if !visited[k] {
			total += k * countsList2[k]
			visited[k] = true
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
