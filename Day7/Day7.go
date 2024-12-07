package Day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type mathProblem struct {
	solution int
	numbers  []int
}

func newMathProblem(sol int) mathProblem {
	return mathProblem{solution: sol, numbers: make([]int, 0)}
}

func (problem *mathProblem) addNumber(value int) {
	problem.numbers = append(problem.numbers, value)
}

func readFile(filename string) []mathProblem {
	ret := make([]mathProblem, 0)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Problem opening File", err)
		return nil
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var problem mathProblem
		splited := strings.Split(scanner.Text(), ":")
		splited[0] = strings.TrimSpace(splited[0])
		splited[1] = strings.TrimSpace(splited[1])
		num, _ := strconv.Atoi(splited[0])
		problem = newMathProblem(num)
		for _, val := range strings.Split(splited[1], " ") {
			val = strings.TrimSpace(val)
			num, _ = strconv.Atoi(val)
			problem.addNumber(num)
		}
		ret = append(ret, problem)
	}
	return ret
}

func generateCombinations(symbols []string, length int) [][]string {
	result := [][]string{}

	var generateCombos func(current string, remaining int)
	generateCombos = func(current string, remaining int) {
		if remaining == 0 {
			splited := strings.Split(current, "")
			result = append(result, splited)
			return
		}

		for _, symbol := range symbols {
			generateCombos(current+symbol, remaining-1)
		}
	}

	generateCombos("", length)

	return result
}

func solve(numbers []int, combination []string) int {
	calcResult := numbers[0]
	for i, _ := range combination {
		if combination[i] == "+" {
			calcResult += numbers[i+1]
		} else if combination[i] == "*" {
			calcResult *= numbers[i+1]
		} else if combination[i] == "|" {
			leftStr := strconv.Itoa(calcResult)
			rightStr := strconv.Itoa(numbers[i+1])
			calcResult, _ = strconv.Atoi(leftStr + rightStr)
		}
	}
	return calcResult
}

func isSolvable(problem *mathProblem, operators []string) bool {
	lenght := len(problem.numbers) - 1
	possibilities := generateCombinations(operators, lenght)
	for _, possibilty := range possibilities {
		sol := solve(problem.numbers, possibilty)
		if sol == problem.solution {
			return true
		}
	}
	return false
}

func solveAll(problems []mathProblem, operators []string) int {
	sum := 0
	for _, problem := range problems {
		if isSolvable(&problem, operators) {
			sum += problem.solution
		}
	}
	return sum
}

func SolutionDay7() {
	input := readFile("./Day7/Day7.txt")
	operators := []string{"+", "*"}
	sum := solveAll(input, operators)
	fmt.Printf("Solution Day7 Part 1: %d\n", sum)
	operators2 := []string{"+", "*", "|"}
	sum2 := solveAll(input, operators2)
	fmt.Printf("Solution Day7 Part 2: %d\n", sum2)
}
