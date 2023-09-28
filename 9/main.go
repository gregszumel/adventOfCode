package main

import (
	"fmt"
    "math"
    "os"
    "strconv"
	"strings"
)


type RopeNode struct {
    X int;
    Y int;
    head *RopeNode;
    tail *RopeNode;
    visitedPositions [][]int;
}

func (currNode *RopeNode) printRope() {
    fmt.Println(currNode, &currNode)
    for currNode.tail != nil {
        currNode = currNode.tail
        fmt.Println(currNode, &currNode)
    }
}

func (currNode *RopeNode) getNumVisited() int {
    for currNode.tail != nil {
        currNode = currNode.tail
    }
    return len(currNode.visitedPositions)
}

func (head *RopeNode) movePosition(x, y int) {
    fmt.Println("incrementing head:", x, ",", y, "-> (", head.X, ",", head.Y, ")" )
    head.X += x
    head.Y += y
}

func (tail *RopeNode) moveRopeNode() {
    relX := float64(tail.head.X - tail.X)
    relY :=  float64(tail.head.Y - tail.Y)

    fmt.Println("moving tail, where head pos is:", tail.head.X, ",", tail.head.Y)
    fmt.Print("moved tail from:", tail.X, ",", tail.Y)

    // moving only horizontally - |X| > 1, |Y| == 0
    if math.Abs(relX) == 2 && math.Abs(relY) == 0 {
        tail.X += int(relX / 2)
    // tail moving vertically
    } else if math.Abs(relX) == 0 && math.Abs(relY) == 2 {
        tail.Y += int(relY / 2)
    // moving diagonally
    } else if math.Abs(relX) == 2 && math.Abs(relY) == 1 {
        tail.X += int(relX / 2) 
        tail.Y += int(relY) 

    } else if math.Abs(relX) == 1 && math.Abs(relY) == 2 {
        tail.X += int(relX) 
        tail.Y += int(relY / 2) 
    } else if math.Abs(relX) == 2 && math.Abs(relY) == 2{
        tail.X += int(relX / 2) 
        tail.Y += int(relY / 2) 
    }
    fmt.Print("to ->", tail.X, ",", tail.Y, "\n")
}

func (tail *RopeNode) updateVisitedPositions() {
    if tail.tail != nil { panic("calling updateVisitedPositions on nontail") }

    for _, pos := range tail.visitedPositions {
        if pos[0] == tail.X && pos[1] == tail.Y {
            return 
        }
    }
    newPos := make([]int, 2)
    newPos[0] = tail.X
    newPos[1] = tail.Y
    tail.visitedPositions = append(tail.visitedPositions, newPos)
    fmt.Println("adding new position:", newPos)
}


func handleInstruction(instr string, head *RopeNode) {
    var currNode *RopeNode
    dir, nStr := strings.Split(instr, " ")[0], strings.Split(instr, " ")[1]
    n, err := strconv.Atoi(nStr)
    if err != nil { panic(err) }
    for i:=0; i < n; i++ {
        if dir == "R" {
            head.movePosition(0, 1)
        } else if dir == "L" {
            head.movePosition(0, -1)
        } else if dir == "U" {
            head.movePosition(1, 0)
        } else if dir == "D" {
            head.movePosition(-1, 0)
        }
        currNode = head

        for currNode.tail != nil {
            currNode = currNode.tail
            currNode.moveRopeNode()
        }
        currNode.updateVisitedPositions()
    }
}


func createRope(n int) *RopeNode {
    headPointer := &RopeNode {}
    currNodePointer := headPointer;
    for i:=0; i < (n-1); i++ {
        nextNode := RopeNode { head: currNodePointer }
        currNodePointer.tail = &nextNode
        currNodePointer = &nextNode
    }
    return headPointer
}


func main() {

    head := createRope(10)
    head.printRope()

    dat, err := os.ReadFile("input")
    if err != nil { panic(err) }
    instructions := strings.Split(string(dat), "\n")
    for _, instruction := range instructions {
        if instruction == "" { break }
        fmt.Println(instruction)
        handleInstruction(instruction, head)
        head.printRope()
    }

    fmt.Println("")
    fmt.Println(head.getNumVisited())
}
