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


func main() {
    runningSum := 0
    dat, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    unsplitRucksacks := strings.Split(string(dat), "\n")
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

    fmt.Println(runningSum)

}
