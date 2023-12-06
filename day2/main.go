package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// currently ignore memory load ^^
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
	proSum := 0
	scn := bufio.NewScanner(f)
	for scn.Scan() {
		line := scn.Text()
		info := strings.Split(line, ":")
		summary := sumariseColors(info[1])
		if isPossible(summary) {
			id, _ := strconv.Atoi(strings.Split(info[0], " ")[1])
			sum += id
		}

		proSum += getSummaryProduct(summary)
	}

	fmt.Println("sum: ", sum, "product:", proSum)
}

func isPossible(summary map[string]int) bool {
	for color, qty := range summary {
		switch color {
		case "red":
			if qty > 12 {
				return false
			}
		case "green":
			if qty > 13 {
				return false
			}

		case "blue":
			if qty > 14 {
				return false
			}
		}
	}
	return true
}

func sumariseColors(colorsLine string) map[string]int {
	sumary := make(map[string]int)
	for _, set := range strings.Split(colorsLine, ";") {
		for _, cubes := range strings.Split(set, ",") {
			info := strings.Split(strings.TrimSpace(cubes), " ")
			qty, _ := strconv.Atoi(info[0])
			prev, exists := sumary[info[1]]
			if exists && qty > prev {
				sumary[info[1]] = qty
			} else if !exists {
				sumary[info[1]] = qty
			}
		}
	}
	return sumary
}

func getSummaryProduct(summary map[string]int) int {
	p := 1
	for _, qty := range summary {
		p *= qty
	}
	return p
}
