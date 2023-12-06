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
	testFile = "input.txt"
	testInput = "test_input.txt"
)

type Matrix [][]int

func getText(fileName string) string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func getNumCharacters(inputText string) int {
	var numChar int = 0

	scanner := bufio.NewScanner(strings.NewReader(inputText))
	for scanner.Scan() {
		textLine := scanner.Text()
		numChar = len(textLine)
		break
	}

	return numChar
}

func createEmptyMatrix(val int) Matrix {
	// square matrix
	matrix := make([][]int, val)
	for i := 0; i < val; i++ {
		matrix[i] = make([]int, val)
	}

	return matrix
}

func fillMatrix(matrix Matrix, inputText string) Matrix {
	y := 0

	var char string
	var num int

	scanner := bufio.NewScanner(strings.NewReader(inputText))

	for scanner.Scan() {
		textLine := scanner.Text()

		for i := 0; i < len(matrix); i++ {
			char = string(textLine[i])

			if char == "." {
				num = -2
			} else if val, err := strconv.Atoi(char); err == nil {
				num = val
			} else if char == "*" {
				num = -3
			}else {
				num = -1
			}

			matrix[y][i] = num
		}

		y += 1
	}

	return matrix
}

func combineNums(matrix Matrix) Matrix {
	var curVal int = 0
	var runVal int = 0

	matrixLen := len(matrix)

	for y := 0; y < matrixLen; y++ {
		for x := 0; x < matrixLen; x++ {
			curVal = matrix[y][x]

			if curVal >= 0 {
				runVal = runVal*10 + curVal
			} else if runVal != 0 {
				for i := 0; i < len(strconv.Itoa(runVal)); i++ {
					matrix[y][x-i-1] = runVal
				}
				runVal = 0
			}

		}
		if runVal != 0 {
			for i := 0; i < len(strconv.Itoa(runVal)); i++ {
				matrix[y][matrixLen-i-1] = runVal
			}
			runVal = 0
		}
	}
	return matrix
}

func removeNum(matrix Matrix, y int, x int) Matrix {
	curVal := matrix[y][x]

	for i := x - 1; i > -1; i-- {
		if matrix[y][i] == curVal {
			matrix[y][i] = -2
		} else {
			break
		}
	}
	for i := x + 1; i < len(matrix); i++ {
		if matrix[y][i] == curVal {
			matrix[y][i] = -2
		} else {
			break
		}
	}

	matrix[y][x] = -2

	return matrix
}

func getLocalSum(matrix Matrix, y int, x int) (Matrix, int) {
	var localVal int = 0
	var localSum int = 0

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			localVal = matrix[y+i][x+j]

			if localVal >= 0 {
				localSum += localVal
				matrix = removeNum(matrix, y+i, x+j)
			}
		}
	}

	return matrix, localSum
}

func getNumSum(matrix Matrix) (Matrix, int) {
	var curVal int = 0
	var runSum int = 0
	var localSum int = 0

	matrixLen := len(matrix)

	for y := 0; y < matrixLen; y++ {
		for x := 0; x < matrixLen; x++ {
			curVal = matrix[y][x]

			if curVal == -1 || curVal == -3 {
				matrix, localSum = getLocalSum(matrix, y, x)
				runSum += localSum
			}
		}
	}

	return matrix, runSum
}

func getLocalGearRatio(matrix Matrix, y int, x int) (Matrix, int) {
	var gearCount int = 0
	var gearRatio int = 1
	var localVal int = 0

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			localVal = matrix[y+i][x+j]

			if localVal >= 0 {
				gearCount += 1
				gearRatio *= localVal
				matrix = removeNum(matrix, y+i, x+j)
			}
		}
	}

	if gearCount != 2 {
		gearRatio = 0
	}

	return matrix, gearRatio
}

func getGearRatioSum(matrix Matrix) (Matrix, int) {
	var curVal int = 0
	var ratioSum int = 0
	var gearRatio int = 0

	matrixLen := len(matrix)

	for y := 0; y < matrixLen; y++ {
		for x := 0; x < matrixLen; x++ {
			curVal = matrix[y][x]

			if curVal == -3 {
				matrix, gearRatio = getLocalGearRatio(matrix, y, x)
				ratioSum += gearRatio
			}
		}
	}

	return matrix, ratioSum
}

func runPart1(matrix Matrix) {
	var numSum int

	_, numSum = getNumSum(matrix)

	fmt.Printf("Part 1 Solution: %d\n", numSum)
}

func runPart2(matrix Matrix) {
	var gearRatioSum int

	_, gearRatioSum = getGearRatioSum(matrix)

	fmt.Printf("Part 2 Solution: %d\n", gearRatioSum)
}

func main() {
	var inputText string = getText(testFile)
	var numChar int = getNumCharacters(inputText)
	var matrix Matrix = createEmptyMatrix(numChar)

	matrix = fillMatrix(matrix, inputText)
	matrix = combineNums(matrix)

	runPart1(matrix)

	matrix = fillMatrix(matrix, inputText)
	matrix = combineNums(matrix)

	runPart2(matrix)
}
