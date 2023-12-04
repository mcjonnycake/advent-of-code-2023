package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

var (
	testfile = "input.txt"
	testinput = "test_input.txt"
)

const (
	maxValidRed = 12
	maxValidGreen = 13
	maxValidBlue = 14
)

func getColorCount(reveal string) (int, int, int) {
	bcount := 0
	rcount := 0
	gcount := 0

	s := strings.Split(reveal, ", ")

	for _, color := range s {
		colorS := strings.Split(color, " ")

		if colorS[1] == "blue" {
			bcount, _ = strconv.Atoi(colorS[0])
		} else if colorS[1] == "red" {
			rcount, _ = strconv.Atoi(colorS[0])
		} else if colorS[1] == "green" {
			gcount, _ = strconv.Atoi(colorS[0])
		}
	}

	return bcount, rcount, gcount
}

func isValidGame(reveals string) bool {
	var maxbcount int = 0
	var maxrcount int = 0
	var maxgcount int = 0

	s := strings.Split(reveals, "; ")
	for _, reveal := range s {
		bcount, rcount, gcount := getColorCount(reveal)

		if bcount > maxbcount {
			maxbcount = bcount
		}
		if rcount > maxrcount {
			maxrcount = rcount
		}
		if gcount > maxgcount {
			maxgcount = gcount
		}
	}

	validr := maxValidRed >= maxrcount
	validg := maxValidGreen >= maxgcount
	validb := maxValidBlue >= maxbcount

	if validr && validg && validb {
		return true
	} else {
		return false
	}
}

func getGamePower(reveals string) int {
	var maxbcount int
	var maxrcount int
	var maxgcount int

	s := strings.Split(reveals, "; ")
	for _, reveal := range s {
		bcount, rcount, gcount := getColorCount(reveal)

		if bcount > maxbcount {
			maxbcount = bcount
		}
		if rcount > maxrcount {
			maxrcount = rcount
		}
		if gcount > maxgcount {
			maxgcount = gcount
		}
	}

	return maxbcount * maxrcount * maxgcount
}

func parseValidGame(input_text string) (bool, int) {
	gameSplit := strings.Split(input_text, ": ")
	gameIDSplit := strings.Split(gameSplit[0], " ")
	gameID, _ := strconv.Atoi(gameIDSplit[1])

	if isValidGame(gameSplit[1]) {
		return true, gameID
	} else {
		return false, gameID
	}
}

func parseGamePower(input_text string) int {
	gameSplit := strings.Split(input_text, ": ")
	gamePower := getGamePower(gameSplit[1])

	return gamePower
}

func runPart1() {
	validSum := 0

	content, err := ioutil.ReadFile(testfile)
	if err != nil {
		log.Fatal(err)
	}

	inputText := string(content)

	scanner := bufio.NewScanner(strings.NewReader(inputText))
	for scanner.Scan() {
		textLine := scanner.Text()
		valid, id:= parseValidGame(textLine)
		
		if valid {
			validSum += id
		}
	}

	fmt.Printf("Total Valid ID Sum: %d\n", validSum)
}

func runPart2() {
	gamePowerSum := 0

	content, err := ioutil.ReadFile(testfile)
	if err != nil {
		log.Fatal(err)
	}

	inputText := string(content)

	scanner := bufio.NewScanner(strings.NewReader(inputText))
	for scanner.Scan() {
		textLine := scanner.Text()

		gamePowerSum += parseGamePower(textLine)
	}

	fmt.Printf("Total Power: %d\n", gamePowerSum)
}

func main() {
	runPart1()
	runPart2()
}
