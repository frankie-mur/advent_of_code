package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data := GetInputDataFromFile("2023/day4/input.txt")
	//Seperate at the :
	part1Res := part1(data)
	part2Res := part2(data)

	fmt.Printf("Part1: %d\nPart2: %d\n", part1Res, part2Res)
}

func part1(data []string) int {
	total := 0
	for _, d := range data {
		points := 0
		numbers := strings.Split(d, ":")[1]
		winningAndMyNums := strings.Split(numbers, "|")
		//Make winning nums a set and iterate through my nums, if in set then a match
		winningNums, myNums := strings.Fields(winningAndMyNums[0]), strings.Fields(winningAndMyNums[1])
		winningSet := makeNumsSet(winningNums)
		for _, num := range myNums {
			toInt, _ := strconv.Atoi(num)
			if _, ok := winningSet[toInt]; ok {
				//In the set
				if points >= 1 {
					points *= 2
				} else {
					points += 1
				}
			}
		}
		total += points
	}
	return total
}

func part2(data []string) int {
	var cardWinsNoOfCards = make(map[int]int)
	var cardQueue []int

	// Calculate what card wins how many cards
	for i, card := range data {
		var currentCardWins int

		numbersPresent := strings.Split(strings.Split(card, ":")[1], " | ")
		winningNumbers, myNumbers := strings.Fields(numbersPresent[0]), strings.Fields(numbersPresent[1])
		for _, number := range myNumbers {
			if slices.Contains(winningNumbers, number) {
				currentCardWins += 1
			}
		}

		cardWinsNoOfCards[i+1] = currentCardWins
		cardQueue = append(cardQueue, i+1)
	}

	var cardsProcessed int
	var queueLength = len(cardQueue)
	for cardsProcessed < queueLength {
		generatedRange := generateRange(cardQueue[cardsProcessed], cardQueue[cardsProcessed]+cardWinsNoOfCards[cardQueue[cardsProcessed]])
		cardQueue = append(cardQueue, generatedRange...)
		cardsProcessed++
		queueLength += len(generatedRange)
	}

	return cardsProcessed
}

func generateRange(start, end int) []int {
	var numRange = make([]int, 0, end-start+1)
	for i := start + 1; i <= end; i++ {
		numRange = append(numRange, i)
	}

	return numRange
}

func makeNumsSet(data []string) map[int]bool {
	var set map[int]bool = make(map[int]bool)

	for _, num := range data {
		num, _ := strconv.Atoi(num)
		set[num] = true
	}

	return set
}

func GetInputDataFromFile(fileName string) []string {

	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	contents := strings.Trim(string(file), "\n")

	data := strings.Split(contents, "\n")
	return data
}
