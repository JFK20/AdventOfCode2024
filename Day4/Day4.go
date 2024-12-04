package Day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(filename string) [][]string {
	var ret [][]string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Problem opening File", err)
		return nil
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret = append(ret, strings.Split(scanner.Text(), ""))
	}
	return ret
}

func searchForXMAS(charMatrix [][]string) int {
	rows, cols := len(charMatrix), len(charMatrix[0])
	count := 0

	// Directions: right, down, diagonal down-right, diagonal down-left
	directions := [][]int{
		{0, 1},
		{1, 0},
		{1, 1},
		{1, -1},
	}

	for startRow := 0; startRow < rows; startRow++ {
		for startCol := 0; startCol < cols; startCol++ {
			for _, dir := range directions {
				//forward
				if checkXMASSequence(charMatrix, startRow, startCol, dir[0], dir[1]) {
					count++
				}

				//reverse
				if checkXMASSequence(charMatrix, startRow, startCol, -dir[0], -dir[1]) {
					count++
				}
			}
		}
	}

	return count
}

func checkXMASSequence(charMatrix [][]string, row, col, rowStep, colStep int) bool {
	rows, cols := len(charMatrix), len(charMatrix[0])

	target := []string{"X", "M", "A", "S"}

	if row+rowStep*3 < 0 || row+rowStep*3 >= rows ||
		col+colStep*3 < 0 || col+colStep*3 >= cols {
		return false
	}

	for i, char := range target {
		checkRow := row + i*rowStep
		checkCol := col + i*colStep

		if checkRow < 0 || checkRow >= rows ||
			checkCol < 0 || checkCol >= cols ||
			charMatrix[checkRow][checkCol] != char {
			return false
		}
	}

	return true
}

func searchForX_MAS(charMatrix [][]string) int {
	rows, cols := len(charMatrix), len(charMatrix[0])
	count := 0

	for startRow := 1; startRow < rows-1; startRow++ {
		for startCol := 1; startCol < cols-1; startCol++ {
			ms := 0
			ss := 0
			var mCords [][]int
			var sCords [][]int
			if charMatrix[startRow][startCol] != "A" {
				continue
			}
			//top left
			if charMatrix[startRow-1][startCol-1] == "M" {
				ms++
				mCords = append(mCords, []int{startRow - 1, startCol - 1})
			} else if charMatrix[startRow-1][startCol-1] == "S" {
				ss++
				sCords = append(sCords, []int{startRow - 1, startCol - 1})
			}
			//top right
			if charMatrix[startRow-1][startCol+1] == "M" {
				ms++
				mCords = append(mCords, []int{startRow - 1, startCol + 1})
			} else if charMatrix[startRow-1][startCol+1] == "S" {
				ss++
				sCords = append(sCords, []int{startRow - 1, startCol + 1})
			}
			//bottom left
			if charMatrix[startRow+1][startCol-1] == "M" {
				ms++
				mCords = append(mCords, []int{startRow + 1, startCol - 1})
			} else if charMatrix[startRow+1][startCol-1] == "S" {
				ss++
				sCords = append(sCords, []int{startRow + 1, startCol - 1})
			}
			//bottom right
			if charMatrix[startRow+1][startCol+1] == "M" {
				ms++
				mCords = append(mCords, []int{startRow + 1, startCol + 1})
			} else if charMatrix[startRow+1][startCol+1] == "S" {
				ss++
				sCords = append(sCords, []int{startRow + 1, startCol + 1})
			}
			if ms == 2 && ss == 2 {
				//check if y or x cords are equal for m
				if !(mCords[0][0] == mCords[1][0]) && !(mCords[0][1] == mCords[1][1]) {
					continue
				}
				//check if y or x cords are equal for s
				if !(sCords[0][0] == sCords[1][0]) && !(sCords[0][1] == sCords[1][1]) {
					continue
				}
				count++
			}
		}
	}

	return count
}

// 1948 is too high
func SolutionDay4() {
	charMatrix := readFile("./Day4/Day4.txt")
	/*testMatrix := []string{
		".M.S......",
		"..A..MSMS.",
		".M.S.MAA..",
		"..A.ASMSM.",
		".M.S.M....",
		"..........",
		"S.S.S.S.S.",
		".A.A.A.A..",
		"M.M.M.M.M.",
		"..........",
	}
	var charMatrix [][]string
	for i := 0; i < len(testMatrix); i++ {
		charMatrix = append(charMatrix, strings.Split(testMatrix[i], ""))
	}*/
	//fmt.Printf("%+v\n", charMatrix)
	occurences := searchForXMAS(charMatrix)
	fmt.Printf("Solution Day4 Part 1: %d\n", occurences)
	occurences2 := searchForX_MAS(charMatrix)
	fmt.Printf("Solution Day4 Part 2: %d\n", occurences2)
}
