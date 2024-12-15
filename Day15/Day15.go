package Day15

import (
	"AdventOfCode/mathUtil"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func printGrid(gridMap map[mathUtil.Vector2D[int]]rune) {
	// Determine grid bounds
	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt

	for vec := range gridMap {
		if vec.X < minX {
			minX = vec.X
		}
		if vec.Y < minY {
			minY = vec.Y
		}
		if vec.X > maxX {
			maxX = vec.X
		}
		if vec.Y > maxY {
			maxY = vec.Y
		}
	}

	// Create a 2D slice to represent the grid
	width := maxX - minX + 1
	height := maxY - minY + 1
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = '.' // Default character
		}
	}

	// Populate the grid with runes from the map
	for vec, r := range gridMap {
		grid[vec.Y-minY][vec.X-minX] = r
	}

	// Print the grid row by row
	for _, row := range grid {
		rowStr := string(row) // Convert rune slice to string
		fmt.Println(rowStr)   // Print the entire row
	}
}

func gridToSlice(gridMap map[mathUtil.Vector2D[int]]rune) [][]rune {
	// Determine grid bounds
	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt

	for vec := range gridMap {
		if vec.X < minX {
			minX = vec.X
		}
		if vec.Y < minY {
			minY = vec.Y
		}
		if vec.X > maxX {
			maxX = vec.X
		}
		if vec.Y > maxY {
			maxY = vec.Y
		}
	}

	// Create a 2D slice to represent the grid
	width := maxX - minX + 1
	height := maxY - minY + 1
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = '.' // Default character
		}
	}

	// Populate the grid with runes from the map
	for vec, r := range gridMap {
		grid[vec.Y-minY][vec.X-minX] = r
	}

	return grid
}

func readFile(filename string) (map[mathUtil.Vector2D[int]]rune, []rune) {
	grid := make(map[mathUtil.Vector2D[int]]rune)
	moves := make([]rune, 0)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Problem opening File", err)
		return nil, nil
	}
	scanner := bufio.NewScanner(file)
	ver := 0
	i := 0
	for scanner.Scan() {
		if len(scanner.Text()) < 1 {
			ver++
		}
		if ver == 0 {
			splited := []rune(scanner.Text())
			for j, v := range splited {
				grid[mathUtil.Vector2D[int]{X: j, Y: i}] = v
			}
			i++
		}
		if ver == 1 {
			splited := []rune(scanner.Text())
			moves = append(moves, splited...)
		}

	}
	return grid, moves
}

func readFile2(filename string) (map[mathUtil.Vector2D[int]]rune, []rune) {
	grid := make(map[mathUtil.Vector2D[int]]rune)
	moves := make([]rune, 0)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Problem opening File", err)
		return nil, nil
	}
	scanner := bufio.NewScanner(file)
	ver := 0
	i := 0
	for scanner.Scan() {
		if len(scanner.Text()) < 1 {
			ver++
		}
		if ver == 0 {
			splited := []rune(scanner.Text())
			offsett := 0
			for j, v := range splited {
				if v == '.' {
					grid[mathUtil.Vector2D[int]{X: j + offsett, Y: i}] = v
					offsett++
					grid[mathUtil.Vector2D[int]{X: j + offsett, Y: i}] = v
				}
				if v == 'O' {
					grid[mathUtil.Vector2D[int]{X: j + offsett, Y: i}] = '['
					offsett++
					grid[mathUtil.Vector2D[int]{X: j + offsett, Y: i}] = ']'
				}
				if v == '#' {
					grid[mathUtil.Vector2D[int]{X: j + offsett, Y: i}] = v
					offsett++
					grid[mathUtil.Vector2D[int]{X: j + offsett, Y: i}] = v
				}
				if v == '@' {
					grid[mathUtil.Vector2D[int]{X: j + offsett, Y: i}] = v
					offsett++
					grid[mathUtil.Vector2D[int]{X: j + offsett, Y: i}] = '.'
				}

			}
			i++
		}
		if ver == 1 {
			splited := []rune(scanner.Text())
			moves = append(moves, splited...)
		}

	}
	return grid, moves
}

func move(grid map[mathUtil.Vector2D[int]]rune, symbol rune, pos *mathUtil.Vector2D[int], m rune) (map[mathUtil.Vector2D[int]]rune, bool) {
	var direction = mathUtil.Vector2D[int]{X: 0, Y: 0}
	if m == '^' {
		direction.Y -= 1
	} else if m == 'v' {
		direction.Y += 1
	} else if m == '<' {
		direction.X -= 1
	} else if m == '>' {
		direction.X += 1
	}
	newPos := mathUtil.AddVector2D(*pos, direction)
	if grid[newPos] == '#' {
		return grid, false
	}
	if grid[newPos] == '.' {
		grid[*pos] = '.'
		*pos = mathUtil.AddVector2D(*pos, direction)
		grid[*pos] = symbol
		return grid, true
	}
	if grid[newPos] == 'O' {
		newGrid, suc := move(grid, grid[newPos], &newPos, m)
		if !suc {
			return grid, false
		} else {
			grid = newGrid
			grid[*pos] = '.'
			*pos = mathUtil.AddVector2D(*pos, direction)
			grid[*pos] = symbol
			return grid, true
		}
	}

	return grid, true
}

func findStartPos2(warehouse [][]rune) mathUtil.Vector2D[int] {
	for y, l := range warehouse {
		for x, r := range l {
			if r == '@' {
				return mathUtil.Vector2D[int]{x, y}
			}
		}
	}
	return mathUtil.Vector2D[int]{}
}

func runeToDirVec(r rune) mathUtil.Vector2D[int] {
	if r == '<' {
		return mathUtil.Vector2D[int]{-1, 0}
	}
	if r == '>' {
		return mathUtil.Vector2D[int]{1, 0}
	}
	if r == '^' {
		return mathUtil.Vector2D[int]{0, -1}
	}
	if r == 'v' {
		return mathUtil.Vector2D[int]{0, 1}
	}

	return mathUtil.Vector2D[int]{}
}

func moveRobot2(warehouse [][]rune, movement []rune) {
	pos := findStartPos2(warehouse)

	for _, m := range movement {
		dir := runeToDirVec(m)
		lookAhead := mathUtil.AddVector2D(pos, dir)

		if dir.Y != 0 {
			toMove := make([]mathUtil.Vector2D[int], 1)
			toMove[0] = mathUtil.AddVector2D(pos, dir)
			lookAheads := make([]mathUtil.Vector2D[int], 1)
			nextLookAheads := make([]mathUtil.Vector2D[int], 0)
			lookAheads[0] = lookAhead
			moved := true
			isAnyBox := true

			for isAnyBox && moved {
				isAnyBox = false
				for _, lA := range lookAheads {
					if warehouse[lookAhead.Y][lookAhead.X] == '[' {
						nextLookAheads = append(nextLookAheads, mathUtil.AddVector2D(lA, dir))
						nextLookAheads = append(nextLookAheads, mathUtil.AddVector2D(mathUtil.AddVector2D(lA, mathUtil.Vector2D[int]{1, 0}), dir))
						isAnyBox = true
					} else if warehouse[lookAhead.Y][lookAhead.X] == ']' {
						nextLookAheads = append(nextLookAheads, mathUtil.AddVector2D(lA, dir))
						nextLookAheads = append(nextLookAheads, mathUtil.AddVector2D(mathUtil.AddVector2D(lA, mathUtil.Vector2D[int]{-1, 0}), dir))
						isAnyBox = true
					} else if warehouse[lookAhead.Y][lookAhead.X] == '#' {
						moved = false
					}
				}
				toMove = append(toMove, nextLookAheads...)
				lookAheads = nextLookAheads
				nextLookAheads = make([]mathUtil.Vector2D[int], 0)
			}
			if !moved {
				continue
			}
			toMove = mathUtil.Distinct(toMove)
			fmt.Println(len(toMove))
			for i := len(toMove) - 1; i >= 0; i-- {
				v := toMove[i]
				moveFrom := mathUtil.SubVector2D(v, dir)
				toPlace := warehouse[moveFrom.Y][moveFrom.X]
				warehouse[v.Y][v.X] = toPlace
				warehouse[moveFrom.Y][moveFrom.X] = '.'
			}

			pos = mathUtil.AddVector2D(pos, dir)
			continue
		}

		boxesToMove := 0

		for warehouse[lookAhead.Y][lookAhead.X] == ']' || warehouse[lookAhead.Y][lookAhead.X] == '[' {
			lookAhead = mathUtil.AddVector2D(lookAhead, dir)
			boxesToMove++
		}

		if warehouse[lookAhead.Y][lookAhead.X] == '#' {
			continue
		}
		behind := mathUtil.SubVector2D(lookAhead, dir)
		for range boxesToMove {
			warehouse[lookAhead.Y][lookAhead.X] = warehouse[behind.Y][behind.X]
			warehouse[behind.Y][behind.X] = '.'

			lookAhead = mathUtil.SubVector2D(lookAhead, dir)
			behind = mathUtil.SubVector2D(lookAhead, dir)
		}
		warehouse[lookAhead.Y][lookAhead.X] = '@'
		warehouse[behind.Y][behind.X] = '.'
		pos = lookAhead
	}
}

func findStartPos(grid map[mathUtil.Vector2D[int]]rune) mathUtil.Vector2D[int] {
	for pos, val := range grid {
		if val == '@' {
			return pos
		}
	}
	return mathUtil.Vector2D[int]{}
}

func sumBoxes(grid map[mathUtil.Vector2D[int]]rune) int {
	sum := 0
	for pos, val := range grid {
		if val == 'O' {
			sum += pos.Y*100 + pos.X
		}
	}
	return sum
}

func moveAll(grid map[mathUtil.Vector2D[int]]rune, moves []rune) map[mathUtil.Vector2D[int]]rune {
	pos := findStartPos(grid)
	printGrid(grid)
	for i := 0; i < len(moves); i++ {
		move(grid, grid[pos], &pos, moves[i])
	}
	printGrid(grid)
	return grid
}

func getBoxDistance2(warehouse [][]rune) int {
	sum := 0
	for y, l := range warehouse {
		for x, r := range l {
			if r == '[' {
				sum += 100*y + x
			}
		}
	}
	return sum
}

func SolutionDay15() {
	grid, moves := readFile("./Day15/Day15Test.txt")
	fmt.Println(moves)
	fmt.Println(grid)
	grid = moveAll(grid, moves)
	fmt.Printf("Solution Day15 Part 1: %d\n", sumBoxes(grid))
}
