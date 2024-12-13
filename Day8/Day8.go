package Day8

import (
	"AdventOfCode/mathUtil"
	"bufio"
	"fmt"
	"log"
	"os"
)

var length int
var width int

func printRuneMatrix(matrix [][]rune) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(string(matrix[i][j]))
		}
		fmt.Println()
	}
}

type antenna struct {
	identifier rune
	pos        mathUtil.Vector2D[int]
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
	length = len(ret)
	width = len(ret[0])
	return ret
}

func getAntennas(matrix [][]rune) []antenna {
	antennas := make([]antenna, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != '.' {
				antennas = append(antennas, antenna{identifier: matrix[i][j], pos: mathUtil.Vector2D[int]{X: i, Y: j}})
			}
		}
	}
	return antennas
}

func getAntiNode(pos1 mathUtil.Vector2D[int], pos2 mathUtil.Vector2D[int], solution bool) []mathUtil.Vector2D[int] {
	found := make([]mathUtil.Vector2D[int], 0)
	newX := pos2.X + (pos2.X - pos1.X)
	newY := pos2.Y + (pos2.Y - pos1.Y)
	if solution {
		found = append(found, mathUtil.Vector2D[int]{X: pos2.X, Y: pos2.Y})
		for newX >= 0 && newX < length && newY >= 0 && newY < width {
			found = append(found, mathUtil.Vector2D[int]{X: newX, Y: newY})
			newX += pos2.X - pos1.X
			newY += pos2.Y - pos1.Y
		}
	} else {
		if newX >= 0 && newX < length && newY >= 0 && newY < width {
			found = append(found, mathUtil.Vector2D[int]{X: newX, Y: newY})
		}
	}

	return found
}

func addUnique(s *[]mathUtil.Vector2D[int], vector mathUtil.Vector2D[int]) {
	flag := false
	for _, current := range *s {
		if current == vector {
			flag = true
		}
	}
	if !flag {
		*s = append(*s, vector)
	}
}

func matchAntennas(antennas []antenna, solution bool) int {
	antiNodes := make([]mathUtil.Vector2D[int], 0)
	for i := 0; i < len(antennas); i++ {
		for j := 0; j < len(antennas); j++ {
			if i == j {
				continue
			}
			if antennas[i].identifier == antennas[j].identifier {
				newPoses := getAntiNode(antennas[i].pos, antennas[j].pos, solution)
				for _, newPos := range newPoses {
					addUnique(&antiNodes, newPos)
				}
				newPoses = getAntiNode(antennas[j].pos, antennas[i].pos, solution)
				for _, newPos := range newPoses {
					addUnique(&antiNodes, newPos)
				}
			}
		}
	}
	return len(antiNodes)
}

func SolutionDay8() {
	input := readFile("./Day8/Day8.txt")
	//printRuneMatrix(input)
	antennas := getAntennas(input)
	amount := matchAntennas(antennas, false)
	fmt.Printf("Solution Day8 Part 1: %d\n", amount)
	amount2 := matchAntennas(antennas, true)
	fmt.Printf("Solution Day8 Part 2: %d\n", amount2)

}
