package main

import (
	"fmt"
	"os"
)


func hasRepeats(s string) bool {

    for left := 0; left < len(s) - 1; left ++ {
        for right := left + 1; right < len(s); right++ {
            if s[left] == s[right] { return true }
        }
    }

    return false
}

func main()  {
    dat, err := os.ReadFile("input.txt")
    if err != nil { panic(err) }
    signalStr := string(dat)

    part1Found, part2Found := false, false

    for i:=4; i < len(signalStr); i++ {
        if !part1Found && !hasRepeats(signalStr[i-4:i]) {
            fmt.Println("Part 1 solution:", i, signalStr[i-4: i])
            part1Found = true
        }

        if !part2Found && i-14 >= 0 && !hasRepeats(signalStr[i-14: i]) {
            fmt.Println("Part 2 solution:", i, signalStr[i-14: i])
            part2Found = true
        }
    }
}
