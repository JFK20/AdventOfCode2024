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

var memo = make(map[int]int)

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

func findEndStones(stone int, iters int) int {
	stones := make([]int, 0)
	stonesAfter := make([]int, 0)
	stones = append(stones, stone)
	for i := 0; i < iters; i++ {
		for j := 0; j < len(stones); j++ {
			/*if memo[stones[j]] != 0 {
				return memo[stones[j]]
			}*/
			stonesAfter = append(stonesAfter, rules(stones[j])...)
			//memo[stones[j]] = len(stonesAfter)
		}
		stones = stonesAfter
		stonesAfter = make([]int, 0)
	}
	return len(stones)
}

func findAllEndStones(stones []int, iters int) int {
	sum := 0
	endStones := make([]int, 0)

	var wg sync.WaitGroup
	mu := sync.Mutex{}
	results := make(chan []int, len(stones))
	counts := make(chan int, len(stones))

	for _, stone := range stones {
		wg.Add(1)
		go func(stone int) {
			defer wg.Done()
			count := findEndStones(stone, iters)
			counts <- count
		}(stone)
	}

	// Close channels once all goroutines are done.
	go func() {
		wg.Wait()
		close(counts)
		close(results)
	}()

	// Aggregate results from channels.
	for count := range counts {
		mu.Lock()
		sum += count
		mu.Unlock()
	}

	for newStones := range results {
		mu.Lock()
		endStones = append(endStones, newStones...)
		mu.Unlock()
	}

	return sum
}

func SolutionDay11() {
	input := readFile("./Day11/Day11Test.txt")
	sum := findAllEndStones(input, 25)
	fmt.Printf("Solution Day11 Part 1: %d\n", sum)
	//sum, endStones = findAllEndStones(endStones, 25)
	//fmt.Println("50 Iters Completed")
	//sum, endStones = findAllEndStones(endStones, 25)
	//fmt.Printf("Solution Day11 Part 2: %d\n", sum)
}
