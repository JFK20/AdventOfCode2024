package main

import (
	"AdventOfCode/Day13"
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
	Day7.SolutionDay7()
	Day8.SolutionDay8()
	Day9.SolutionDay9()
	Day10.SolutionDay10()
	Day11.SolutionDay11()
	Day12.SolutionDay12()*/
	Day13.SolutionDay13()
	elapsed := time.Since(start)
	fmt.Printf("AoC took %s", elapsed)
}
