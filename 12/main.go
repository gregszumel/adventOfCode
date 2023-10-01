package main

import (
	"fmt"
	"os"
	"strings"
)

type Tuple struct { x, y int; }

type GraphNode struct {
    val, x, y int;
    parent *GraphNode;
    explored bool
}

func getAdjacentNodes(graph [][]*GraphNode, node *GraphNode) []*GraphNode {
    nodes := make([]*GraphNode, 0)
    if node.x > 0 {
        nodes = append(nodes, graph[node.x-1][node.y])
    }
    if node.x < len(graph) - 1 {
        nodes = append(nodes, graph[node.x+1][node.y])
    }
    if node.y > 0 {
        nodes = append(nodes, graph[node.x][node.y-1])
    }
    if node.y < len(graph[0]) - 1 {
        nodes = append(nodes, graph[node.x][node.y+1])
    }
    return nodes
}


func BFS(graph [][]*GraphNode, root *GraphNode) *GraphNode {
    queue := make([]*GraphNode, 0)
    root.explored = true
    queue = append(queue, root)

    for len(queue) > 0 {
        node := queue[0]

        if node.val == 27 { return node }

        for _, adjNode := range getAdjacentNodes(graph, node) {
            if !adjNode.explored && adjNode.val <= node.val + 1 {
                adjNode.explored = true
                adjNode.parent = node
                queue = append(queue, adjNode)
            }
        }
        queue = queue[1:]
    }
    return root
}

func printSoln(node *GraphNode) {
    for node.parent != nil {
        fmt.Println(node)
        node = node.parent
    }
    fmt.Println(node)
}

func getSolnLen(node *GraphNode) int {
    length := 0
    if node.val != 27 {
        return 9999999999
    }

    for node.parent != nil {
        length += 1
        node = node.parent
    }
    return length
}

func resetExploredGraph(graph [][]*GraphNode) {
    for i:=0; i < len(graph); i++ {
        for j:=0; j < len(graph[i]); j++ {
            node := graph[i][j]
            node.explored = false
            node.parent = nil
        }
    }
}

func main() {
    dat, err := os.ReadFile("input")
    if err != nil { panic(err) }
    field := make([][]*GraphNode, 0)
    var root *GraphNode
    var potentialRoots []*GraphNode
    for i, row := range strings.Split(string(dat), "\n") {
        if row == "" { continue }
        elements := make([]*GraphNode, len(row))
        for j:=0; j < len(elements); j++ {
            var val int
            if string(row[j]) == "E" {
                val = int(byte('z')) + 1
            } else if string(row[j]) == "S" {
                val = int(byte('a')) - 1
            } else {
                val = int(row[j])
            }
            node := GraphNode{val: val - int(byte('a')) + 1, x:i, y:j}
            elements[j] = &node
            if string(row[j]) == "S" { 
                root = &node 
                potentialRoots = append(potentialRoots, &node)
            }
            if string(row[j]) == "a" { 
                potentialRoots = append(potentialRoots, &node)
            }
        }
        field = append(field, elements)
    }

    fmt.Println(getSolnLen(BFS(field, root)))

    minSolnLen := 9999999999
    for _, potentialRoot := range potentialRoots {
        resetExploredGraph(field)
        soln := BFS(field, potentialRoot)
        solnLen := getSolnLen(soln)
        if solnLen < minSolnLen {
            minSolnLen = solnLen
        }
    }
    fmt.Println(minSolnLen)
}

