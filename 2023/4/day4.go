package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func PartA() {

	file, err := os.Open("2023/4/day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {

		parts := strings.Fields(scanner.Text())

		winCount := 0
		isNums := false
		winners := make([]string, 10)

		for _, part := range parts[2:] { //first two elements will be card and number (e.g. Card 2:)

			// Should probably just split on | and loop through two separate collections
			if part == "|" {
				isNums = true
				continue
			}

			// Winning numbers
			if !isNums {
				winners = append(winners, part)
			} else {

				if slices.Contains(winners, part) {
					if winCount == 0 {
						winCount = 1
					} else {
						winCount *= 2
					}
				}

			}

		}

		total += winCount

	}

	fmt.Printf("%d\n", total)

}

func PartB() {

	file, err := os.Open("2023/4/day4b.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cardScores := make([]int, 0, 200)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		cardScores = append(cardScores, calculateWins(parts[2:]))
	}

	total := expandCards(0, len(cardScores), cardScores)

	fmt.Println(total)

}

//TODO: Could parallelise ?

func expandCards(start int, count int, allCards []int) int {

	finish := start + count

	proc := 0

	for idx, cardWins := range allCards[start:finish] {

		if cardWins > 0 {
			proc += expandCards(start+idx+1, cardWins, allCards)
		}
	}

	return proc + count

}

func calculateWins(card []string) int {

	processWinners := true
	winners := make([]string, 0, 10)
	winCount := 0

	for _, part := range card {

		// TODO Should benchmark, but it might be more obvious to split on | and loop through two separate collections
		if part == "|" {
			processWinners = false
			continue
		}

		if processWinners {
			winners = append(winners, part)
		} else {
			if slices.Contains(winners, part) {
				winCount += 1
			}
		}

	}

	return winCount

}
