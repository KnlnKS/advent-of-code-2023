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
		for _, c := range line {
			if c > 47 && c < 58 {
				if num1 < 0 {
					num1 = int(c - 48)
					num2 = num1
				} else {
					num2 = int(c - 48)
				}
			}
		}

		// Put the digits together and add to the total
		total += (num1 * 10) + num2
	}

	fmt.Println(total)
}
