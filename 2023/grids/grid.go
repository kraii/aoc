package grids

import (
	"aoc"
	"fmt"
)

type Point struct {
	X, Y int
}

type Direction int

const (
	Up    Direction = 0
	Right Direction = 1
	Down  Direction = 2
	Left  Direction = 3
)

var Directions = [4]Direction{Up, Right, Down, Left}

func PrintDir(dir Direction) string {
	switch dir {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	default:
		panic(badDirection(dir))
	}
}

func Move1(p Point, d Direction) Point {
	return Move(p, d, 1)
}

func Move(p Point, d Direction, distance int) Point {
	switch d {
	case Up:
		return Point{p.X, p.Y - distance}
	case Down:
		return Point{p.X, p.Y + distance}
	case Left:
		return Point{p.X - distance, p.Y}
	case Right:
		return Point{p.X + distance, p.Y}
	}
	panic(badDirection(d))
}

func badDirection(d Direction) string {
	return fmt.Sprintf("What the heck direction is %d", d)
}

func TurnLeft(current Direction) Direction {
	switch current {
	case Up:
		return Left
	case Down:
		return Right
	case Left:
		return Down
	case Right:
		return Up
	default:
		panic(badDirection(current))
	}
}

func TurnRight(current Direction) Direction {
	switch current {
	case Up:
		return Right
	case Down:
		return Left
	case Left:
		return Up
	case Right:
		return Down
	default:
		panic(badDirection(current))
	}
}

func InRange[T any](grid [][]T, p Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.Y < len(grid) && p.X < len(grid[0])
}

func ParseGrid(file string) [][]rune {
	var result [][]rune
	scanner := aoc.OpenScanner(file)
	for scanner.Scan() {
		result = append(result, []rune(scanner.Text()))
	}
	return result
}

func FindPoint[T comparable](grid [][]T, value T) Point {
	for y, row := range grid {
		for x, v := range row {
			if v == value {
				return Point{x, y}
			}
		}
	}
	return Point{-1, -1}
}

func Make2DSlice[T any](lenX, lenY int) [][]T {
	outer := make([][]T, lenY)
	contents := make([]T, lenY*lenX)
	for i := range outer {
		outer[i], contents = contents[:lenX], contents[lenX:]
	}
	return outer
}
