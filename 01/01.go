package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var test_1_1 = "1abc2"
var test_1_2 = "pqr3stu8vwx"
var test_1_3 = "a1b2c3d4e5f"
var test_1_4 = "treb7uchet"

var test_2_1 = "two1nine"
var test_2_2 = "eightwothree"
var test_2_3 = "abcone2threexyz"
var test_2_4 = "xtwone3four"
var test_2_5 = "4nineeightseven2"
var test_2_6 = "zoneight234"
var test_2_7 = "7pqrstsixteen"

var test_file string = "input.txt"

func get_first_last_digits(input_text string) (int, int) {
	var first_digit int
	var last_digit int
	var first_digit_found = false

	for _, character := range input_text {
		if digit, err := strconv.Atoi(string(character)); err == nil {
			if !first_digit_found {
				first_digit_found = true
				first_digit = digit
			}
			last_digit = digit
		}
	}

	return first_digit, last_digit
}

func get_first_last_digits_with_text(input_text string) (int, int) {
	var first_digit int
	var last_digit int
	var first_digit_found = false

	input_text = replace_text_numbers(input_text)

	for _, character := range input_text {
		if digit, err := strconv.Atoi(string(character)); err == nil {
			if !first_digit_found {
				first_digit_found = true
				first_digit = digit
			}
			last_digit = digit
		}
	}

	return first_digit, last_digit
}

func replace_text_numbers(input_text string) string {
	var lowest_idx int = -1
	var elem_idx int = -1
	var lowest_idx_num = ""
	var lowest_idx_text = ""
	var found_num bool = false

	num_text_arr := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for {
		for index, element := range num_text_arr { 
			if strings.Contains(input_text, element) {
				elem_idx = strings.Index(input_text, element)
				if lowest_idx == -1 {
					found_num = true
					lowest_idx = elem_idx
					lowest_idx_num = strconv.Itoa(index)
					lowest_idx_text = element
				} else if lowest_idx > elem_idx {
					lowest_idx = elem_idx
					lowest_idx_num = strconv.Itoa(index)
					lowest_idx_text = element
				}
			}
		}

		if found_num {
			input_text = strings.Replace(input_text, lowest_idx_text, lowest_idx_num, 1)
			found_num = false
			lowest_idx = -1
		} else {
			break
		}
	}

	return input_text
}

func create_calibration_number(first_digit int, last_digit int) int {
	return first_digit*10 + last_digit
}

func run_part_1_test(text_input string) {
	first_digit, last_digit := get_first_last_digits(text_input)
	calibration_number := create_calibration_number(first_digit, last_digit)
	fmt.Printf("%s, %d\n", text_input, calibration_number)
}

func run_part_1() {
	content, err := ioutil.ReadFile(test_file)
	if err != nil {
		log.Fatal(err)
	}

	input_text := string(content)

	var running_total int = 0

	scanner := bufio.NewScanner(strings.NewReader(input_text))
	for scanner.Scan() {
		textLine := scanner.Text()
		first_digit, last_digit := get_first_last_digits(textLine)
		calibration_number := create_calibration_number(first_digit, last_digit)

		running_total += calibration_number
	}
	fmt.Printf("Part 1: %d\n", running_total)
}

func run_part_2_test(text_input string) {
	first_digit, last_digit := get_first_last_digits_with_text(text_input)
	calibration_number := create_calibration_number(first_digit, last_digit)
	fmt.Printf("%s, %d\n", text_input, calibration_number)
}

func run_part_2() {
	content, err := ioutil.ReadFile(test_file)
	if err != nil {
		log.Fatal(err)
	}

	input_text := string(content)

	var running_total int = 0

	scanner := bufio.NewScanner(strings.NewReader(input_text))
	for scanner.Scan() {
		textLine := scanner.Text()
		first_digit, last_digit := get_first_last_digits_with_text(textLine)
		calibration_number := create_calibration_number(first_digit, last_digit)

		fmt.Printf("Text Input: %s - Calibration Number: %d\n", textLine, calibration_number)

		running_total += calibration_number
	}
	fmt.Printf("Part 2: %d\n", running_total)
}

func main() {
	fmt.Println("Part 1 Testing")
	run_part_1_test(test_1_1)
	run_part_1_test(test_1_2)
	run_part_1_test(test_1_3)
	run_part_1_test(test_1_4)

	fmt.Println()
	run_part_1()
	fmt.Println()

	fmt.Println("Part 2 Testing")
	run_part_2_test(test_2_1)
	run_part_2_test(test_2_2)
	run_part_2_test(test_2_3)
	run_part_2_test(test_2_4)
	run_part_2_test(test_2_5)
	run_part_2_test(test_2_6)
	run_part_2_test(test_2_7)

	fmt.Println()
	run_part_2()

	
}
