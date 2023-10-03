package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    f, _ := os.ReadFile("test")

    parseInput(f)
}

func parseInput(f []byte)  {
    input := strings.TrimRight(string(f), "\n")
    rockMap := make(map[int]map[int]int)
    for _, line := range strings.Split(input, "\n") {
        var prevX, prevY = -1, -1
        for _, coord := range strings.Split(line, " -> ") {
            coords := strings.Split(coord, ",")
            x, _ := strconv.Atoi(coords[0])
            y, _ := strconv.Atoi(coords[1])
            _, mapMade := rockMap[x]
            if !mapMade { rockMap[x] = make(map[int]int) }

            rockMap[x][y] = 1

            fmt.Println(prevX, prevY)
            if prevX != -1 && prevY != -1 {
                fmt.Println(x, prevX, y, prevY)
                for prevX != x || prevY != y {
                    if prevX != x {
                        if prevX > x {
                            prevX--
                        } else {
                            prevX++
                        }
                    } else {
                        if prevY > y {
                            prevY--
                        } else {
                            prevY++
                        }
                    }
                    fmt.Println(prevX, prevY)
                    _, mapMade := rockMap[prevX]
                    if !mapMade { rockMap[prevX] = make(map[int]int) }
                    rockMap[prevX][prevY] = 1
                }
            } else {
                prevX = x; prevY = y;

            }
            fmt.Println(x, prevX, y, prevY)
                

            fmt.Println(x, y)
            fmt.Println("--------")
        }
    }
}

