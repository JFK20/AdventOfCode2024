package Day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func stringAppendArray(leftside []int, rightside []int, line string) ([]int, []int) {
	stringValues := strings.Split(line, " ")
	for _, str := range stringValues {
		strings.ReplaceAll(str, "\n", "")
		strings.ReplaceAll(str, " ", "")
	}
	val, _ := strconv.Atoi(stringValues[0])
	leftside = append(leftside, val)
	val, _ = strconv.Atoi(stringValues[len(stringValues)-1])
	rightside = append(rightside, val)
	return leftside, rightside
}

func readFile(filename string) ([]int, []int) {
	leftside := make([]int, 0)
	rightside := make([]int, 0)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Problem opening File", err)
		return nil, nil
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		leftside, rightside = stringAppendArray(leftside, rightside, scanner.Text())
	}
	file.Close()
	return leftside, rightside
}

func calulateDistance(leftside []int, rightside []int) int {
	sort.Ints(leftside)
	sort.Ints(rightside)
	sum := 0
	if len(leftside) != len(rightside) {
		log.Fatal("Left side and right side don't match")
		return -1
	}
	for index, _ := range leftside {
		tmp := leftside[index] - rightside[index]
		if tmp < 0 {
			tmp = -tmp
		}
		sum += tmp
	}
	return sum
}

func calculateSimilarity(leftside []int, rightside []int) int {
	sort.Ints(leftside)
	sort.Ints(rightside)
	score := 0
	if len(leftside) != len(rightside) {
		log.Fatal("Left side and right side don't match")
		return -1
	}
	for _, leftValue := range leftside {
		occurences := 0
		for _, rightValue := range rightside {
			if leftValue == rightValue {
				occurences++
			}
		}
		score += occurences * leftValue
	}
	return score
}

func SolutionDay1() {
	leftside, rightside := readFile("./Day1/Day1.txt")
	distance := calulateDistance(leftside, rightside)
	fmt.Printf("Solution Day1 Part 1: %d\n", distance)
	score := calculateSimilarity(leftside, rightside)
	fmt.Printf("Solution Day1 Part 2: %d\n", score)
}
