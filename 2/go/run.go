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

func score2(picks []string) int{
    player1, outcome := picks[0], picks[1]
    switch {
    case player1=="A" && outcome=="X": return 0 + 3
    case player1=="B" && outcome=="X": return 0 + 1
    case player1=="C" && outcome=="X": return 0 + 2

    case player1=="A" && outcome=="Y": return 3 + 1
    case player1=="B" && outcome=="Y": return 3 + 2
    case player1=="C" && outcome=="Y": return 3 + 3

    case player1=="A" && outcome=="Z": return 6 + 2
    case player1=="B" && outcome=="Z": return 6 + 3
    case player1=="C" && outcome=="Z": return 6 + 1
    }

    return -1
}


func main() {
    dat, err := os.ReadFile("input.txt")
    if err != nil { panic(err) }
    gamesStrings := strings.Split(string(dat), "\n")
    scoreSum := 0
    part2ScoreSum := 0

    for i:=0; i < len(gamesStrings); i++ {
        gameStr := gamesStrings[i]
        if len(gameStr) > 0 {
            game := strings.Split(gamesStrings[i], " ")
            gameScore := score(game)
            scoreSum += gameScore
            score2Sum := score2(game)
            part2ScoreSum += score2Sum
            fmt.Println(game, "->", gameScore, "||", score2Sum)
        }
    }

    fmt.Println(scoreSum)
    fmt.Println(part2ScoreSum)
}
