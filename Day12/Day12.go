package Day12

import (
	"AdventOfCode/vec"
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

var length int
var width int
var visited []vec.Vector2D

func readFile(filename string) [][]rune {
	ret := make([][]rune, 0)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Problem opening File", err)
		return nil
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		runeLine := make([]rune, 0)
		for _, char := range scanner.Text() {
			runeLine = append(runeLine, char)
		}
		ret = append(ret, runeLine)
	}
	length = len(ret)
	width = len(ret[0])
	return ret
}

func findCosts(garden [][]rune) int {
	allVisited := make([]vec.Vector2D, 0)
	sum := 0

	for y := range garden {
		for x := range garden {
			pos := vec.Vector2D{X: x, Y: y}
			if !slices.Contains(allVisited, pos) {
				visited = make([]vec.Vector2D, 0)
				a, b := regionCircumference(garden, pos)
				sum += a * b
				allVisited = append(allVisited, visited...)
			}
		}
	}
	return sum
}

func regionCircumference(garden [][]rune, pos vec.Vector2D) (int, int) {
	if slices.Contains(visited, pos) {
		return 0, len(visited)
	}
	visited = append(visited, pos)

	bounds := vec.Vector2D{X: length, Y: width}
	regionSymbol := garden[pos.Y][pos.X]
	neighbours := make([]vec.Vector2D, 0)
	circumference := 4
	for _, v := range pos.GetAllNeighbours() {
		if v.IsInBounds(bounds) && garden[v.Y][v.X] == regionSymbol {
			if !slices.Contains(visited, v) {
				neighbours = append(neighbours, v)
			}
			circumference--
		}
	}

	for _, v := range neighbours {
		c, _ := regionCircumference(garden, v)
		circumference += c
	}
	return circumference, len(visited)
}

func checkAll4(input [][]rune, current vec.Vector2D) []vec.Vector2D {
	sameAround := []vec.Vector2D{}
	bounds := vec.Vector2D{X: length, Y: width}
	if !current.IsInBounds(bounds) {
		return sameAround
	}
	regionSymbol := input[current.Y][current.X]
	neighbours := current.GetAllNeighbours()
	for _, neighbour := range neighbours {
		if neighbour.IsInBounds(bounds) && input[neighbour.Y][neighbour.X] == regionSymbol {
			sameAround = append(sameAround, neighbour)
		}
	}
	return sameAround
}

type polynomial struct {
	area  int
	sides int
}

func alternativeSolution(input [][]rune) int {
	cost2 := 0
	visitedCoordinates := make(map[vec.Vector2D]struct{})

	for j, _ := range input {
		for i, _ := range input[j] {
			if _, ok := visitedCoordinates[vec.Vector2D{X: i, Y: j}]; !ok {
				next := []vec.Vector2D{{i, j}}
				shape := polynomial{}
				for len(next) != 0 {
					newShape, traverseNext := findAllGardensNonRecursively(input, next[0], shape, visitedCoordinates)
					shape = newShape
					next = append(next, traverseNext...)
					next = next[1:]
				}
				cost2 += shape.area * shape.sides
			}
		}
	}
	return cost2
}

func checkCorners(input [][]rune, current vec.Vector2D) int {
	count := 0
	gardenType := input[current.Y][current.X]
	x, y := current.X, current.Y

	if x == 0 && y == 0 {
		count += 1
	}

	if x == 0 && y == len(input)-1 {
		count += 1
	}

	if x == len(input[0])-1 && y == len(input)-1 {
		count += 1
	}

	if x == len(input[0])-1 && y == 0 {
		count += 1
	}

	// top left outside corner
	// ##   __   |#
	// #O   #O   |O
	if (x > 0 && y > 0 && input[y][x-1] != gardenType && input[y-1][x] != gardenType) ||
		(x > 0 && y == 0 && input[y][x-1] != gardenType) || (x == 0 && y > 0 && input[y-1][x] != gardenType) {
		count += 1
	}

	// top left inside corner
	// OO
	// O#
	if x < len(input[0])-1 && y < len(input)-1 && input[y][x+1] == gardenType && input[y+1][x] == gardenType && input[y+1][x+1] != gardenType {
		count += 1
	}

	// top right outside corner
	// ##   __    #|
	// O#   O#    O|
	if (x < len(input[0])-1 && y > 0 && input[y][x+1] != gardenType && input[y-1][x] != gardenType) ||
		(x < len(input[0])-1 && y == 0 && input[y][x+1] != gardenType) || (x == len(input[0])-1 && y > 0 && input[y-1][x] != gardenType) {
		count += 1
	}

	// top right inside corner
	// OO
	// #O
	if x > 0 && y < len(input)-1 && input[y][x-1] == gardenType && input[y+1][x] == gardenType && input[y+1][x-1] != gardenType {
		count += 1
	}

	// bottom left outside corner
	// #O   #O    |O
	// ##   --    |#
	if (x > 0 && y < len(input)-1 && input[y][x-1] != gardenType && input[y+1][x] != gardenType) ||
		(x > 0 && y == len(input)-1 && input[y][x-1] != gardenType) || (x == 0 && y < len(input)-1 && input[y+1][x] != gardenType) {
		count += 1
	}

	// bottom left inside corner
	// O#
	// OO
	if x < len(input[0])-1 && y > 0 && input[y][x+1] == gardenType && input[y-1][x] == gardenType && input[y-1][x+1] != gardenType {
		count += 1
	}

	// bottom right outside corner
	// O#   O#    O|
	// ##   --    #|
	if (x < len(input[0])-1 && y < len(input)-1 && input[y][x+1] != gardenType && input[y+1][x] != gardenType) ||
		(x < len(input[0])-1 && y == len(input)-1 && input[y][x+1] != gardenType) || (x == len(input[0])-1 && y < len(input)-1 && input[y+1][x] != gardenType) {
		count += 1
	}

	// bottom right inside corner
	// #O
	// OO
	if x > 0 && y > 0 && input[y][x-1] == gardenType && input[y-1][x] == gardenType && input[y-1][x-1] != gardenType {
		count += 1
	}
	return count
}

func findAllGardensNonRecursively(input [][]rune, current vec.Vector2D, shape polynomial, visited map[vec.Vector2D]struct{}) (polynomial, []vec.Vector2D) {
	if _, ok := visited[current]; ok {
		return shape, []vec.Vector2D{}
	}

	checkNext := checkAll4(input, current)

	// none surrounding are same garden
	if len(checkNext) == 0 {
		if shape.area == 0 {
			visited[current] = struct{}{}
			shape = polynomial{
				area: 1, sides: 4,
			}
			return shape, []vec.Vector2D{}
		}
		return shape, []vec.Vector2D{}
	}

	shape.area += 1
	visited[current] = struct{}{}
	shape.sides += checkCorners(input, current)

	return shape, checkNext
}

func SolutionDay12() {
	input := readFile("./Day12/Day12.txt")
	cost := findCosts(input)
	fmt.Printf("Solution Day11 Part 1: %d\n", cost)
	cost2 := alternativeSolution(input)
	fmt.Printf("Solution Day11 Part 2: %d\n", cost2)
}
