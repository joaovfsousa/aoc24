package day04

import (
	"bufio"
	"fmt"
	"os"
)

type Dir int8

type Pos struct {
	x int
	y int
}

func (p Pos) Add(p2 Pos) Pos {
	return Pos{x: p.x + p2.x, y: p.y + p2.y}
}

func (p Pos) isWithinBounds(base Pos) bool {
	if p.x < 0 || p.x >= base.x {
		return false
	}

	if p.y < 0 || p.y >= base.y {
		return false
	}

	return true
}

type Node struct {
	p   Pos
	dir Pos
}

var dirsToCheck = [8]Pos{
	// top row
	{x: -1, y: -1},
	{x: 0, y: -1},
	{x: 1, y: -1},
	// middle row
	{x: -1, y: 0},
	{x: 1, y: 0},
	// bottom row
	{x: -1, y: 1},
	{x: 0, y: 1},
	{x: 1, y: 1},
}

func solve1() int {
	total := 0

	file, err := os.Open("days/day04/input01.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	m := [][]byte{}

	for scanner.Scan() {
		line := scanner.Bytes()

		m = append(m, append([]byte{}, line...))
	}

	lCount := len(m)
	cCount := len(m[0])

	base := Pos{x: cCount, y: lCount}

	for l := 0; l < lCount; l++ {
		for c := 0; c < cCount; c++ {
			currentPos := Pos{x: c, y: l}

			char := m[l][c]

			if char == 'X' {
				mNodes := []Node{}

				for _, d := range dirsToCheck {
					neighbor := currentPos.Add(d)
					if neighbor.isWithinBounds(base) && m[neighbor.y][neighbor.x] == 'M' {
						mNodes = append(mNodes, Node{
							p:   neighbor,
							dir: d,
						})
					}
				}

				aNodes := []Node{}

				for _, n := range mNodes {
					neighbor := n.p.Add(n.dir)

					if neighbor.isWithinBounds(base) && m[neighbor.y][neighbor.x] == 'A' {
						aNodes = append(aNodes, Node{
							p:   neighbor,
							dir: n.dir,
						})
					}

				}

				for _, n := range aNodes {
					neighbor := n.p.Add(n.dir)

					if neighbor.isWithinBounds(base) && m[neighbor.y][neighbor.x] == 'S' {
						total++
					}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return total
}

func solve2() int {
	total := 0

	file, err := os.Open("days/day04/input01.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		_ = scanner.Text()
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
