package Day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(filename string) [][]string {
	ret := make([][]string, 0)
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

func getStartPosition(input [][]string) (int, int) {
	x, y := 0, 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == "^" {
				x = i
				y = j
			}
		}
	}
	return x, y
}

func print2DArray(input [][]string) {
	for _, value := range input {
		fmt.Printf("%v\n", value)
	}
	fmt.Println()
}

func walkThrough(input [][]string) [][]string {
	//find start pos
	x, y := getStartPosition(input)
	/*
	* -1 0 oben
	* 0 1 rechts
	* 1 0 unten
	* 0 -1 links
	 */
	direction := []int{-1, 0}
	i := 0
	for {
		/*if i > 500 {
			break
		}*/
		//fmt.Println("X, Y", x, y)
		//fmt.Printf("direction %v\n", direction)
		if x+direction[0] >= len(input) || x+direction[0] < 0 {
			//fmt.Println("out of range", x, y)
			break
		}
		if y+direction[1] >= len(input[x]) || y+direction[1] < 0 {
			//fmt.Println("out of range", x, y)
			break
		}
		if input[x+direction[0]][y+direction[1]] == "." || input[x+direction[0]][y+direction[1]] == "X" || input[x+direction[0]][y+direction[1]] == "^" {
			x = x + direction[0]
			y = y + direction[1]
			input[x][y] = "X"
		} else if input[x+direction[0]][y+direction[1]] == "#" {
			//fmt.Println("found obtacle")
			//rotate Vektor
			tmp := direction[0]
			direction[0] = direction[1]
			direction[1] = -tmp
		}
		i++
	}
	return input
}

func countX(input [][]string) int {
	sum := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == "X" || input[i][j] == "^" {
				sum++
			}
		}
	}
	return sum
}

func SolutionDay6() {
	input := readFile("./Day6/Day6.txt")
	//fmt.Printf("%v\n", input)
	walked := walkThrough(input)
	//print2DArray(walked)
	Xs := countX(walked)
	fmt.Printf("Solution Day6 Part 1: %d\n", Xs)
	/*sum2 := validateAllUpdates(false)
	fmt.Printf("Solution Day6 Part 2: %d\n", sum2)*/
}
