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

			CheckAdjacent:
				for k := -1; k <= 1; k++ {
					for l := -1; l <= 1; l++ {
						if i+k >= 0 && i+k < len(schematic) && j+l >= 0 && j+l < len(schematic[i]) {
							if schematic[i+k][j+l] == '*' {
								gear := Coordinate{i + k, j + l}
								for m := range gearLocations {
									if gearLocations[m] == gear {
										break CheckAdjacent
									}
								}
								gearLocations = append(gearLocations, gear)
								break CheckAdjacent
							}
						}
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
