package main 

import (
	
	"fmt"
	"os"
	"strconv"
	"bufio"
	"log"
	"strings"
	"math"
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

	f, err := os.Open("Data/03.txt")

	if err != nil {
        log.Fatal(err)
    }

	defer f.Close()

	scanner := bufio.NewScanner(f)

	totals := make(map[int]int)
	total_lines := 0

	for scanner.Scan() {

		line := scanner.Text()
		line_split := strings.Split(line, "")
		line_length := len(line_split)

		for i, s:= range line_split {

			totals[line_length - i - 1] = totals[line_length - i - 1] + convToInt(s)

		}
		total_lines++



	}

	gamma := make(map[int]int) 

	epsilon := make(map[int]int) 

	var avg_totals float32

	for i, _:= range totals {

		avg_totals = float32(totals[i])/float32(total_lines) 

		if avg_totals > .50 {

			gamma[len(totals) - i - 1] = 1
			epsilon[len(totals) - i - 1] = 0

		} else {

			gamma[len(totals) - i - 1] = 0
			epsilon[len(totals) - i - 1] = 1
		}
	}

	var gamma_calc float64
	var epsilon_calc float64

	for i:=0; i < len(gamma); i++ {

		gamma_calc = gamma_calc + float64(gamma[i])*float64(math.Pow(2, float64(i)))
		epsilon_calc = epsilon_calc + float64(epsilon[i])*float64(math.Pow(2, float64(i)))

	}

	fmt.Print(gamma_calc*epsilon_calc)




}