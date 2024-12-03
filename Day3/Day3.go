package Day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFile(filename string) string {
	ret := ""
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Problem opening File", err)
		return "nil"
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret += scanner.Text()
	}
	return ret
}

func calculateSum(matches [][]string) int {
	sum := 0
	for _, value := range matches {
		var numbers []string = strings.Split(value[1], ",")
		left, _ := strconv.Atoi(numbers[0])
		right, _ := strconv.Atoi(numbers[1])
		sum += left * right
	}
	return sum
}

func getDoandDonts(input string) []int {
	do := regexp.MustCompile(`do\(\)`)
	donts := regexp.MustCompile(`don't\(\)`)
	doIndex := make([]int, 0)
	dontIndex := make([]int, 0)
	for _, value := range do.FindAllStringIndex(input, -1) {
		doIndex = append(doIndex, value[1])
	}
	for _, value := range donts.FindAllStringIndex(input, -1) {
		dontIndex = append(dontIndex, value[0])
	}
	laufen := []int{0, dontIndex[0]}
	for i := 0; i < len(dontIndex); i++ {
		//the new doIndex should be greater then the last appended dontIndex
		if laufen[len(laufen)-1] < doIndex[i] {
			laufen = append(laufen, doIndex[i])
		} else {
			continue
		}
		flag := false
		for j := i; j < len(dontIndex); j++ {
			//check that the dontIndex is greater then the doIndex
			if dontIndex[j] > doIndex[i] {
				flag = true
				laufen = append(laufen, dontIndex[j])
				break
			}
		}
		//if for the last doIndex no matching dont was found run till end
		if !flag {
			laufen = append(laufen, len(input))
		}
	}
	return laufen
}

func manipInput(input string, index []int) string {
	ret := ""
	for i := 0; i < len(index); i += 2 {
		ret += input[index[i]:index[i+1]]
	}
	return ret
}

func SolutionDay3() {
	input := readFile("./Day3/Day3.txt")
	re := regexp.MustCompile(`mul\((\d+,\d+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	sum := calculateSum(matches)
	fmt.Printf("Solution Day3 Part 1: %d\n", sum)
	toCut := getDoandDonts(input)
	newInput := manipInput(input, toCut)
	matches2 := re.FindAllStringSubmatch(newInput, -1)
	sum2 := calculateSum(matches2)
	fmt.Printf("Solution Day3 Part 2: %d\n", sum2)
}
