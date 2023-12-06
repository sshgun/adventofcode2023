package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scn := bufio.NewScanner(f)

	schematic := make([]string, 0, 140)

	for scn.Scan() {
		schematic = append(schematic, scn.Text())
	}

	sum := 0
	gearsRatio := 0

	for i := 0; i < len(schematic); i++ {
		lineLen := len(schematic[i])

		for j := 0; j < len(schematic[i]); j++ {

			if string(schematic[i][j]) == "." {
				continue
			}

			if unicode.IsNumber(rune(schematic[i][j])) {
				jStart := j - 1
				number := []byte{schematic[i][j]}
				for ; j < lineLen-1 &&
					unicode.IsNumber(rune(schematic[i][j+1])); j++ {
					number = append(number, schematic[i][j+1])
				}

				mustAdd := false
				rnumber, _ := strconv.Atoi(string(number))

				if jStart >= 0 && isChar(schematic[i][jStart]) {
					mustAdd = true

				} else if j < lineLen-1 && isChar(schematic[i][j+1]) {
					mustAdd = true

				}

				if !mustAdd && i-1 >= 0 {
					for x := jStart; x < lineLen && x < j+2; x++ {
						if x >= 0 && isChar(schematic[i-1][x]) {
							mustAdd = true
							break
						}
					}
				}

				if !mustAdd && i+1 < len(schematic) {
					for x := jStart; x < lineLen && x < j+2; x++ {
						if x >= 0 && isChar(schematic[i+1][x]) {
							mustAdd = true
							break
						}
					}
				}

				if mustAdd {
					sum += rnumber
				}
			}

			if string(schematic[i][j]) == "*" {
				ratios := make([]int, 0)

				if j+1 < lineLen && unicode.IsNumber(rune(schematic[i][j+1])) {
					n, _ := getCompleteNumber(schematic[i], j+1)
					ratios = append(ratios, n)
				}

				if j-1 >= 0 && unicode.IsNumber(rune(schematic[i][j-1])) {
					n, _ := getCompleteNumber(schematic[i], j-1)
					ratios = append(ratios, n)
				}

				if i-1 >= 0 {
					for x := j - 1; x < j+2 && x < lineLen; x++ {
						if x >= 0 && unicode.IsNumber(rune(schematic[i-1][x])) {
							n, rightlength := getCompleteNumber(schematic[i-1], x)
							ratios = append(ratios, n)
							x += rightlength
						}
					}
				}

				if i+1 < len(schematic) {
					for x := j - 1; x < j+2 && x < lineLen; x++ {
						if x >= 0 && unicode.IsNumber(rune(schematic[i+1][x])) {
							n, rightlength := getCompleteNumber(schematic[i+1], x)
							ratios = append(ratios, n)
							x += rightlength
						}
					}
				}

				if len(ratios) == 2 {
					gr := ratios[0] * ratios[1]
					gearsRatio += gr
				}
			}
		}
	}
	fmt.Println("sum:", sum)
	fmt.Println("gearsRation:", gearsRatio)
}

func getCompleteNumber(line string, pos int) (int, int) {

	numbers := []byte{line[pos]}
	for i := pos - 1; i >= 0; i-- {
		if unicode.IsNumber(rune(line[i])) {
			numbers = append([]byte{line[i]}, numbers...)
		} else {
			break
		}
	}

	rightLength := 0
	for i := pos + 1; i < len(line); i++ {
		if unicode.IsNumber(rune(line[i])) {
			numbers = append(numbers, line[i])
			rightLength += 1
		} else {
			break
		}
	}

	n, _ := strconv.Atoi(string(numbers))
	return n, rightLength
}

func isChar(x byte) bool {

	return !unicode.IsNumber(rune(x)) && !(string(x) == ".")
}
