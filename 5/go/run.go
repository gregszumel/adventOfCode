package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInitState(initStateString string) map[int]*Stack {
    m := make(map[int]*Stack)
    for i:=1; i <= 9 ; i++ { m[i] = &Stack{} }

    rows := strings.Split(initStateString, "\n")

    for i:=len(rows)-1; i >= 0; i-- {
        // skip last row
        if strings.ContainsAny(rows[i], "0123456789") { continue }
        // each line has 35 characters. 
        for j := 0; j*4+1<35; j++ {
            targetChar := string(rows[i][j*4+1])
            if targetChar != " " { m[j+1].Push(targetChar) }
        }
    }
    return m
}

func processInstruction(instructionStr string) (int, int, int) {
    // "move 2 from 9 to 8"
    instructionsSlice := strings.Split(instructionStr, " ")
    loop, err := strconv.Atoi(instructionsSlice[1])
    if err != nil { panic(err) }
    from, err := strconv.Atoi(instructionsSlice[3])
    if err != nil { panic(err) }
    to, err := strconv.Atoi(instructionsSlice[5])
    return loop, from, to
}

func handleInstruction(stacks map[int]*Stack, loop, from, to int) {

    for i:=0; i<loop; i++ {
        val, success := stacks[from].Pop()
        if !success { panic("can't pop when handling instruction") }
        stacks[to].Push(val)
    }
}


func handleInstructionPart2(stacks map[int]*Stack, loop, from, to int) {

    intermediateStack := Stack{}
    for i:=0; i<loop; i++ {
        val, success := stacks[from].Pop()
        if !success { panic("can't pop when handling instruction") }
        intermediateStack.Push(val)
    }

    for i:=0; i<loop; i++ {
        val, success := intermediateStack.Pop()
        if !success { panic("can't pop when handling instruction") }
        stacks[to].Push(val)
    }

}

func main() {
    dat, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }

    lines := strings.Split(string(dat), "\n\n")
    initStateStr := lines[0]
    instructions := lines[1]

    stacks := getInitState(initStateStr)

    for _, instruction := range strings.Split(instructions, "\n") {
        if len(instruction) > 0 {
            loop, from, to := processInstruction(instruction)
            handleInstruction(stacks, loop, from, to)
        }
    }

    fmt.Print("Part 1 solution: ")
    for i:=1; i<=9; i++ {
        val, success := stacks[i].Peek()
        if !success {panic("issue with success")}
        fmt.Print(val)

    }
    fmt.Print("\n")


    /// PART 2

    stacks = getInitState(initStateStr)
    for _, instruction := range strings.Split(instructions, "\n") {
        if len(instruction) > 0 {
            loop, from, to := processInstruction(instruction)
            handleInstructionPart2(stacks, loop, from, to)
        }
    }
    fmt.Print("Part 2 solution: ")
    for i:=1; i<=9; i++ {
        val, success := stacks[i].Peek()
        if !success {panic("issue with success")}
        fmt.Print(val)

    }
    fmt.Print("\n")

}
