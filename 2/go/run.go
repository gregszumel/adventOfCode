package main

import (
	"fmt"
	"os"
	"strings"
)

func score(picks []string) int {
    player1, player2 := picks[0], picks[1]
    switch {
    case player1=="A" && player2=="X": return 1 + 3
    case player1=="B" && player2=="X": return 1 + 0
    case player1=="C" && player2=="X": return 1 + 6

    case player1=="A" && player2=="Y": return 2 + 6
    case player1=="B" && player2=="Y": return 2 + 3
    case player1=="C" && player2=="Y": return 2 + 0

    case player1=="A" && player2=="Z": return 3 + 0
    case player1=="B" && player2=="Z": return 3 + 6
    case player1=="C" && player2=="Z": return 3 + 3
    }

    return -1
}

func main() {
    dat, err := os.ReadFile("input.txt")
    if err != nil { panic(err) }
    gamesStrings := strings.Split(string(dat), "\n")
    scoreSum := 0

    for i:=0; i < len(gamesStrings); i++ {
        gameStr := gamesStrings[i]
        if len(gameStr) > 0 {
            game := strings.Split(gamesStrings[i], " ")
            gameScore := score(game)
            scoreSum += gameScore
            fmt.Println(game, "->", gameScore)
        }
    }

    fmt.Println(scoreSum)
}
