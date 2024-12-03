package day03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func solve1() int {
	total := 0

	file, err := os.Open("days/day03/input01.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	regex, err := regexp.Compile(`mul\((?P<arg1>\d{1,3}),(?P<arg2>\d{1,3})\)`)
	if err != nil {
		panic(err)
	}

	for scanner.Scan() {
		line := scanner.Text()

		for _, match := range regex.FindAllStringSubmatch(line, -1) {
			arg1, err1 := strconv.Atoi(match[1])
			arg2, err2 := strconv.Atoi(match[2])

			if err1 != nil || err2 != nil {
				panic("Failed to parse ints")
			}

			total += arg1 * arg2
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return total
}

func solve2() int {
	total := 0
	do := true

	file, err := os.Open("days/day03/input01.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	regex, err := regexp.Compile(`(?P<fn>do|don't|mul)\((?P<args>[\d,]{0,7})\)`)
	if err != nil {
		panic(err)
	}

	for scanner.Scan() {
		line := scanner.Text()

		for _, match := range regex.FindAllStringSubmatch(line, -1) {
			matchType := match[1]

			switch matchType {
			case "mul":
				if do {
					args := match[2]

					separateArgs := strings.Split(args, ",")

					if len(separateArgs) != 2 {
						fmt.Println("match")
						fmt.Println(match)
						fmt.Println(args)
						fmt.Println(separateArgs)
						panic("Oh no")
					}

					arg1, err1 := strconv.Atoi(separateArgs[0])
					arg2, err2 := strconv.Atoi(separateArgs[1])

					if err1 != nil || err2 != nil {
						panic("Failed to parse ints")
					}

					total += arg1 * arg2
				}
			case "do":
				do = true
			case "don't":
				do = false
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
