package Day14

import (
	"AdventOfCode/mathUtil"
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

type robot struct {
	pos mathUtil.Vector2D[int]
	vel mathUtil.Vector2D[int]
}

func (r *robot) inQuadrant(start mathUtil.Vector2D[int], end mathUtil.Vector2D[int]) bool {
	if r.pos.X < start.X {
		return false
	}
	if r.pos.Y < start.Y {
		return false
	}
	if r.pos.X > end.X {
		return false
	}
	if r.pos.Y > end.Y {
		return false
	}
	return true
}

func toString(robots []robot, bounds mathUtil.Vector2D[int]) string {
	// Initialize the grid
	matrix := make([][]rune, bounds.Y)
	for y := 0; y < bounds.Y; y++ {
		row := make([]rune, bounds.X)
		for x := 0; x < bounds.X; x++ {
			row[x] = '.' // Fill empty positions with '.'
		}
		matrix[y] = row
	}

	// Place robots on the grid
	for _, r := range robots {
		if r.pos.Y >= 0 && r.pos.Y < bounds.Y && r.pos.X >= 0 && r.pos.X < bounds.X {
			matrix[r.pos.Y][r.pos.X] = '#' // Place '#' where the robot is
		}
	}

	// Convert the grid to a string
	var result string
	for _, row := range matrix {
		result += string(row) + "\n" // Convert each row to a string
	}
	return result
}

func readFile(filename string) []robot {
	ret := make([]robot, 0)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Problem opening File", err)
		return nil
	}
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`p=(-?\d+),(-?\d+)\sv=(-?\d+),(-?\d+)`)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		posX, _ := strconv.Atoi(matches[0][1])
		posY, _ := strconv.Atoi(matches[0][2])
		velX, _ := strconv.Atoi(matches[0][3])
		velY, _ := strconv.Atoi(matches[0][4])
		ro := robot{pos: mathUtil.Vector2D[int]{X: posX, Y: posY}, vel: mathUtil.Vector2D[int]{X: velX, Y: velY}}
		ret = append(ret, ro)
	}
	return ret
}

func calcMovment(rob robot, sec int, bounds mathUtil.Vector2D[int]) robot {
	for i := 0; i < sec; i++ {
		newPos := mathUtil.Vector2D[int]{X: rob.pos.X + rob.vel.X, Y: rob.pos.Y + rob.vel.Y}
		if !newPos.IsInBounds(bounds) {
			rob.pos = newPos
		}
		if newPos.X < 0 {
			newPos.X += bounds.X
		}
		if newPos.Y < 0 {
			newPos.Y += bounds.Y
		}
		newPos.X = newPos.X % bounds.X
		newPos.Y = newPos.Y % bounds.Y
		rob.pos = newPos
	}
	return rob
}

func moveAllRobots(robots []robot, sec int, bounds mathUtil.Vector2D[int]) []robot {
	for i, _ := range robots {
		robots[i] = calcMovment(robots[i], sec, bounds)
	}
	return robots
}

func getQuadrants(bounds mathUtil.Vector2D[int]) []mathUtil.Vector2D[int] {
	upperleft := mathUtil.Vector2D[int]{X: bounds.X/2 - 1, Y: bounds.Y/2 - 1}
	upperright := mathUtil.Vector2D[int]{X: bounds.X, Y: bounds.Y/2 - 1}
	lowerleft := mathUtil.Vector2D[int]{X: bounds.X/2 - 1, Y: bounds.Y}
	lowerright := mathUtil.Vector2D[int]{X: bounds.X, Y: bounds.Y}
	return []mathUtil.Vector2D[int]{upperleft, upperright, lowerleft, lowerright}
}

func getRobotsInQuadarants(quadrants []mathUtil.Vector2D[int], robots []robot) int {
	upperleft := 0
	upperright := 0
	lowerleft := 0
	lowerright := 0
	for _, r := range robots {
		//upper left
		if r.inQuadrant(mathUtil.Vector2D[int]{X: 0, Y: 0}, quadrants[0]) {
			upperleft++
		}
		//upright
		if r.inQuadrant(mathUtil.Vector2D[int]{quadrants[0].X + 2, 0}, quadrants[1]) {
			upperright++
		}
		//lower left
		if r.inQuadrant(mathUtil.Vector2D[int]{0, quadrants[0].Y + 2}, quadrants[2]) {
			lowerleft++
		}
		//lower right
		if r.inQuadrant(mathUtil.Vector2D[int]{quadrants[2].X + 2, quadrants[0].Y + 2}, quadrants[3]) {
			lowerright++
		}

	}
	fmt.Println(upperleft, upperright, lowerleft, lowerright)
	return upperleft * upperright * lowerright * lowerleft
}

func saveGridAsImageWithText(grid string, filename string, cellSize int) error {
	// Parse the grid string into rows
	rows := strings.Split(strings.TrimSpace(grid), "\n")

	height := len(rows)
	if height == 0 {
		return nil // Nothing to render
	}
	width := len(rows[0])

	// Create a white image canvas
	img := image.NewRGBA(image.Rect(0, 0, width*cellSize, height*cellSize))
	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{0, 0, 0, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{black}, image.Point{}, draw.Src)

	// Draw each character onto the image

	for y, row := range rows {
		for x, char := range row {
			addLabel(img, x*cellSize, y*cellSize, string(char), white)
		}
	}

	// Save the image to a file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}

func addLabel(img *image.RGBA, x, y int, label string, col color.Color) {
	point := fixed.Point26_6{
		X: fixed.I(x),
		Y: fixed.I(y + 15), // Offset to position the text inside the cell
	}
	d := &font.Drawer{
		Dst:  img,
		Src:  &image.Uniform{col},
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

func printPictures(robots []robot) {
	input := readFile("./Day14/Day14.txt")
	bounds := mathUtil.Vector2D[int]{X: 101, Y: 103} // 101 103
	for i := 1; i < 9998; i++ {
		robots = moveAllRobots(input, 1, bounds)
		str := toString(robots, bounds)
		if i < 5000 {
			continue
		}
		path := "./Day14/pics/" + strconv.Itoa(i) + ".png"
		saveGridAsImageWithText(str, path, 10)
	}

}

// 6587
func SolutionDay14() {
	input := readFile("./Day14/Day14.txt")
	bounds := mathUtil.Vector2D[int]{X: 101, Y: 103} // 101 103
	robots := moveAllRobots(input, 100, bounds)
	amount := getRobotsInQuadarants(getQuadrants(bounds), robots)
	fmt.Printf("Solution Day14 Part 1: %d\n", amount)
	//for solution of part2 use two inputs to guess a range oof ca 10k numbers and print them
	//printPictures(robots)
}
