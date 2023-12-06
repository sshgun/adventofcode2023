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
	sum := 0
	for scn.Scan() {
		line := scn.Text()
		points := 0
		parts := strings.Split(line, "|")
		wining := strings.Split(strings.TrimSpace(strings.Split(parts[0], ":")[1]), " ")
		numbers := strings.Split(strings.TrimSpace(parts[1]), " ")
		for _, number := range numbers {
			if number == " " || number == "" {
				continue
			}

			for _, w := range wining {
				if number == w {
					if points == 0 {
						points = 1
					} else {
						points *= 2
					}
					break
				}
			}
		}
		sum += points
	}
	fmt.Println("worth: ", sum)
}
