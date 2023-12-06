package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	test11 = "1abc2"
	test12 = "pqr3stu8vwx"
	test13 = "a1b2c3d4e5f"
	test14 = "treb7uchet"

	test21 = "two1nine"
	test22 = "eightwothree"
	test23 = "abcone2threexyz"
	test24 = "xtwone3four"
	test25 = "4nineeightseven2"
	test26 = "zoneight234"
	test27 = "7pqrstsixteen"

	testFile = "input.txt"
)



func getFirstLastDigits(inputText string) (int, int) {
	var firstDigit int
	var lastDigit int
	var firstDigitFound = false

	for _, character := range inputText {
		if digit, err := strconv.Atoi(string(character)); err == nil {
			if !firstDigitFound {
				firstDigitFound = true
				firstDigit = digit
			}
			lastDigit = digit
		}
	}

	return firstDigit, lastDigit
}

func getFirstLastDigitsWithText(inputText string) (int, int) {
	var firstDigit int
	var lastDigit int
	var firstDigitFound = false

	inputText = replaceTextNumbers(inputText)

	for _, character := range inputText {
		if digit, err := strconv.Atoi(string(character)); err == nil {
			if !firstDigitFound {
				firstDigitFound = true
				firstDigit = digit
			}
			lastDigit = digit
		}
	}

	return firstDigit, lastDigit
}

func replaceTextNumbers(inputText string) string {
	var lowestIdx int = -1
	var elemIdx int = -1
	var lowestIdxNum = ""
	var foundNum bool = false
	numTextArr := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for {
		for index, element := range numTextArr { 
			if strings.Contains(inputText, element) {
				elemIdx = strings.Index(inputText, element)
				if lowestIdx == -1 {
					foundNum = true
					lowestIdx = elemIdx
					lowestIdxNum = strconv.Itoa(index)
				} else if lowestIdx > elemIdx {
					lowestIdx = elemIdx
					lowestIdxNum = strconv.Itoa(index)
				}
			}
		}

		if foundNum {
			inputText = inputText[:lowestIdx] + lowestIdxNum + inputText[lowestIdx+1:]
			foundNum = false
			lowestIdx = -1
		} else {
			break
		}
	}

	return inputText
}

func createCalNum(firstDigit int, lastDigit int) int {
	return firstDigit*10 + lastDigit
}

func runPart1Test(textInput string) {
	firstDigit, lastDigit := getFirstLastDigits(textInput)
	calNum := createCalNum(firstDigit, lastDigit)
	fmt.Printf("%s, %d\n", textInput, calNum)
}

func runPart1() {
	content, err := ioutil.ReadFile(testFile)
	if err != nil {
		log.Fatal(err)
	}

	inputText := string(content)

	var runningTotal int = 0

	scanner := bufio.NewScanner(strings.NewReader(inputText))
	for scanner.Scan() {
		textLine := scanner.Text()
		firstDigit, lastDigit := getFirstLastDigits(textLine)
		calNum := createCalNum(firstDigit, lastDigit)

		runningTotal += calNum
	}
	fmt.Printf("Part 1: %d\n", runningTotal)
}

func runPart2Test(textInput string) {
	firstDigit, lastDigit := getFirstLastDigitsWithText(textInput)
	calNum := createCalNum(firstDigit, lastDigit)
	fmt.Printf("%s, %d\n", textInput, calNum)
}

func runPart2() {
	content, err := ioutil.ReadFile(testFile)
	if err != nil {
		log.Fatal(err)
	}

	inputText := string(content)

	var runningTotal int = 0

	scanner := bufio.NewScanner(strings.NewReader(inputText))
	for scanner.Scan() {
		textLine := scanner.Text()
		firstDigit, lastDigit := getFirstLastDigitsWithText(textLine)
		calNum := createCalNum(firstDigit, lastDigit)

		fmt.Printf("Text Input: %s - Calibration Number: %d\n", textLine, calNum)

		runningTotal += calNum
	}
	fmt.Printf("Part 2: %d\n", runningTotal)
}

func main() {
	fmt.Println("Part 1 Testing")
	runPart1Test(test11)
	runPart1Test(test12)
	runPart1Test(test13)
	runPart1Test(test14)

	fmt.Println()
	runPart1()
	fmt.Println()

	fmt.Println("Part 2 Testing")
	runPart2Test(test21)
	runPart2Test(test22)
	runPart2Test(test23)
	runPart2Test(test24)
	runPart2Test(test25)
	runPart2Test(test26)
	runPart2Test(test27)
	runPart2Test("oneight")

	fmt.Println()
	runPart2()
}
