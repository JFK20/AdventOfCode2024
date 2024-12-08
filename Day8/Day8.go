package Day8

import (
	"AdventOfCode/helper"
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
	pos        helper.Vector2D
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
				antennas = append(antennas, antenna{identifier: matrix[i][j], pos: helper.Vector2D{X: i, Y: j}})
			}
		}
	}
	return antennas
}

func getAntiNode(pos1 helper.Vector2D, pos2 helper.Vector2D) (helper.Vector2D, bool) {
	newX := pos2.X + (pos2.X - pos1.X)
	newY := pos2.Y + (pos2.Y - pos1.Y)
	if newX >= 0 && newX < length && newY >= 0 && newY < width {
		return helper.Vector2D{X: newX, Y: newY}, true
	}
	return helper.Vector2D{}, false
}

func matchAntennas(antennas []antenna) int {
	antiNodes := make([]helper.Vector2D, 0)
	for i := 0; i < len(antennas); i++ {
		for j := 0; j < len(antennas); j++ {
			if i == j {
				continue
			}
			if antennas[i].identifier == antennas[j].identifier {
				newPos, found := getAntiNode(antennas[i].pos, antennas[j].pos)
				if found {
					flag := false
					for _, anti := range antiNodes {
						if anti == newPos {
							flag = true
						}
					}
					if !flag {
						antiNodes = append(antiNodes, newPos)
					}

				}
				newPos, found = getAntiNode(antennas[j].pos, antennas[i].pos)
				if found {
					flag := false
					for _, anti := range antiNodes {
						if anti == newPos {
							flag = true
						}
					}
					if !flag {
						antiNodes = append(antiNodes, newPos)
					}

				}
			}
		}
	}
	return len(antiNodes)
}

func SolutionDay8() {
	input := readFile("./Day8/Day8.txt")
	printRuneMatrix(input)
	antennas := getAntennas(input)
	amount := matchAntennas(antennas)
	fmt.Println(amount)

}
