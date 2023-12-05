package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"strconv"
)

type mapRange struct {
	min  int
	max  int
	dest int
}

func toInt(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal("Input isn't int")
	}
	return val
}

func PartA() {

	file, err := os.Open("2023/5/day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	seeds := make([]string, 0, 20)

	maps := make([][]mapRange, 0, 10)

	first := true
	mapCount := -1
	for scanner.Scan() {

		txt := scanner.Text()

		if first {
			//We're on seeds part
			seeds = strings.Fields(txt)[1:]
			first = false
			continue
		}

		if txt == "" {
			continue
		}

		if strings.Contains(txt, ":") {
			mapCount += 1
			maps = append(maps, make([]mapRange, 0, 5))
			continue
		}

		items := strings.Fields(txt)
		maps[mapCount] = append(maps[mapCount], mapRange{min: toInt(items[1]),
			max: toInt(items[1]) + toInt(items[2]) - 1, dest: toInt(items[0])})

	}

	min := 0

	for _, seed := range seeds {

		loc := getFinalLocation(toInt(seed), maps)
		if min == 0 {
			min = loc
		}

		if loc < min {
			min = loc
		}

	}

	fmt.Println(min)

}

func PartB() {

	file, err := os.Open("2023/5/day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	seeds := make([]string, 0, 20)

	maps := make([][]mapRange, 0, 10)

	first := true
	mapCount := -1
	for scanner.Scan() {

		txt := scanner.Text()

		if first {
			//We're on seeds part
			seeds = strings.Fields(txt)[1:]
			first = false
			continue
		}

		if txt == "" {
			continue
		}

		if strings.Contains(txt, ":") {
			mapCount += 1
			maps = append(maps, make([]mapRange, 0, 5))
			continue
		}

		items := strings.Fields(txt)
		maps[mapCount] = append(maps[mapCount], mapRange{min: toInt(items[1]),
			max: toInt(items[1]) + toInt(items[2]) - 1, dest: toInt(items[0])})

	}

	min := 0

	seedRange := make([]*mapRange, len(seeds)/2)

	idx := 0
	snd := false
	for _, seed := range seeds {

		if !snd {
			seedRange[idx] = &mapRange{min: toInt(seed)}
			snd = true
		} else {
			seedRange[idx].max = seedRange[idx].min + toInt(seed) - 1
			snd = false
			idx += 1
		}
	}

	for _, rangeSeed := range seedRange {

		for i := rangeSeed.min; i <= rangeSeed.max; i++ {

			loc := getFinalLocation(i, maps)
			if min == 0 {
				min = loc
			}

			if loc < min {
				min = loc
			}

		}

	}

	fmt.Println(min)
}

func getFinalLocation(source int, maps [][]mapRange) int {
	dest := source
	// Can cheat by assuming that the maps are always in order
	for _, m := range maps {
		dest = getDestination(dest, m)
	}
	return dest
}

func getDestination(source int, m []mapRange) int {

	dest := 0

	for _, r := range m {

		if r.min <= source && r.max >= source {
			moved := source - r.min
			dest = r.dest + moved
		}
	}

	// If we didn't find anything then, the source equals the destination
	if dest == 0 {
		dest = source
	}

	return dest

}
