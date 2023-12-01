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

	// Keep track of the total
	var total int

	// Create a buffered reader
	reader := bufio.NewReader(file)
	for {
		// Read line-by-line
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		// Find the first and last digits in the line
		num1, num2 := -1, -1
		for i, c := range line {
			switch c {
			// Check if the character is a digit
			case 49, 50, 51, 52, 53, 54, 55, 56, 57:
				if num1 < 0 {
					num1 = int(c - 48)
					num2 = num1
				} else {
					num2 = int(c - 48)
				}
			// Check if one
			case 'o':
				if i < len(line)-2 && line[i+1] == 'n' && line[i+2] == 'e' {
					if num1 < 0 {
						num1 = 1
						num2 = num1
					} else {
						num2 = 1
					}
				}
			// Check if two or three
			case 't':
				if i < len(line)-2 && line[i+1] == 'w' && line[i+2] == 'o' {
					if num1 < 0 {
						num1 = 2
						num2 = num1
					} else {
						num2 = 2
					}
				} else if i < len(line)-4 && line[i+1] == 'h' && line[i+2] == 'r' && line[i+3] == 'e' && line[i+4] == 'e' {
					if num1 < 0 {
						num1 = 3
						num2 = num1
					} else {
						num2 = 3
					}
				}
			// Check if four or five
			case 'f':
				if i < len(line)-3 && line[i+1] == 'o' && line[i+2] == 'u' && line[i+3] == 'r' {
					if num1 < 0 {
						num1 = 4
						num2 = num1
					} else {
						num2 = 4
					}
				} else if i < len(line)-3 && line[i+1] == 'i' && line[i+2] == 'v' && line[i+3] == 'e' {
					if num1 < 0 {
						num1 = 5
						num2 = num1
					} else {
						num2 = 5
					}
				}
			// Check if six or seven
			case 's':
				if i < len(line)-2 && line[i+1] == 'i' && line[i+2] == 'x' {
					if num1 < 0 {
						num1 = 6
						num2 = num1
					} else {
						num2 = 6
					}
				} else if i < len(line)-4 && line[i+1] == 'e' && line[i+2] == 'v' && line[i+3] == 'e' && line[i+4] == 'n' {
					if num1 < 0 {
						num1 = 7
						num2 = num1
					} else {
						num2 = 7
					}
				}
			// Check if eight
			case 'e':
				if i < len(line)-4 && line[i+1] == 'i' && line[i+2] == 'g' && line[i+3] == 'h' && line[i+4] == 't' {
					if num1 < 0 {
						num1 = 8
						num2 = num1
					} else {
						num2 = 8
					}
				}
			// Check if nine
			case 'n':
				if i < len(line)-3 && line[i+1] == 'i' && line[i+2] == 'n' && line[i+3] == 'e' {
					if num1 < 0 {
						num1 = 9
						num2 = num1
					} else {
						num2 = 9
					}
				}
			}

		}

		// Put the digits together and add to the total
		total += (num1 * 10) + num2
	}

	fmt.Println(total)
}
