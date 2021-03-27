package main

import "fmt"

type Point struct {
	i, j int
}

func at(maze [][]int, a Point) (int, bool) {
	if a.i > len(maze)-1 || a.j >= len(maze[0]) || a.j < 0 || a.i < 0 {
		return 0, false
	}
	return maze[a.i][a.j], true
}

func (p Point) add(r Point) Point {
	return Point{p.i + r.i, p.j + r.j}
}

var dirs = [4]Point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func walk(maze [][]int, start Point, end Point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	Q := []Point{start}
	for len(Q) > 0 {
		for i := range dirs {
			cur := Q[0]
			Q = Q[1:]
			if cur == end {
				break
			}
			newPoint := cur.add(dirs[i])
			val, ok := at(maze, newPoint)
			if !ok || val == 1 {
				continue
			}
			val, ok = at(steps, newPoint)
			if val != 0 || !ok {
				continue
			}
			if newPoint == start {
				continue
			}
			curStep, _ := at(steps, cur)
			steps[newPoint.i][newPoint.j] = curStep + 1
			Q = append(Q, newPoint)
		}
	}
	return steps
}

func main() {
	fmt.Println()
}
