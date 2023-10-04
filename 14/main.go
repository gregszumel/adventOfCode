package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    f, _ := os.ReadFile("input")
    rockMap, maxY := parseInput(f)
    part1Counter := 0
    for true {
        if !addGrainPart1(&rockMap, maxY) {
            break
        }
        part1Counter++
    }
    fmt.Println(part1Counter)

    rockMap, maxY = parseInput(f)
    part2Counter := 0
    for true {
        if !addGrainPart2(&rockMap, maxY) {
            part2Counter++
            break
        }
        part2Counter++
    }
    fmt.Println(part2Counter)
}

func parseInput(f []byte) (map[int]map[int]int, int)  {
    input := strings.TrimRight(string(f), "\n")
    rockMap := make(map[int]map[int]int)
    maxY := 0
    for _, line := range strings.Split(input, "\n") {
        var prevX, prevY = -1, -1
        for _, coord := range strings.Split(line, " -> ") {
            coords := strings.Split(coord, ",")
            x, _ := strconv.Atoi(coords[0])
            y, _ := strconv.Atoi(coords[1])
            if y > maxY { maxY = y }
            _, mapMade := rockMap[x]
            if !mapMade { rockMap[x] = make(map[int]int) }

            rockMap[x][y] = 1

            if prevX != -1 && prevY != -1 {
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
                    _, mapMade := rockMap[prevX]
                    if !mapMade { rockMap[prevX] = make(map[int]int) }
                    rockMap[prevX][prevY] = 1
                }
            } else {
                prevX = x; prevY = y;

            }
        }
    }
    return rockMap, maxY
}


func addGrainPart1(rockMap *map[int]map[int]int, maxY int) bool {
    sandX := 500
    sandY := 0

    for true {
        if sandY > maxY + 1 {
            return false
        } else if (*rockMap)[sandX][sandY + 1] == 0 {
            sandY++
        } else if (*rockMap)[sandX - 1][sandY + 1] == 0 {
            sandY++
            sandX--
        } else if (*rockMap)[sandX + 1][sandY + 1] == 0 {
            sandY++
            sandX++
        } else {
            (*rockMap)[sandX][sandY] = 2
            return true
        }
    }
    panic("never terminated")
}

func addGrainPart2(rockMap *map[int]map[int]int, maxY int) bool {
    sandX := 500
    sandY := 0

    for true {
        if sandY > maxY  {
            _, exists := (*rockMap)[sandX]
            if !exists { (*rockMap)[sandX] = make(map[int]int) }
            (*rockMap)[sandX][sandY] = 2
            return true
        } else if (*rockMap)[sandX][sandY + 1] == 0 {
            sandY++
        } else if (*rockMap)[sandX - 1][sandY + 1] == 0 {
            sandY++
            sandX--
        } else if (*rockMap)[sandX + 1][sandY + 1] == 0 {
            sandY++
            sandX++
        } else {
            (*rockMap)[sandX][sandY] = 2
            return !(sandX == 500 && sandY == 0)
        }
    }
    panic("never terminated")
}
