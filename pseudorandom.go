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
	fmt.Printf("{%d %v %s}\n", c.myScore, c.myCurrentNumber, c.behavior)
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
		if len(c.numbersIHaveChosenInThePast) == 0 {
			c.myCurrentNumber = generateRandomNumberFromOneToNInclusive(10)
		} else {
			leastChosenNumbers := findLeastCommonNumbers(numbersChosenLastRound)
			c.myCurrentNumber = chooseRandomNumberFromSlice(leastChosenNumbers)
		}
	case "realEstateAgent":
		if len(c.numbersIHaveChosenInThePast) == 0 {
			c.myCurrentNumber = generateRandomNumberFromOneToNInclusive(10)
		} else {
			mostChosenNumbers := findMostCommonNumbers(numbersChosenLastRound)
			c.myCurrentNumber = chooseRandomNumberFromSlice(mostChosenNumbers)
		}
	case "completelyRandom":
		c.myCurrentNumber = generateRandomNumberFromOneToNInclusive(10)
	}
	c.numbersIHaveChosenInThePast = append(c.numbersIHaveChosenInThePast, c.myCurrentNumber) // add the number chosen to the history
}

func simulateOneRound(allContestants []*Contestant, numbersChosenLastRound []int) {
	for i := 0; i < len(allContestants); i++ {
		contestant := allContestants[i]
		contestant.chooseNumber(numbersChosenLastRound)
	}

	numbersChosenLastRound = []int{}
	for _, contestant := range allContestants {
		numbersChosenLastRound = append(numbersChosenLastRound, contestant.myCurrentNumber)
	}

	numbersMap := make(map[int]int)
	for _, num := range numbersChosenLastRound {
		numbersMap[num]++
	}

	for _, contestant := range allContestants {
		if numbersMap[contestant.myCurrentNumber] == 1 {
			contestant.myScore++
		}
		contestant.printInfo()
	}
	fmt.Println("round:", numbersChosenLastRound)
	fmt.Println(" ")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	numbersChosenLastRound := []int{0, 0, 0, 0}
	allContestants := []*Contestant{
		{1, 0, 0, []int{}, []int{}, "rock"},
		{2, 0, 0, []int{}, []int{}, "realEstateAgent"},
		{3, 0, 0, []int{}, []int{}, "realEstateAgent"},
		{4, 0, 0, []int{}, []int{}, "opportunist"},
	}
	for i := 1; i <= 1000; i++ {
		simulateOneRound(allContestants, numbersChosenLastRound)
	}
}
