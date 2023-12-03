package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinate struct {
	i int
	j int
}

// See if the gear location already exists in the map
func find(gearLocations []Coordinate, gear Coordinate) bool {
	for _, gearLocation := range gearLocations {
		if gearLocation == gear {
			return true
		}
	}
	return false
}

func main() {
	// Open the input file
	file, err := os.Open("./input.txt")
	if err != nil {
		return
	}
	defer file.Close()

	// Create a buffered reader
	reader := bufio.NewReader(file)

	// Load file into a 2D array
	var schematic [][]rune
	for i := 0; ; i++ {
		// Read line-by-line
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		// add line except the last character
		schematic = append(schematic, []rune(line))
	}

	// Keep track of the gear locations, and any numbers associated with them
	var gearMap = make(map[Coordinate][]int)

	// Iterate over the 2D array
	for i := range schematic {
		bufferNum := 0
		var gearLocations []Coordinate

		for j := range schematic[i] {
			switch schematic[i][j] {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				bufferNum = (bufferNum * 10) + int(schematic[i][j]-'0')

				if i > 0 && schematic[i-1][j] == '*' {
					if !find(gearLocations, Coordinate{i - 1, j}) {
						gearLocations = append(gearLocations, Coordinate{i - 1, j})
					}
				}
				if i > 0 && j > 0 && schematic[i-1][j-1] == '*' {
					if !find(gearLocations, Coordinate{i - 1, j - 1}) {
						gearLocations = append(gearLocations, Coordinate{i - 1, j - 1})
					}
				}
				if j > 0 && schematic[i][j-1] == '*' {
					if !find(gearLocations, Coordinate{i, j - 1}) {
						gearLocations = append(gearLocations, Coordinate{i, j - 1})
					}
				}
				if i > 0 && j != len(schematic[i])-1 && schematic[i-1][j+1] == '*' {
					if !find(gearLocations, Coordinate{i - 1, j + 1}) {
						gearLocations = append(gearLocations, Coordinate{i - 1, j + 1})
					}
				}
				if i < len(schematic)-1 && schematic[i+1][j] == '*' {
					if !find(gearLocations, Coordinate{i + 1, j}) {
						gearLocations = append(gearLocations, Coordinate{i + 1, j})
					}
				}
				if i < len(schematic)-1 && j > 0 && schematic[i+1][j-1] == '*' {
					if !find(gearLocations, Coordinate{i + 1, j - 1}) {
						gearLocations = append(gearLocations, Coordinate{i + 1, j - 1})
					}
				}
				if i < len(schematic)-1 && j != len(schematic[i])-1 && schematic[i+1][j+1] == '*' {
					if !find(gearLocations, Coordinate{i + 1, j + 1}) {
						gearLocations = append(gearLocations, Coordinate{i + 1, j + 1})
					}
				}
				if j < len(schematic[i])-1 && schematic[i][j+1] == '*' {
					if !find(gearLocations, Coordinate{i, j + 1}) {
						gearLocations = append(gearLocations, Coordinate{i, j + 1})
					}
				}

			default:
				if len(gearLocations) > 0 {
					for _, gearLocation := range gearLocations {
						gearMap[gearLocation] = append(gearMap[gearLocation], bufferNum)
					}
				}
				bufferNum = 0
				gearLocations = nil
			}
		}
	}

	// Calculate the total
	total := 0
	for _, v := range gearMap {
		if len(v) == 2 {
			total += v[0] * v[1]
		}
	}

	fmt.Println(total)
}
