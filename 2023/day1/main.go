package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("2023/day1/input.txt")
	if err != nil {
		fmt.Printf("Failed to open main.txt %v", err)
	}
	defer input.Close()

	sc := bufio.NewScanner(input)

	total := 0

	for sc.Scan() {
		text := sc.Text()
		var firstNumRes int
		var lastNumRes int
		fIndex, firstNum := getFirstNumber(text)
		fNWIndex, firstNumWord := getFirstNumberAsWord(text)
		fmt.Printf("firstNum: %d, firstNumWord: %d\n", firstNum, firstNumWord)
		if fIndex < fNWIndex && fIndex != -1 {
			firstNumRes = firstNum
		} else {
			firstNumRes = firstNumWord
		}
		lIndex, lastNum := getLastNumber(text)
		lNWIndex, lastNumWord := getLastNumberAsWord(text)
		fmt.Printf("lastNum: %d, lastNumWord: %d\n", firstNum, firstNumWord)
		if lIndex > lNWIndex && lIndex != -1 {
			lastNumRes = lastNum
		} else {
			lastNumRes = lastNumWord
		}

		fmt.Printf("%d, %d\n", firstNumRes, lastNumRes)
		total += (firstNumRes * 10) + lastNumRes
	}

	fmt.Println(total)
}

//one, two, three, four, five, six, seven, eight, and nine

//abcone2threexyz
//run conatins for each possible number string, if it does store name in map with starting index
//at end see which one last lowest starting index

func getFirstNumberAsWord(text string) (int, int) {
	wordMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	resultMap := map[string]int{}

	for word := range wordMap {
		//get word and index if exits
		index := strings.Index(text, word)
		if index == -1 {
			//does not exist
			continue
		}
		resultMap[word] = index
	}
	//now we have a result map of all occurrences of a word
	//we now want the first occurrence so smalelst index
	fmt.Printf("Result: %v\n", resultMap)
	minKey := ""
	minValue := 10000
	for key, value := range resultMap {
		fmt.Printf("Testing key: %v, value %v\n", key, value)
		if value < minValue {
			fmt.Printf("Found value: %v is less than %v\n", value, minValue)
			minKey = key
			minValue = value
		}
	}
	fmt.Printf("returning minValue %v, and wordMap %v\n", minValue, wordMap[minKey])
	return minValue, wordMap[minKey]
}

func getLastNumberAsWord(text string) (int, int) {
	wordMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	resultMap := map[string]int{}
	for word := range wordMap {
		//get word and index if exits
		index := strings.LastIndex(text, word)
		if index == -1 {
			//does not exist
			continue
		}
		resultMap[word] = index
	}
	//now we have a result map of all occurrences of a word
	//we now want the first occurrence so smalelst index
	maxKey := ""
	maxValue := -1
	fmt.Printf("Result: %v\n", resultMap)
	for key, value := range resultMap {
		if value > maxValue {
			maxKey = key
			maxValue = value
		}
	}
	fmt.Printf("returning maxValue %v, and wordMap %v\n", maxValue, wordMap[maxKey])
	return maxValue, wordMap[maxKey]
}

func getFirstNumber(s string) (int, int) {
	for i, char := range s {
		num, err := strconv.Atoi(string(char))
		if err == nil {
			return i, num
		}
	}
	return -1, -1
}

func getLastNumber(s string) (int, int) {
	for i := len(s) - 1; i >= 0; i-- {
		num, err := strconv.Atoi(string(s[i]))
		if err == nil {
			return i, num
		}
	}
	return -1, -1
}
