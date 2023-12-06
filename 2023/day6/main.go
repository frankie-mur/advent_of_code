package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := GetInputDataFromFile("2023/day6/input.txt")
	time := strings.Split(data[0], ":")[1]
	distance := strings.Split(data[1], ":")[1]
	timeValues := getValues(time)
	distanceValues := getValues(distance)

	//Start from 0 to time incrementing
	//for each number calculate (number of seconds total -  the current number) * current number
	// if that number is greater than distance count that number towards a way to win
	//at the end of "game" calculate multiply times to win with total

	part1res := part1(timeValues, distanceValues)
	part2res := part2(timeValues, distanceValues)

	fmt.Printf("part1: %d, part2: %d", part1res, part2res)
}

func part1(timeValues []int, distanceValues []int) int {
	total := 1
	for index := 0; index < len(timeValues); index++ {
		canWin := 0
		currentTime := timeValues[index]
		for j := 0; j <= currentTime; j++ {
			calcScore := (currentTime - j) * j
			if calcScore > distanceValues[index] {
				canWin += 1
			}
		}
		total *= canWin
	}
	return total
}

func part2(timeValues []int, distanceValues []int) int {
	timeValue := arrayToSingleInt(timeValues)
	distanceValue := arrayToSingleInt(distanceValues)
	total := 0
	for i := 1; i < timeValue; i++ {
		if (timeValue-i)*i > distanceValue {
			total++
		}
	}
	return total
}

func getValues(data string) []int {
	values := []int{}
	fields := strings.Fields(data)

	for _, v := range fields {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Errorf("failed to convert string to int: %v", err)
		}
		values = append(values, num)
	}
	return values
}

func arrayToSingleInt(array []int) int {

	// Convert time values to strings
	var timeStrings []string
	for _, value := range array {
		timeStrings = append(timeStrings, strconv.Itoa(value))
	}

	// Concatenate time values into a single integer
	combinedTime, err := strconv.Atoi(strings.Join(timeStrings, ""))
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return 0
	}

	return combinedTime
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
