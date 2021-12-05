package main 

import (
	
	"fmt"
	"os"
	"strconv"
	"bufio"
	"log"
	"strings"
)

func convToInt(num_ln string) int {

	curr, err := strconv.Atoi(num_ln)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	return curr

}

func readLines(path string) ([]int, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		text := scanner.Text()
		text_int := convToInt(text)
        lines = append(lines, text_int)
    }
    return lines, scanner.Err()
}

func main() {
	
	// part 1

	position_dict := make(map[string]int)

	position_dict["horizontal"] = 0
	position_dict["vertical"] = 0

	f, err := os.Open("Data/02.txt")

	if err != nil {
        log.Fatal(err)
    }

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		line := scanner.Text()
		line_split := strings.Split(line, " ")
		value := convToInt(line_split[1])
		direction := line_split[0]

		switch direction {

		case "forward": position_dict["horizontal"] = position_dict["horizontal"] + value
		case "up": position_dict["vertical"] = position_dict["vertical"] - value
		case "down":  position_dict["vertical"] = position_dict["vertical"] + value

		}


	}

	total_mult := position_dict["horizontal"]*position_dict["vertical"]

	fmt.Printf("Total horizontal*vertical = %v\n", total_mult)

	// part 2 

	position_dict["horizontal"] = 0
	position_dict["vertical"] = 0
	position_dict["aim"] = 0

	f2, err := os.Open("Data/02.txt")

	if err != nil {
        log.Fatal(err)
    }

	defer f2.Close()

	scanner_2 := bufio.NewScanner(f2)

	for scanner_2.Scan() {

		line := scanner_2.Text()
		line_split := strings.Split(line, " ")
		value := convToInt(line_split[1])
		direction := line_split[0]

		switch direction {

		case "forward": 
			position_dict["horizontal"] = position_dict["horizontal"] + value
			position_dict["vertical"] = position_dict["vertical"] + position_dict["aim"]*value
		case "up": position_dict["aim"] = position_dict["aim"] - value
		case "down":  position_dict["aim"] = position_dict["aim"] + value

		}


	}

	total_mult = position_dict["horizontal"]*position_dict["vertical"]

	fmt.Printf("Part 2: Total horizontal*vertical = %v\n", total_mult)



}