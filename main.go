package main

import (
	"AdventOfCode/Day8"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	/*Day1.SolutionDay1()
	Day2.SolutionDay2()
	Day3.SolutionDay3()
	Day4.SolutionDay4()
	Day5.SolutionDay5()
	Day6.SolutionDay6()
	Day7.SolutionDay7()*/
	Day8.SolutionDay8()
	elapsed := time.Since(start)
	fmt.Printf("AoC took %s", elapsed)
}
