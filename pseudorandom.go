package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Contestant struct {
	id                          int
	myScore                     int
	myCurrentNumber             int
	numbersChosenLastRound      []int
	numbersIHaveChosenInThePast []int
	behavior                    string
}

func generateRandomNumberFromOneToNInclusive(n int) int {
	return (rand.Intn(n) + 1)
}

func chooseRandomNumberFromSlice(numbers []int) int {
	randomIndex := rand.Intn(len(numbers))
	return numbers[randomIndex]
}

func findMostCommonNumbers(nums []int) []int {
	countMap := make(map[int]int)
	mostCommon := []int{}
	maxCount := 0

	for _, num := range nums {
		countMap[num]++
		if countMap[num] > maxCount {
			maxCount = countMap[num]
		}
	}

	for num, count := range countMap {
		if count == maxCount {
			mostCommon = append(mostCommon, num)
		}
	}
	return mostCommon
}

func findLeastCommonNumbers(nums []int) []int {
	countMap := make(map[int]int)
	leastCommonNumbers := []int{}
	minCount := len(nums) + 1

	for _, num := range nums {
		countMap[num]++
	}

	for num := 1; num <= 10; num++ {
		if countMap[num] < minCount {
			minCount = countMap[num]
		}
	}

	for num := 1; num <= 10; num++ {
		if countMap[num] == minCount {
			leastCommonNumbers = append(leastCommonNumbers, num)
		}
	}
	return leastCommonNumbers
}

func (c *Contestant) printInfo() {
	fmt.Printf("{%d %d %d %v %s}\n", c.id, c.myScore, c.myCurrentNumber, c.numbersChosenLastRound, c.behavior)
}

func (c *Contestant) chooseRandomNumber(numbersChosenLastRound []int) {
	randomNumber := generateRandomNumberFromOneToNInclusive(10)
	c.myCurrentNumber = randomNumber
}

func (c *Contestant) chooseNumber(numbersChosenLastRound []int) {
	switch c.behavior {
	case "rock":
		if len(c.numbersIHaveChosenInThePast) == 0 {
			c.myCurrentNumber = generateRandomNumberFromOneToNInclusive(10)
		}
	case "opportunist":
		leastChosenNumbers := findLeastCommonNumbers(numbersChosenLastRound)
		c.myCurrentNumber = chooseRandomNumberFromSlice(leastChosenNumbers)
	case "realEstateAgent":
		mostChosenNumbers := findMostCommonNumbers(numbersChosenLastRound)
		c.myCurrentNumber = chooseRandomNumberFromSlice(mostChosenNumbers)
	case "completelyRandom":
		c.myCurrentNumber = generateRandomNumberFromOneToNInclusive(10)
	}
	c.numbersIHaveChosenInThePast = append(c.numbersIHaveChosenInThePast, c.myCurrentNumber) // add the number chosen to the history
}

func main() {
	rand.Seed(time.Now().UnixNano())
	numbersChosenLastRound := []int{0, 0, 0, 0}
	allContestants := []*Contestant{
		{1, 0, 0, []int{}, []int{}, "completelyRandom"},
		{2, 0, 0, []int{}, []int{}, "rock"},
	}
	for _, contestant := range allContestants {
		contestant.chooseRandomNumber(numbersChosenLastRound)
		contestant.printInfo()
	}

	fmt.Println(numbersChosenLastRound)
	numbersChosenLastRound = []int{}
	fmt.Println(numbersChosenLastRound)

	for _, contestant := range allContestants {
		numbersChosenLastRound = append(numbersChosenLastRound, contestant.myCurrentNumber)
	}
	fmt.Println(numbersChosenLastRound)
}