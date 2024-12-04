package Day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type charInfo struct {
	index      int
	char       string
	neighbours []int
}

func newCharInfo(i int, j int, mlength int, char string) charInfo {
	c := charInfo{index: i*mlength + j, char: char}
	tl := (i-1)*mlength + j - 1
	t := (i-1)*mlength + j
	tr := (i-1)*mlength + j + 1
	l := i*mlength + j - 1
	r := i*mlength + j + 1
	bl := (i+1)*mlength + j - 1
	b := (i+1)*mlength + j
	br := (i+1)*mlength + j + 1
	c.neighbours = []int{tl, t, tr, r, br, b, bl, l}
	return c
}

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

func searchAlgo(charInfoArray []charInfo, info charInfo, direction int, nextChar string) int {
	if info.char == "X" {
		if direction == -1 {
			for i := 0; i < len(info.neighbours); i++ {
				if info.neighbours[i] < 0 || info.neighbours[i] >= len(charInfoArray) {
					continue
				}
				foundChar := charInfoArray[info.neighbours[i]].char
				if foundChar == nextChar {
					return searchAlgo(charInfoArray, charInfoArray[info.neighbours[i]], -1, "A")
				}
			}
		} else {
			if info.neighbours[direction] < 0 || info.neighbours[direction] >= len(charInfoArray) {
				return 0
			}
			foundChar := charInfoArray[info.neighbours[direction]].char
			if foundChar == nextChar {
				return searchAlgo(charInfoArray, charInfoArray[info.neighbours[direction]], direction, "A")
			}
		}
	}
	if info.char == "M" {
		if direction == -1 {
			for i := 0; i < len(info.neighbours); i++ {
				if info.neighbours[i] < 0 || info.neighbours[i] >= len(charInfoArray) {
					continue
				}
				foundChar := charInfoArray[info.neighbours[i]].char
				if foundChar == nextChar {
					return searchAlgo(charInfoArray, charInfoArray[info.neighbours[i]], -1, "S")
				}
			}
		} else {
			if info.neighbours[direction] < 0 || info.neighbours[direction] >= len(charInfoArray) {
				return 0
			}
			foundChar := charInfoArray[info.neighbours[direction]].char
			if foundChar == nextChar {
				return searchAlgo(charInfoArray, charInfoArray[info.neighbours[direction]], direction, "S")
			}
		}

	}
	if info.char == "A" {
		if direction == -1 {
			for i := 0; i < len(info.neighbours); i++ {
				if info.neighbours[i] < 0 || info.neighbours[i] >= len(charInfoArray) {
					continue
				}
				foundChar := charInfoArray[info.neighbours[i]].char
				if foundChar == nextChar {
					return 1
				}
			}
		} else {
			if info.neighbours[direction] < 0 || info.neighbours[direction] >= len(charInfoArray) {
				return 0
			}
			foundChar := charInfoArray[info.neighbours[direction]].char
			if foundChar == nextChar {
				return 1
			}
		}
	}
	return 0
}

func createCharInfoArray(charMatrix [][]string) int {
	var charInfoArray []charInfo
	for i := 0; i < len(charMatrix); i++ {
		for j := 0; j < len(charMatrix[i]); j++ {
			var charI = newCharInfo(i, j, len(charMatrix[i]), charMatrix[i][j])
			charInfoArray = append(charInfoArray, charI)
		}
	}
	sum := 0
	for i := 0; i < len(charInfoArray); i++ {
		if charInfoArray[i].char == "X" {
			for j := 0; j < 8; j++ {
				sum += searchAlgo(charInfoArray, charInfoArray[i], j, "M")
			}
		}
	}
	return sum
}

// 3387 is to high 2567 is to high
func SolutionDay4() {
	charMatrix := readFile("./Day4/Day4.txt")
	/*testMatrix := []string{
		"....XXMAS.",
		".SAMXMS...",
		"...S..A...",
		"..A.A.MS.X",
		"XMASAMX.MM",
		"X.....XA.A",
		"S.S.S.S.SS",
		".A.A.A.A.A",
		"..M.M.M.MM",
		".X.X.XMASX",
	}
	var charMatrix [][]string
	for i := 0; i < len(testMatrix); i++ {
		charMatrix = append(charMatrix, strings.Split(testMatrix[i], ""))
	}*/
	//fmt.Printf("%+v\n", charMatrix)
	occurences := createCharInfoArray(charMatrix)
	fmt.Println(occurences)
}
