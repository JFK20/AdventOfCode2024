package Day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

type Node struct {
	value      int
	neighbours []*Node
	visited    bool
}

func newNode(value int) *Node {
	return &Node{value, make([]*Node, 0), false}
}

func readRules(filename string, needNumber []int) map[int]*Node {
	fileData := make([]string, 0)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Problem opening File", err)
		return nil
	}
	scanner := bufio.NewScanner(file)
	nodes := make(map[int]*Node)
	for scanner.Scan() {
		fileData = append(fileData, scanner.Text())
	}
	for _, line := range fileData {
		tmp := strings.Split(line, "|")
		left, _ := strconv.Atoi(tmp[0])
		right, _ := strconv.Atoi(tmp[1])
		if slices.Contains(needNumber, left) && slices.Contains(needNumber, right) {
			nodes[left] = newNode(left)
			nodes[right] = newNode(right)
		}
	}

	for _, line := range fileData {
		tmp := strings.Split(line, "|")
		left, _ := strconv.Atoi(tmp[0])
		right, _ := strconv.Atoi(tmp[1])
		if slices.Contains(needNumber, left) && slices.Contains(needNumber, right) {
			nodes[left].neighbours = append(nodes[left].neighbours, nodes[right])
		}
	}
	return nodes
}

func findPath(nodes map[int]*Node, startNode *Node) []int {
	nodeList := make([]*Node, 0, len(nodes))
	for _, node := range nodes {
		nodeList = append(nodeList, node)
	}

	// Track visited nodes
	visited := make(map[*Node]bool)

	// Path to store the final route
	path := []int{}

	// Backtracking function
	var backtrack func(current *Node, depth int) bool
	backtrack = func(current *Node, depth int) bool {
		// Mark current node as visited
		visited[current] = true
		path = append(path, current.value)

		// If we've visited all nodes, we found a path
		if depth == len(nodes)-1 {
			return true
		}

		// Try each neighbour
		for _, neighbour := range current.neighbours {
			if !visited[neighbour] {
				// Recursively explore this path
				if backtrack(neighbour, depth+1) {
					return true
				}
			}
		}

		// If no path found, backtrack
		visited[current] = false
		path = path[:len(path)-1]
		return false
	}

	if backtrack(startNode, 0) {
		return path
	}

	// No path found
	return nil
}

func remove(slice []int, s int) []int {
	index := slices.Index(slice, s)
	return append(slice[:index], slice[index+1:]...)
}

func uniqueAppend(slice []*Node, node *Node) []*Node {
	if !slices.Contains(slice, node) {
		slice = append(slice, node)
	}
	return slice
}

func findStartNode(nodes map[int]*Node) []int {
	neighbours := make([]*Node, 0)
	var possibility []int

	for key, _ := range nodes {
		possibility = append(possibility, key)
	}

	for _, node := range nodes {
		for _, neighbour := range node.neighbours {
			neighbours = uniqueAppend(neighbours, neighbour)
		}
	}

	for key, node := range nodes {
		if slices.Contains(neighbours, node) {
			possibility = remove(possibility, key)
		}
	}
	return possibility
}

func readPages(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Problem opening File", err)
		return nil
	}
	scanner := bufio.NewScanner(file)
	var pages [][]int
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), ",")
		var page []int
		for _, value := range tmp {
			val, _ := strconv.Atoi(value)
			page = append(page, val)
		}
		pages = append(pages, page)
	}
	return pages
}

func checkOrder(page []int, order []int) bool {
	if len(page) != len(order) {
		return false
	}
	for i := 0; i < len(page); i++ {
		if page[i] != order[i] {
			return false
		}
	}
	return true
}

func validateUpdate(update []int, flag bool) int {
	filename := "./Day5/Day5Rules.txt"
	rules := readRules(filename, update)
	startNode := findStartNode(rules)
	path := findPath(rules, rules[startNode[0]])
	if checkOrder(update, path) && flag {
		mid := len(path) / 2
		return update[mid]
	}
	if !checkOrder(update, path) && !flag {
		mid := len(path) / 2
		return path[mid]
	}
	return 0
}

func validateAllUpdates(flag bool) int {
	pages := readPages(filepath.Join("./Day5/Day5Pages.txt"))
	sum := 0
	for _, page := range pages {
		num := validateUpdate(page, flag)
		sum += num
	}
	return sum
}

func SolutionDay5() {
	sum := validateAllUpdates(true)
	fmt.Printf("Solution Day5 Part 1: %d\n", sum)
	sum2 := validateAllUpdates(false)
	fmt.Printf("Solution Day5 Part 2: %d\n", sum2)
}
