package Day13

import (
	"AdventOfCode/mathUtil"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type automate struct {
	m     mathUtil.Matrix[float64]
	price mathUtil.Vector2D[int]
}

func readFile(filename string, factor int) []automate {
	ret := make([]automate, 0)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Problem opening File", err)
		return nil
	}
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`[XY][+=](\d+)`)
	i := 0
	ma := mathUtil.Matrix[float64]{Prefactor: 1}
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		num1, _ := strconv.Atoi(matches[0][1])
		num2, _ := strconv.Atoi(matches[1][1])
		if i == 0 {
			ma.A = float64(num1)
			ma.C = float64(num2)
			i++
		} else if i == 1 {
			ma.B = float64(num1)
			ma.D = float64(num2)
			i++
		} else if i == 2 {
			x, _ := strconv.Atoi(matches[0][1])
			y, _ := strconv.Atoi(matches[1][1])
			ret = append(ret, automate{m: ma, price: mathUtil.Vector2D[int]{X: x + factor, Y: y + factor}})
			ma = mathUtil.Matrix[float64]{Prefactor: 1}
			i = 0
		}

	}
	return ret
}

func getSolution(auto automate) mathUtil.Vector2D[float64] {
	mat := auto.m
	mat.Invert()
	sol := mat.Multiply(mathUtil.Vector2D[float64]{X: float64(auto.price.X), Y: float64(auto.price.Y)})
	return sol
}

func getAllSolution(autos []automate) []mathUtil.Vector2D[int] {
	solutions := make([]mathUtil.Vector2D[int], 0)
	for _, v := range autos {
		sol := getSolution(v)
		solInt, err := sol.ConvertToInt()
		if err != nil {
			continue
		}
		solutions = append(solutions, solInt)
	}
	return solutions
}

func calcCost(sols []mathUtil.Vector2D[int]) int {
	cost := 0
	for _, v := range sols {
		cost += v.X * 3
		cost += v.Y * 1
	}
	return cost
}

// 34731664719093 too low
func SolutionDay13() {
	input := readFile("./Day13/Day13Text.txt", 0)
	solution := getAllSolution(input)
	fmt.Printf("Solution Day13 Part 1: %d\n", calcCost(solution))
	input2 := readFile("./Day13/Day13Text.txt", 10e12)
	solution2 := getAllSolution(input2)
	fmt.Printf("Solution Day13 Part 1: %d\n", calcCost(solution2))
}
