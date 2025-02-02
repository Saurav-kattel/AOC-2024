package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func main() {

	repeatCount := map[int]int{}
	isCounted := map[int]bool{}

	location, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}

	leftTeam := []int{}
	rightTeam := []int{}

	data := strings.Join(strings.Split(string(location), " "), " ")
	columns := strings.Split(data, "\n")
	for _, rows := range columns {
		if len(rows) <= 0 {
			continue
		}
		fixedString := standardizeSpaces(rows)
		seperatedRows := strings.Split(fixedString, " ")

		leftTeamId, _ := strconv.Atoi(seperatedRows[0])
		rightTeamId, _ := strconv.Atoi(seperatedRows[1])
		leftTeam = append(leftTeam, leftTeamId)
		rightTeam = append(rightTeam, rightTeamId)
	}

	sort.Slice(leftTeam, func(i, j int) bool {
		return leftTeam[i] <= leftTeam[j]
	})

	sort.Slice(rightTeam, func(i, j int) bool {
		return rightTeam[i] <= rightTeam[j]
	})

	arrayLength := len(rightTeam)
	for i := 0; i < arrayLength; i++ {
		if _, ok := isCounted[leftTeam[i]]; ok {
			continue
		}
		for j := 0; j < arrayLength; j++ {
			isCounted[leftTeam[i]] = true
			currentLeftId := leftTeam[i]
			if currentLeftId == rightTeam[j] {
				repeatCount[currentLeftId]++
			}
		}
	}

	sum := 0
	for _, nums := range leftTeam {
		count, _ := repeatCount[nums]
		sum += (nums * count)
	}

	fmt.Println(sum)

}

/*
 -> check the repetion of left[i] in right[]
		HashMap to store the repetion count
		loop till the array length;
		Use LeftId as hashkey
		every iteration check if the rightId is in the hashmap,
*/
