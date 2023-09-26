package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


type FileNode struct {
    name string
    size int
}

type DirNode struct {
    name string
    size int
    parent *DirNode
    expanded bool
    files []*FileNode
    dirs []*DirNode
}


type Stream struct {
    index int
    length int
    dat []string
}

func (s *Stream) next() string {
    if s.index < s.length {
        ret := s.dat[s.index]
        s.index += 1
        return ret
    } else {
        return ""
    }
}

func (s *Stream) peek() string {
    if s.index < s.length {
        return s.dat[s.index]
    } else {
        return ""
    }
}

func (s *Stream) hasNext() bool {
    return s.index < s.length
}

func initStream(s string) Stream {
    dat := strings.Split(s, "\n")
    datLen := len(dat) - 1
    return Stream{ index: 0, length: datLen, dat: dat}
}

func initDirNode() *DirNode {
    root := DirNode {
        name: "/",
        expanded: false,
    }
    rootList := []*DirNode { &root }
    return &DirNode{ 
        name:"top",
        expanded: true,
        dirs: rootList,
    }
}

func handleNext(s *Stream, currentNodePointer *DirNode) *DirNode {
    next := s.next()
    if next == "$ ls" { 
        handleLs(s, currentNodePointer)
    } else if next == "$ cd .." { 
        currentNodePointer = handleCdUp(s, currentNodePointer)
    } else if len(next) > 4 && next[0:4] == "$ cd" {
        currentNodePointer = handleCdDown(s, currentNodePointer, next)
    } else {
        fmt.Println("else:", next)
    }
    return currentNodePointer
}

func handleLs(s *Stream, currentNodePointer *DirNode) {
    for s.hasNext() && string(s.peek()[0]) != "$" {
        next := s.next() 
        if next[0:3] == "dir" {
            name := strings.Split(next, " ")
            d := DirNode{
                name: name[1], 
                expanded: false, 
                parent: currentNodePointer,
            }

            currentNodePointer.dirs = append(currentNodePointer.dirs, &d)
        }  else {
            name := strings.Split(next, " ")
            size, err := strconv.Atoi(name[0])
            if err != nil { panic(err) }
            f := FileNode{name: name[1], size: size}
            currentNodePointer.files = append(currentNodePointer.files, &f)
        }
    }
    currentNodePointer.expanded = true;
}

func handleCdUp(s *Stream, currentNodePointer *DirNode) *DirNode {
    d := currentNodePointer.parent
    return d
}

func handleCdDown(s *Stream, currentNodePointer *DirNode, next string) *DirNode {
    targets := strings.Split(next, " ")
    name := targets[2]

    for i:=0; i < len(currentNodePointer.dirs); i++ {
        d := currentNodePointer.dirs[i]
        if d.name == name  { return d }
    }
    panic("Couldn't find subdir")
}

func getSize(parent *DirNode)int {
    sum := 0
    for _, f := range parent.files { sum += f.size }
    for _, child := range parent.dirs { sum += getSize(child) }
    parent.size = sum
    if parent.name == "/" {
        fmt.Println(parent.name, parent.size)
    } else {
        fmt.Println(parent.parent.name, "/", parent.name, parent.size)
    }
    return sum
}


func part1(parent *DirNode)int {
    sum := 0
    if parent.size < 100000 {
        sum += parent.size
        fmt.Println(parent.parent.name, "/", parent.name, parent.size)
    } 
    for _, child := range parent.dirs { sum += part1(child) }
    return sum
}


func part2(parent *DirNode)int {
    neededToDelete := 45717263 - (70000000 - 30000000)
    filtered := make([]int, 0)

    for _, child := range parent.dirs {
        if child.size >= neededToDelete {
            minFromChild := part2(child)
            filtered = append(filtered, minFromChild)
        }
    }

    filtered = append(filtered, parent.size)
    lowest := 70000000
    for _, potentialMin := range filtered {
        if potentialMin < lowest {
            lowest = potentialMin
        }
    }
    return lowest
}

func main() {
    dat, err := os.ReadFile("input.txt")
    if err != nil { panic(err) }
    s := initStream(string(dat))
    var currentNodePointer *DirNode
    currentNodePointer = initDirNode()
    topNode := currentNodePointer
    for s.hasNext() {
        currentNodePointer = handleNext(&s, currentNodePointer)
    }

    rootDirNode := topNode.dirs[0]
    fmt.Println(getSize(rootDirNode))
    fmt.Println(rootDirNode.size)
    fmt.Println(part1(rootDirNode))
    fmt.Println("needed:", 45717263 - (70000000 - 30000000) )
    fmt.Println(part2(rootDirNode))
}

