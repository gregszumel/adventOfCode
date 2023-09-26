package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func initTrees(height, width int) [][]int {
    trees := make([][]int, height)
    for i := 0; i < height; i++ {
        trees[i] = make([]int, width)
    }
    return trees
}

func canSeeFromTop(trees [][]int)[][]int {
    visibleTrees := initTrees(len(trees), len(trees[0]))
    for col := 0; col < len(trees[0]); col++ {
        maxHeight := -1
        for row := 0; row < len(trees); row++ {
            if trees[row][col] > maxHeight {
                visibleTrees[row][col] = 1
                maxHeight = trees[row][col]
            }
        }
    }
    return visibleTrees
}

func canSeeFromLeft(trees [][]int)[][]int {
    visibleTrees := initTrees(len(trees), len(trees[0]))
    for row := 0; row < len(trees); row++ {
        maxHeight := -1
        for col := 0; col < len(trees[0]); col++ {
            if trees[row][col] > maxHeight {
                visibleTrees[row][col] = 1
                maxHeight = trees[row][col]
            }
        }
    }
    return visibleTrees
}

func canSeeFromRight(trees [][]int)[][]int {
    visibleTrees := initTrees(len(trees), len(trees[0]))
    for row := 0; row < len(trees); row++ {
        maxHeight := -1
        for col := len(trees[0]) -1; col >= 0; col-- {
            if trees[row][col] > maxHeight {
                visibleTrees[row][col] = 1
                maxHeight = trees[row][col]
            }
        }
    }
    return visibleTrees
}

func canSeeFromBot(trees [][]int)[][]int {
    visibleTrees := initTrees(len(trees), len(trees[0]))
    for col := 0; col < len(trees[0]); col++ {
        maxHeight := -1
        for row := len(trees) - 1; row >= 0; row-- {
            if trees[row][col] > maxHeight {
                visibleTrees[row][col] = 1
                maxHeight = trees[row][col]
            }
        }
    }
    return visibleTrees
}

func getMaxTrees(trees [][]int) [][]int {
    isVisibleFromAny := initTrees(len(trees), len(trees[0]))
    topVisible := canSeeFromTop(trees)
    leftVisible := canSeeFromLeft(trees)
    rightVisible := canSeeFromRight(trees)
    botVisible := canSeeFromBot(trees)

    for i:=0; i < len(trees); i++ {
        for j:=0; j < len(trees[0]); j++ {
            isVisibleFromAny[i][j] = max(
                topVisible[i][j], 
                leftVisible[i][j],
                rightVisible[i][j],
                botVisible[i][j],
            )
        }
    }
    return isVisibleFromAny
}



func main() {
    dat, err := os.ReadFile("input.txt")
    if err != nil { panic(err) }
    treeRows := strings.Split(string(dat), "\n")
    trees := make([][]int, len(treeRows))

    for i, row := range treeRows {
        cols := strings.Split(row, "")
        trees[i] = make([]int, len(cols))
        for j, col := range cols {
            treeHeight, err := strconv.Atoi(col)
            if err != nil {panic(err)}
            trees[i][j] = treeHeight
        }
    }

    trees = trees[:len(trees) - 1]

    visibleTrees := getMaxTrees(trees)
    sum := 0
    for i := 0; i < len(visibleTrees); i++ {
        for j := 0; j < len(visibleTrees[0]); j++ {
            sum += visibleTrees[i][j]
        }
    }
    fmt.Println(sum)
    fmt.Println(visibleTrees)
    fmt.Println(trees)

}

