package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scn := bufio.NewScanner(f)

	card := 0
	copies := make([]int, 202)
	for scn.Scan() {
		line := scn.Text()
		copies[card] += 1
		matches := getLineMatches(line)

		for i := 1; i <= matches && i < 202; i++ {
			copies[card+i] += 1 * copies[card]
		}

		fmt.Printf("card %d: matches %d\n", card, matches)
		fmt.Printf("deck: %v\n", copies)
		card += 1
	}

	sum := 0
	for _, x := range copies {
		sum += x
	}
	fmt.Println("total:", sum)
}

func getLineMatches(line string) int {
	matches := 0
	parts := strings.Split(line, "|")
	wining := strings.Split(strings.TrimSpace(strings.Split(parts[0], ":")[1]), " ")
	numbers := strings.Split(strings.TrimSpace(parts[1]), " ")
	for _, number := range numbers {
		if number == " " || number == "" {
			continue
		}

		for _, w := range wining {
			if number == w {
				matches += 1
				break
			}
		}
	}

	return matches
}
