package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("u need to pass only one argument")
		return
	}
	f, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	sum := 0
	scn := bufio.NewScanner(f)
	for scn.Scan() {
		line := bytes.Runes(scn.Bytes())
		first, last := "", ""
		for i := 0; i < len(line); i++ {
			c := string(line[i])
			if _, err := strconv.Atoi(c); err != nil { // is a number ?
				wn := ""
				word := []rune{line[i]}
				for j := i + 1; j < len(line) && len(word) <= 5; j++ {
					if !unicode.IsLetter(line[j]) {
						break
					}
					word = append(word, line[j])
					if num, ok := numbers[string(word)]; ok {
						wn = num
						break
					}
				}
				if wn != "" {
					c = wn
				} else {
					continue
				}

			}

			if first == "" {
				first = string(c)
			} else {
				last = string(c)
			}
		}
		if last == "" {
			last = first
		}
		num, _ := strconv.Atoi(string(first) + string(last))
		fmt.Println(" num:", num)
		sum += num

	}
	fmt.Printf("the result is %d\n", sum)
}

var numbers = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}
