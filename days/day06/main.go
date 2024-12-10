package day06

import (
	"bufio"
	"fmt"
	"os"

	"github.com/joaovfsousa/advent_of_code/core/math"
)

var UP = math.Vector2[int]{
	X: 0,
	Y: -1,
}

var DOWN = math.Vector2[int]{
	X: 0,
	Y: 1,
}

var RIGHT = math.Vector2[int]{
	X: 1,
	Y: 0,
}

var LEFT = math.Vector2[int]{
	X: -1,
	Y: 0,
}

var dirs = []math.Vector2[int]{
	UP, RIGHT, DOWN, LEFT,
}

type Guard struct {
	loc math.Point[int]
	dir int
}

func (g Guard) IsWithinBounds(b math.Point[int]) bool {
	return g.loc.IsWithinBounds(b)
}

func (g Guard) GetDir() math.Vector2[int] {
	return dirs[g.dir%4]
}

func (g *Guard) Turn() {
	g.dir++
}

type Node struct {
	isOpen     bool
	wasVisited bool
}

func PrintMatrix(m [][]*Node, g Guard) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			n := m[i][j]

			if g.loc.X == j && g.loc.Y == i {
				fmt.Print("^")
				continue
			}

			if !n.isOpen {
				fmt.Print("#")
				continue
			}

			if !n.wasVisited {
				fmt.Print(".")
				continue
			}

			if n.wasVisited {
				fmt.Print("X")
			}
		}

		fmt.Println()
	}
}

func solve1() int {
	total := 0

	file, err := os.Open("days/day06/input")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var guard Guard

	m := [][]*Node{}

	y := 0

	for scanner.Scan() {
		line := scanner.Text()

		parsedLine := []*Node{}

		x := 0

		for _, b := range line {
			switch b {
			case '.':
				parsedLine = append(parsedLine, &Node{
					isOpen:     true,
					wasVisited: false,
				})
			case '#':
				parsedLine = append(parsedLine, &Node{
					isOpen:     false,
					wasVisited: false,
				})
			case '^':
				parsedLine = append(parsedLine, &Node{
					isOpen:     true,
					wasVisited: true,
				})
				guard = Guard{
					loc: math.Point[int]{X: x, Y: y},
					dir: 0,
				}
			}

			x++
		}

		m = append(m, parsedLine)
		y++
	}

	b := math.Point[int]{
		X: len(m[0]),
		Y: len(m),
	}

	for guard.IsWithinBounds(b) {
		n := m[guard.loc.Y][guard.loc.X]

		if !n.isOpen {
			guard.loc.Move(guard.GetDir(), -1)
			guard.Turn()
			// PrintMatrix(m, guard)
			// fmt.Println("\n---")
		}

		n.wasVisited = true

		guard.loc.Move(guard.GetDir(), 1)
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			n := m[i][j]

			if n.wasVisited && n.isOpen {
				total++
			}
		}
	}

	// PrintMatrix(m, guard)
	return total
}

func solve2() int {
	total := 0

	file, err := os.Open("days/day06/input")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		_ = line
	}

	return total
}

func Run() {
	res1 := solve1()
	res2 := solve2()

	fmt.Printf("Result 1: %d\n", res1)
	fmt.Printf("Result 2: %d\n", res2)
}
