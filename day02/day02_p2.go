package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

	// Regex to find the digits followed by colour
	re := regexp.MustCompile(`\d+ (r|g|b)`)

	for {
		// Read line-by-line
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		r, g, b := 0, 0, 0

		matches := re.FindAllString(line, -1)

		for i := range matches {
			switch matches[i][len(matches[i])-1] {
			case 'r':
				if len(matches[i]) == 3 {
					if num, _ := strconv.Atoi(matches[i][0:1]); num > r {
						r = num
					}
				} else {
					if num, _ := strconv.Atoi(matches[i][0:2]); num > r {
						r = num
					}
				}
			case 'g':
				if len(matches[i]) == 3 {
					if num, _ := strconv.Atoi(matches[i][0:1]); num > g {
						g = num
					}
				} else {
					if num, _ := strconv.Atoi(matches[i][0:2]); num > g {
						g = num
					}
				}
			case 'b':
				if len(matches[i]) == 3 {
					if num, _ := strconv.Atoi(matches[i][0:1]); num > b {
						b = num
					}
				} else {
					if num, _ := strconv.Atoi(matches[i][0:2]); num > b {
						b = num
					}
				}
			}
		}

		total += r * g * b
	}

	fmt.Println(total)
}
