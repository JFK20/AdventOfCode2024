package Day10

import (
	"AdventOfCode/helper"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) map[helper.Vector2D]int {
	ret := make(map[helper.Vector2D]int)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Problem opening File", err)
		return nil
	}
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		splitted := strings.Split(scanner.Text(), "")
		for i, v := range splitted {
			num, _ := strconv.Atoi(v)
			ret[helper.Vector2D{X: i, Y: y}] = num
		}
		y++
	}

	return ret
}

func walk(topos map[helper.Vector2D]int, startPos helper.Vector2D, startValue int, visited map[helper.Vector2D]bool) int {
	if startValue == 9 {
		if visited[startPos] {
			return 0
		}
		visited[startPos] = true
		return 1
	}
	found := 0
	allNeighbours := startPos.GetAllNeighbours()
	newValue := startValue + 1
	for _, neighbour := range allNeighbours {
		if topos[neighbour] == newValue {
			found += walk(topos, neighbour, newValue, visited)
		} else {
			continue
		}
	}
	return found
}

func walkAll(topos map[helper.Vector2D]int) int {
	sum := 0

	for pos, v := range topos {
		if v == 0 {
			visited := make(map[helper.Vector2D]bool)
			sum += walk(topos, pos, v, visited)
			fmt.Printf("%d for pos X:%d Y:%d\n", sum, pos.X, pos.Y)
		}
	}
	return sum
}

func SolutionDay10() {
	input := readFile("./Day10/Day10.txt")
	fmt.Println(input)
	sum := walkAll(input)
	fmt.Printf("Solution Day9 Part 1: %d\n", sum)
	//fmt.Printf("Solution Day9 Part 2: %d\n", checksum2)
}
