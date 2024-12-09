package Day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

var inputLength int

func readFile(filename string) string {
	ret := ""
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Problem opening File", err)
		return ""
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret += scanner.Text()
	}
	inputLength = len(ret)
	return ret
}

func getLength(input string) int {
	sum := 0
	for _, val := range input {
		num, _ := strconv.Atoi(string(val))
		sum += num
	}
	return sum
}

func convertData(input string) []int {

	ret := make([]int, getLength(input))
	id := 0
	retIndex := 0
	for i := 0; i < len(input); i++ {
		num, _ := strconv.Atoi(string(input[i]))
		if i%2 == 0 {
			for j := 0; j < num; j++ {
				ret[retIndex] = id
				retIndex++
			}
			id++
		} else {
			for j := 0; j < num; j++ {
				ret[retIndex] = -1
				retIndex++
			}
		}
	}
	return ret
}

func compressData(input []int) []int {
	ret := make([]int, len(input))

	left, right := 0, len(input)-1
	for left <= right {
		if input[left] != -1 {
			ret[left] = input[left]
			left++
		} else if input[left] == -1 && input[right] != -1 {
			ret[left] = input[right]
			right--
			left++
		} else {
			right--
		}
	}
	for i := left; i < len(ret); i++ {
		ret[i] = -1
	}

	return ret
}

func getBlockLength(input []int) map[int][]int {
	length := len(input)

	blockLengths := map[int][]int{}

	for i := length - 1; i >= 0; i-- {
		blockLength := 0
		foundNum := input[i]
		if foundNum == -1 {
			continue
		}
		for j := length - 1; j >= 0; j-- {
			if foundNum == input[j] {
				blockLengths[foundNum] = append(blockLengths[foundNum], j)
				blockLength++
				i--
			}
		}
	}
	return blockLengths
}

func getFreeSpace(input []int) map[int][]int {
	length := len(input)

	freeSpaces := map[int][]int{}

	for i := length - 1; i >= 0; i-- {
		blockLength := 0
		foundNum := input[i]
		if foundNum != -1 {
			continue
		}
		for j := i; j >= 0; j-- {
			if foundNum == input[j] {
				blockLength++
				i--
			} else {
				freeSpaces[blockLength] = append(freeSpaces[blockLength], j+1)
				slices.Sort(freeSpaces[blockLength])
				break
			}
		}
	}
	return freeSpaces
}

func compressDataBlock(input []int) []int {
	ret := make([]int, len(input))
	ret = input
	blockLengths := getBlockLength(input)
	freeSpaces := getFreeSpace(input)
	mapLength := (inputLength - 1) / 2
	for i := mapLength; i > 0; i-- {
		currentBlockLength := len(blockLengths[i])
		smallestPossibleFit := -1
		for j := currentBlockLength; j <= 9; j++ {
			if len(freeSpaces[j]) > 0 {
				smallestPossibleFit = freeSpaces[j][0]
				freeSpaces[j] = freeSpaces[j][1:]
				remaningSpace := j - currentBlockLength
				newIndex := smallestPossibleFit + currentBlockLength
				freeSpaces[remaningSpace] = append(freeSpaces[remaningSpace], newIndex)
				slices.Sort(freeSpaces[remaningSpace])
			}
		}
		if smallestPossibleFit == -1 {
			continue
		}
		for j := smallestPossibleFit; j < smallestPossibleFit+currentBlockLength; j++ {
			ret[j] = i
		}
		for j := 0; j < len(blockLengths[i]); j++ {
			ret[blockLengths[i][j]] = -1
		}
	}

	return ret
}

func calculateChecksum(input []int) uint64 {
	var check uint64 = 0
	for i := 0; i < len(input); i++ {
		if input[i] != -1 {
			num := input[i]
			check += uint64(num * i)
		}
	}
	return check
}

// 14477716715259 to high
func SolutionDay9() {
	input := readFile("./Day9/Day9Test.txt")
	data := convertData(input)
	fmt.Println(data)
	compressedData := compressData(data)
	fmt.Println(compressedData)
	checksum := calculateChecksum(compressedData)
	fmt.Printf("Solution Day9 Part 1: %d\n", checksum)
	compressedDataBlock := compressDataBlock(data)
	fmt.Println(compressedDataBlock)
	checksum2 := calculateChecksum(compressedDataBlock)
	fmt.Printf("Solution Day9 Part 2: %d\n", checksum2)
}
