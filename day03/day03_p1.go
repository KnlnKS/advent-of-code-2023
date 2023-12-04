package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Open the input file
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
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

	// Keep track of the total
	var total int

	// Iterate over the 2D array
	for i := range schematic {
		bufferNum := 0
		isBufferNumValid := false

		for j := range schematic[i] {
			switch schematic[i][j] {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				bufferNum = (bufferNum * 10) + int(schematic[i][j]-'0')
				if !isBufferNumValid {
				CheckAdjacent:
					for k := -1; k <= 1; k++ {
						for l := -1; l <= 1; l++ {
							if i+k >= 0 && i+k < len(schematic) && j+l >= 0 && j+l < len(schematic[i]) {
								if schematic[i+k][j+l] != '.' && schematic[i+k][j+l] != '\n' && (schematic[i+k][j+l] < '0' || schematic[i+k][j+l] > '9') {
									isBufferNumValid = true
									break CheckAdjacent
								}
							}
						}
					}
				}
			default:
				if isBufferNumValid {
					total += bufferNum
				}
				bufferNum = 0
				isBufferNumValid = false
			}
		}
	}

	fmt.Println(total)

}
