package main

import (
	"fmt"
	"os"
	"strings"
)

func getVal(num int) int {
    if num > 96 {
        return num - 96
    } else {
        return num - 38
    }
}

func bothContains(letter, sack2, sack3 string) bool {
    return strings.Contains(sack2, letter) && strings.Contains(sack3, letter)
}


func main() {
    runningSum := 0
    dat, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    unsplitRucksacks := strings.Split(string(dat), "\n")

    // Part 1
    for i:=0; i < len(unsplitRucksacks); i++ {
        targetRucksack := unsplitRucksacks[i]
        left := targetRucksack[:len(targetRucksack) / 2]
        right := targetRucksack[len(targetRucksack) / 2:]

        for i:=0; i < len(left); i++ {
            if strings.Contains(right, string(left[i])) {
                runningSum += getVal(int(left[i]))
                break;
            }
        }
    }

    fmt.Println("part 1 solution:", runningSum)

    // Part 2
    part2Sum := 0
    for i:=0; i < len(unsplitRucksacks) - 3; i = i + 3 {
        sack1 := unsplitRucksacks[i]
        sack2 := unsplitRucksacks[i+1]
        sack3 := unsplitRucksacks[i+2]
        for letterIdx:=0; letterIdx < len(sack1); letterIdx++ {
            targetLetter := string(sack1[letterIdx])
            if bothContains(targetLetter, sack2, sack3) {
                part2Sum += getVal(int(sack1[letterIdx]))
                break;
            }
        }
    }
    fmt.Println("part 2 solution:", part2Sum)
}
