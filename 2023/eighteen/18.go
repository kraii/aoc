package eighteen

import (
	. "aoc"
	. "aoc/grids"
	"fmt"
	"strconv"
	"strings"
)

type dig struct {
	dir      Direction
	distance int
}

func solve(file string, useHex bool) int {
	digs := parse(file, useHex)
	area := 0
	current := Point{X: 0, Y: 0}
	for _, d := range digs {
		println(PrintDir(d.dir), d.distance)
		next := Move(current, d.dir, d.distance)
		// shoelace formula
		area += (current.X*next.Y - current.Y*next.X) + d.distance
		current = next
	}
	return area/2 + 1
}

func parse(file string, useHex bool) []dig {
	var digs []dig

	scanner := OpenScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if useHex {
			hexString := strings.Trim(fields[2], "(#)")

			distance := parseHex(hexString[:len(hexString)-1])
			direction := toDirection(string(hexString[len(hexString)-1]))
			digs = append(digs, dig{direction, distance})
		} else {
			digs = append(digs, dig{parseDir(fields[0]), ToInt(fields[1])})
		}
	}
	return digs
}

func toDirection(hexDigit string) Direction {
	hex := parseHex(hexDigit)
	switch hex {
	case 0:
		return Right
	case 1:
		return Down
	case 2:
		return Left
	case 3:
		return Up
	default:
		panic(fmt.Sprintf("Dunno %s %d", hexDigit, hex))

	}
}

func parseHex(hexString string) int {
	i, err := strconv.ParseInt(hexString, 16, strconv.IntSize)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func parseDir(dir string) Direction {
	switch dir {
	case "U":
		return Up
	case "D":
		return Down
	case "L":
		return Left
	case "R":
		return Right
	default:
		panic(fmt.Sprintf("Unknown dir %s", dir))
	}
}
