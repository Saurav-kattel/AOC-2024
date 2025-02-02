package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func checkIsSafe(individualData []string) (int, bool) {
	isIncreasing := false
	isFirstIteration := true
	for i := 0; i < len(individualData)-1; i++ {

		firstVal, _ := strconv.Atoi(individualData[i])
		secondVal, _ := strconv.Atoi(individualData[i+1])

		if isFirstIteration {
			isFirstIteration = false
			isIncreasing = (firstVal - secondVal) < 0
		}

		condition := int(math.Abs(float64(firstVal) - float64(secondVal)))
		// if difference is 0 or if difference is greater then 3
		if (condition == 0) || (condition > 3) {
			return i, false
		}

		if isIncreasing && firstVal-secondVal > 0 {
			return i, false
		}

		if !isIncreasing && firstVal-secondVal < 0 {
			return i, false
		}

	}
	return 1, true
}

func main() {

	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("unable to read file")
		os.Exit(1)
	}

	safeCountA := 0
	safeCountB := 0
	reports := strings.Split(string(data), "\n")
	for _, report := range reports {
		data := strings.Split(report, " ")
		_, isSafe := checkIsSafe(data)
		if isSafe {
			safeCountA++
		} else {
			for i := 0; i < len(data); i++ {
				sliced := make([]string, 0, len(data))
				sliced = append(sliced, data[:i]...)
				sliced = append(sliced, data[i+1:]...)
				_, isSecondSafe := checkIsSafe(sliced)
				fmt.Println(isSecondSafe)
				if isSecondSafe {
					safeCountB++
					break
				}
			}
		}
	}

	fmt.Println(safeCountA, safeCountA+safeCountB)
}
