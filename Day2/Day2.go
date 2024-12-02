package Day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func stringAppendArray(line string) []int {
	intSlice := make([]int, 0)
	stringValues := strings.Split(line, " ")
	for _, str := range stringValues {
		str = strings.ReplaceAll(str, "\n", "")
		str = strings.ReplaceAll(str, " ", "")
		val, _ := strconv.Atoi(str)
		intSlice = append(intSlice, val)
	}
	return intSlice
}

func readFile(filename string) [][]int {
	lines := make([][]int, 0)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Problem opening File", err)
		return nil
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, stringAppendArray(scanner.Text()))
	}
	file.Close()
	return lines
}

func isArraySafe(val []int, order int) bool {
	if len(val) <= 1 {
		return true
	}

	if order == 0 {
		return false
	}

	for i := 1; i < len(val); i++ {
		if order == 1 {
			// Ascending order check
			if val[i-1] >= val[i] || val[i]-val[i-1] > 3 {
				return false
			}
		} else if order == -1 {
			// Descending order check
			if val[i-1] <= val[i] || val[i-1]-val[i] > 3 {
				return false
			}
		}
	}
	return true
}

func checkArray(val []int, order int, flag bool) bool {
	if isArraySafe(val, order) {
		return true
	}

	//brute force
	if flag {
		for i := 0; i < len(val); i++ {
			newVal := make([]int, 0, len(val)-1)
			newVal = append(newVal, val[:i]...)
			newVal = append(newVal, val[i+1:]...)

			newOrder := getOrder(newVal)
			if isArraySafe(newVal, newOrder) {
				return true
			}
		}
	}

	return false
}

func getOrder(arr []int) int {
	if arr[0] == arr[1] {
		//first Elements Equal
		return 0
	} else if arr[0] > arr[1] {
		//Descending
		return -1
	} else {
		//Ascending
		return 1
	}
}

func checkOrder(lines [][]int, flag bool) int {
	hits := 0
	for _, val := range lines {
		order := getOrder(val)
		if checkArray(val, order, flag) {
			hits++
		}
	}
	return hits
}

func SolutionDay2() {
	linesArray := readFile("./Day2/Day2.txt")
	/*linesArray := [][]int{
		{19, 1, 2, 3, 5},
		{1, 2, 3, 5, 19},
		{1, 2, 19, 3, 5},
		{7, 6, 19, 4, 2},
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9}, //error to big gap
		{9, 7, 6, 2, 1}, //error to big gap
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}*/
	hits := checkOrder(linesArray, false)
	fmt.Printf("Solution Day2 Part 1: %d\n", hits)
	hits = checkOrder(linesArray, true)
	fmt.Printf("Solution Day2 Part 2: %d\n", hits)
}
