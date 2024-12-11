package Day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

// der erste key ist der stone
var memo = make(map[int]map[int][]int) // Cache resulting stones for stone and iteration
var memoMu sync.Mutex

func readFile(filename string) []int {
	ret := make([]int, 0)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Problem opening File", err)
		return nil
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitted := strings.Split(scanner.Text(), " ")
		for _, v := range splitted {
			num, _ := strconv.Atoi(v)
			ret = append(ret, num)
		}
	}

	return ret
}

func rules(stone int) []int {
	ret := make([]int, 0)
	if stone == 0 {
		ret = append(ret, 1)
		return ret
	}
	stoneString := strconv.Itoa(stone)
	if len(stoneString)%2 == 0 {
		num, _ := strconv.Atoi(stoneString[:len(stoneString)/2])
		ret = append(ret, num)
		num, _ = strconv.Atoi(stoneString[len(stoneString)/2:])
		ret = append(ret, num)
		return ret
	} else {
		stone *= 2024
		ret = append(ret, stone)
		return ret
	}
}

func findAllEndStones(stones []int, iters int) int {
	amountByValue := make(map[int]int)
	for _, v := range stones {
		amountByValue[v] = amountByValue[v] + 1
	}
	for i := 0; i < iters; i++ {
		resultMap := make(map[int]int)
		for k, v := range amountByValue {
			amount := v

			a := rules(k)
			for _, val := range a {
				resultMap[val] = resultMap[val] + amount
			}
		}
		amountByValue = resultMap
	}
	ret := 0
	for _, v := range amountByValue {
		ret += v
	}
	return ret
}

func SolutionDay11() {
	input := readFile("./Day11/Day11.txt")
	sum := findAllEndStones(input, 25)
	fmt.Printf("Solution Day11 Part 1: %d\n", sum)
	sum2 := findAllEndStones(input, 75)
	fmt.Printf("Solution Day11 Part 2: %d\n", sum2)
}
