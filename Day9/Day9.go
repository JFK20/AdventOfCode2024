package Day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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

func convertData(input string) string {

	ret := make([]rune, getLength(input))
	var id rune = '0'
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
				ret[retIndex] = '.'
				retIndex++
			}
		}
	}
	return string(ret)
}

func compressData(input string) string {
	ret := make([]rune, len(input))

	left, right := 0, len(input)-1
	for left <= right {
		if input[left] != '.' {
			ret[left] = rune(input[left])
			left++
		} else if input[left] == '.' && input[right] != '.' {
			ret[left] = rune(input[right])
			right--
			left++
		} else {
			right--
		}
	}
	for i := left; i < len(ret); i++ {
		ret[i] = '.'
	}

	return string(ret)
}

func calculateChecksum(input string) uint64 {
	var check uint64 = 0
	for i := 0; i < len(input); i++ {
		if input[i] != '.' {
			num := input[i] - '0'
			check += uint64(int(num) * i)
		}
	}
	return check
}

// 90167081070 to low
// 210728992977 to low
// 5027798647049 to low
func SolutionDay9() {
	input := readFile("./Day9/Day9.txt")
	data := convertData(input)
	//fmt.Println(data)
	compressedData := compressData(data)
	//fmt.Println(compressedData)
	checksum := calculateChecksum(compressedData)
	fmt.Printf("Solution Day9 Part 1: %d\n", checksum)

	//fmt.Printf("Solution Day8 Part 2: %d\n", amount2)
}
