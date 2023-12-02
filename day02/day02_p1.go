package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	// Open the input file
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Keep track of the total
	var total int

	// Create a buffered reader
	reader := bufio.NewReader(file)

	// Regex to find the first two digits followed by colour
	// We only need to check the first two digits because the
	// input is guaranteed to be valid for a single digit # of cubes
	re := regexp.MustCompile(`\d\d+ (r|g|b)`)

GameLoop:
	for gameID := 1; true; gameID++ {
		// Read line-by-line
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		//
		matches := re.FindAllString(line, -1)

		// If there are no matches, the game is valid
		if len(matches) > 0 {
			for i := range matches {
				// If the first digit is 2 or greater, or the second digit is greater than the mentioned amount, the game is invalid
				switch matches[i][len(matches[i])-1] {
				case 'r':
					if matches[i][0] > '1' || matches[i][1] > '2' {
						continue GameLoop
					}
				case 'g':
					if matches[i][0] > '1' || matches[i][1] > '3' {
						continue GameLoop
					}
				case 'b':
					if matches[i][0] > '1' || matches[i][1] > '4' {
						continue GameLoop
					}
				}
			}
		}

		total += gameID
	}

	fmt.Println(total)
}
