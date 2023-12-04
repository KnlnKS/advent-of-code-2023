package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the input file
	file, err := os.Open("./input.txt")
	if err != nil {
		return
	}
	defer file.Close()

	// Create a buffered reader
	reader := bufio.NewReader(file)

	// Total number of cards
	total := 0

	var cards [220]int

	// iterate over each card
	for cardNum := 0; ; cardNum++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		// Add the default card to the array
		cards[cardNum]++
		// Add all the cards to the total
		total += cards[cardNum]

		// Skip the `Card   1: ` part
		var i int
		for i = 3; line[i-2] != ':'; i++ {
		}

		tempNum := 0

		// Read the winning numbers
		var winningNums [10]int
		for c := 0; line[i] != '|'; i++ {
			switch line[i] {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				tempNum = (tempNum * 10) + int(line[i]-'0')
			default:
				if tempNum != 0 {
					winningNums[c] = tempNum
					tempNum = 0
					c++
				}
			}
		}

		// Skip to first number
		i += 2

		// Card point total
		cardTotal := 0

		// Read the card numbers
		for ; i < len(line); i++ {
			switch line[i] {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				tempNum = (tempNum * 10) + int(line[i]-'0')
			default:
				if tempNum != 0 {
					for j := range winningNums {
						if tempNum == winningNums[j] {
							cardTotal++
							break
						}
					}
					tempNum = 0
				}
			}
		}

		for j := 1; j <= cardTotal; j++ {
			cards[cardNum+j] += cards[cardNum]
		}
	}

	// Print the total
	fmt.Println(total)
}
