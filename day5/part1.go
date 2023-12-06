package main 

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1() {

	seeds := make([]int, 0)

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scn := bufio.NewScanner(f)
	for scn.Scan() {
		line := scn.Text()
		if strings.HasPrefix(line, "seeds:") {
			for _, seedsStrings := range strings.Split(line[7:], " ") {
				if seedsStrings != "" {
					number, _ := strconv.Atoi(seedsStrings)
					seeds = append(seeds, number)
				}
			}
		}

		if strings.HasSuffix(line, "map:") {

			farms := make([][]int, 0)
			for scn.Scan() {
				line = scn.Text()
				if line == "" {
					break
				}

				fm := make([]int, 3)
				for i, nw := range strings.Split(strings.TrimSpace(line), " ") {
					fm[i], _ = strconv.Atoi(nw)
				}
				farms = append(farms, fm)
			}

			for i, seed := range seeds {
				for _, fm := range farms {
					sourceStart := fm[1]
					sourceFinish := sourceStart + fm[2]
					if seed >= sourceStart && seed < sourceFinish {
						newpos := seed - sourceStart

						targetStart := fm[0]
						seeds[i] = targetStart + newpos
						break
					}
				}
			}
		}
	}

	near := slices.Min(seeds)

	fmt.Println("mapped seeds: ", seeds, "near location:", near)
}
