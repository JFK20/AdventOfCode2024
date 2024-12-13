package Day10

import (
	"AdventOfCode/mathUtil"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) map[mathUtil.Vector2D[int]]int {
	ret := make(map[mathUtil.Vector2D[int]]int)
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
			ret[mathUtil.Vector2D[int]{X: i, Y: y}] = num
		}
		y++
	}

	return ret
}

func walk(topos map[mathUtil.Vector2D[int]]int, startPos mathUtil.Vector2D[int], startValue int, visited map[mathUtil.Vector2D[int]]bool, flag bool) int {
	if flag {
		visited = make(map[mathUtil.Vector2D[int]]bool)
	}
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
			found += walk(topos, neighbour, newValue, visited, flag)
		} else {
			continue
		}
	}
	return found
}

func walkAll(topos map[mathUtil.Vector2D[int]]int, flag bool) int {
	sum := 0

	for pos, v := range topos {
		if v == 0 {
			visited := make(map[mathUtil.Vector2D[int]]bool)
			sum += walk(topos, pos, v, visited, flag)
		}
	}
	return sum
}

func SolutionDay10() {
	input := readFile("./Day10/Day10.txt")
	sum := walkAll(input, false)
	fmt.Printf("Solution Day10 Part 1: %d\n", sum)
	sum2 := walkAll(input, true)
	fmt.Printf("Solution Day10 Part 2: %d\n", sum2)
}
