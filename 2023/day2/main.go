package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("2023/day2/input.txt")
	if err != nil {
		fmt.Printf("Failed to open main.txt %v", err)
	}
	defer input.Close()

	sc := bufio.NewScanner(input)
	score := 0

	for sc.Scan() {
		text := sc.Text()
		//Game id
		gameId := getGameId(text)
		gameSet := getGameSet(text)
		fmt.Printf("Game id: %v\n", gameId)
		//		isValid := true
		maxGreen := -math.MaxInt
		maxBlue := -math.MaxInt
		maxRed := -math.MaxInt

		for _, game := range gameSet {
			red := 0
			green := 0
			blue := 0
			trimed := strings.TrimSpace(game)
			fmt.Printf("Game set: %v\n", trimed)
			split := strings.Split(trimed, ",")
			for _, color := range split {
				//split at space
				trimEachColor := strings.TrimSpace(color)
				eachSet := strings.Split(trimEachColor, " ")
				for idx, v := range eachSet {
					switch v {
					case "green":
						num, err := strconv.Atoi(eachSet[idx-1])
						if err != nil {
							fmt.Printf("Error: %v\n", err)
						}
						green += num
					case "red":
						num, err := strconv.Atoi(eachSet[idx-1])
						if err != nil {
							fmt.Printf("Error: %v\n", err)
						}
						red += num
					case "blue":
						num, err := strconv.Atoi(eachSet[idx-1])
						if err != nil {
							fmt.Printf("Error: %v\n", err)
						}
						blue += num
					default:
						continue
					}
				}
				fmt.Printf("Red: %v, Green: %v, Blue: %v\n", red, green, blue)
			}
			// //Check if set is valid
			// err := checkScores(red, green, blue)
			// if err != nil {
			// 	isValid = false
			// 	break
			// }
			if red > maxRed {
				maxRed = red
			}
			if green > maxGreen {
				maxGreen = green
			}
			if blue > maxBlue {
				maxBlue = blue
			}
		}
		// if isValid {
		// 	score += gameId
		// } else {
		// 	fmt.Printf("Game id: %v is not valid\n", gameId)
		// }
		fmt.Printf("maxBlue: %v maxGreen: %v maxBlue: %v\n", maxBlue, maxGreen, maxBlue)
		score += (maxBlue * maxGreen * maxRed)
	}

	fmt.Printf("SCORE: %d\n", score)
}

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func getGameId(s string) int {
	split := strings.Split(s, ":")
	gameId := strings.Split(split[0], " ")[1]
	toNum, err := strconv.Atoi(gameId)
	if err != nil {
		return -1
	}
	return toNum
}

func getGameSet(s string) []string {
	split := strings.Split(s, ":")
	gameSets := strings.Split(split[1], ";")
	return gameSets
}

func checkScores(red int, green int, blue int) error {
	maxSets := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	if maxSets["red"] < red {
		return errors.New("maxed out on red")
	}

	if maxSets["green"] < green {
		return errors.New("maxed out on green")
	}
	if maxSets["blue"] < blue {
		return errors.New("maxed out on blue")
	}

	return nil

}
