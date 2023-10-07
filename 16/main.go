package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


type Valve struct {
    name string
    flow int
    children []string
}

func main() {
    valves := parseInput()
    currValve := "AA"
}

func parseInput() map[string]Valve {
    dat, _ := os.ReadFile("test")
    datStr := strings.TrimRight(string(dat), "\n")
    valves := make(map[string]Valve)
    for _, line := range strings.Split(datStr, "\n") {
        fmt.Println(line)
        valveName := strings.Split(line, " has flow")[0][6:]
        fmt.Println(valveName)
        flowRate := strings.Split(line, "has flow rate=")[1]
        flowRate = strings.Split(flowRate, ";")[0]
        flow, _ := strconv.Atoi(flowRate)
        fmt.Println(flow)
        childrenStrSet := regexp.MustCompile("valves? ").Split(line, -1)
        children := strings.Split(childrenStrSet[len(childrenStrSet)-1], ", ")
        fmt.Println(children)
        valves[valveName] = Valve{
            name: valveName, flow: flow, 
            children: children}
    }

    return valves
}
