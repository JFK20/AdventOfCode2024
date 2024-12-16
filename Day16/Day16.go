package Day16

import (
	"AdventOfCode/mathUtil"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

var height int
var width int

func printRuneMatrix(matrix [][]rune, width int, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Print(string(matrix[y][x]))
		}
		fmt.Println()
	}
}

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
	height = len(ret)
	width = len(ret[0])
	return ret
}

func findCoordinate(grid [][]rune, target rune) mathUtil.Vector2D[int] {
	// Scan grid to find the coordinate of 'S' or 'E'
	for y := range height {
		for x := range width {
			if grid[x][y] == target {
				return mathUtil.Vector2D[int]{X: x, Y: y}
			}
		}

	}
	return mathUtil.Vector2D[int]{X: -1, Y: -1}
}

func isValidMove(grid [][]rune, position mathUtil.Vector2D[int]) bool { // Check if position is within grid bounds and not a wall
	return position.X >= 0 && position.X < width && position.Y >= 0 && position.Y < height && At(grid, position) != '#'
}

func At(grid [][]rune, position mathUtil.Vector2D[int]) rune {
	return grid[position.Y][position.X]
}

func dijkstraShortestPath(grid [][]rune) []mathUtil.Vector2D[int] {
	// Find start and end coordinates
	start := findCoordinate(grid, 'S')
	end := findCoordinate(grid, 'E')

	// Distance map to track shortest distances
	distances := make(map[mathUtil.Vector2D[int]]int)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			distances[mathUtil.Vector2D[int]{X: x, Y: y}] = math.MaxInt32
		}
	}
	distances[start] = 0

	// Priority queue to get next closest node
	pq := &mathUtil.PriorityQueue{}
	pq.Push(start, 0)

	// Track parent to reconstruct path
	parent := make(map[mathUtil.Vector2D[int]]mathUtil.Vector2D[int])

	// Track visited nodes
	visited := make(map[mathUtil.Vector2D[int]]bool)

	for pq.Len() > 0 {
		// Get current node with minimum distance
		current, _ := pq.Pop()

		// Skip if already visited
		if visited[current] {
			continue
		}
		visited[current] = true

		// Check if reached end
		if current == end {
			break
		}

		// Explore neighbors
		for _, neighbour := range current.GetAllNeighbours() {
			// Skip invalid moves
			if !isValidMove(grid, neighbour) {
				continue
			}

			// Calculate new distance (assuming each move costs 1)
			newDistance := distances[current] + 1

			// Update if new path is shorter
			if newDistance < distances[neighbour] {
				distances[neighbour] = newDistance
				parent[neighbour] = current
				pq.Push(neighbour, newDistance)
			}
		}
	}

	// Reconstruct path
	if distances[end] == math.MaxInt32 {
		return nil // No path found
	}

	path := []mathUtil.Vector2D[int]{}
	current := end
	for current != start {
		path = append([]mathUtil.Vector2D[int]{current}, path...)
		current = parent[current]
	}
	path = append([]mathUtil.Vector2D[int]{start}, path...)

	return path
}

func visPath(grid [][]rune, path []mathUtil.Vector2D[int]) [][]rune {
	newGrid := grid
	for _, pos := range path {
		grid[pos.Y][pos.X] = 'P'
	}
	return newGrid
}

func SolutionDay16() {
	input := readFile("./Day16/Day16Test.txt")
	printRuneMatrix(input, width, height)
	path := dijkstraShortestPath(input)
	newGrid := visPath(input, path)
	printRuneMatrix(newGrid, width, height)

}
