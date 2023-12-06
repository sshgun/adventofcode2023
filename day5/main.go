package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	result := part2()
	fmt.Println("near:", result)
}

func part2() int {

	seedsRanges := make([]int, 0)

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
					seedsRanges = append(seedsRanges, number)
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

			newSeedsRanges := make([]int, 0)

			for i := 0; i < len(seedsRanges)-1; i += 2 {

				seedst := seedsRanges[i]
				seedfn := seedst + seedsRanges[i+1]
				intersepted := false

				for _, fm := range farms {

					sourcest := fm[1]
					sourcefn := sourcest + fm[2]

					if !areInterseted(seedst, seedfn, sourcest, sourcefn) {
						continue
					}

					its := getInterseption(seedst, seedfn, sourcest, sourcefn)

					mappedStart := fm[0] + its[0] - fm[1]
					mappedlength := its[1]

					newSeedsRanges = append(newSeedsRanges, mappedStart, mappedlength)
					intersepted = true

					for _, leftover := range getSeedsLeftovers(seedst, seedfn, sourcest, sourcefn) {
						seedsRanges = append(seedsRanges, leftover...)
					}
				}

				if !intersepted {
					newSeedsRanges = append(newSeedsRanges, seedst, seedsRanges[i+1])
				}
			}

			seedsRanges = newSeedsRanges
		}
	}

	near := -1
	for i := 0; i < len(seedsRanges)-1; i += 2 {
		if near == -1 || seedsRanges[i] < near {
			near = seedsRanges[i]
		}
	}

	return near
}

func areInterseted(seedst, seedfn, sourcest, sourcefn int) bool {
	return seedst >= sourcest && seedst < sourcefn ||
		seedfn > sourcest && seedfn <= sourcefn
}

func getSeedsLeftovers(seedst, seedfn, sourcest, sourcefn int) [][]int {
	sets := make([][]int, 0)
	intersept := getInterseption(seedst, seedfn, sourcest, sourcefn)

	// lower leftover
	if seedst < intersept[0] {
		sets = append(sets, []int{seedst, intersept[0] - seedst})
	}

	//upper leftover
	intfn := intersept[0] + intersept[1]
	if seedfn > intfn {
		sets = append(sets, []int{intfn, seedfn - intfn})
	}
	return sets
}

func getInterseption(seedst, seedfn, sourcest, sourcefn int) []int {
	start := max(seedst, sourcest)
	finish := min(seedfn, sourcefn)
	return []int{start, finish - start}
}
