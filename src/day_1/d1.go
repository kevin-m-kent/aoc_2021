package main 

import (
	"bufio"
    "fmt"
	"os"
	"log"
	"strconv"
)

var num_increases int
var prev int 
var curr int 
var num_array []int

func convToInt(num_ln string) int {

	curr, err := strconv.Atoi(num_ln)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	return curr

}

// https://stackoverflow.com/questions/5884154/read-text-file-into-string-array-and-write

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

	f, err := os.Open("Data/01_p1.txt")

	if err != nil {
        log.Fatal(err)
    }

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for i:=0; scanner.Scan(); i++ {

		if i > 0 {

			raw := scanner.Text()

			curr := convToInt(raw)

			fmt.Println(curr, prev)

			if curr > prev {

				num_increases++ 

			}

			prev = curr 

		} else {

			curr = prev 	

		}

	}
	
	fmt.Printf("Part 1: There were %v increases \n", num_increases)

	//part 2 

	num_array, err = readLines("Data/01_p1.txt")

	num_increases = 0

	for i:=0; i < (len(num_array) - 3); i++ {

		is_bigger := (num_array[i] + num_array[i + 1] + num_array[i + 2]) < (num_array[i + 1] + num_array[i + 2] + num_array[i + 3])

		if is_bigger {

			num_increases++
		}


	}

	fmt.Printf("Part 2: There were %v increases \n", num_increases)

	
}