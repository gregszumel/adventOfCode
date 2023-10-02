package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ListNode struct {
    val int;
    hasVal bool;

    subList []*ListNode;
}

func compareLists(l1, l2 []*ListNode) (bool, bool) {
    for i:=0; i < min(len(l1), len(l2)); i++ {
        if (l1[i].hasVal && l2[i].hasVal) && l1[i].val < l2[i].val {
            return true, false
        } else if (l1[i].hasVal && l2[i].hasVal) && l1[i].val > l2[i].val { 
            // l1 and l2 are not in the right order
            return false, false
        } else if l1[i].hasVal && !l2[i].hasVal {
            // l1 is a number, l2 is a list
            // create a sublist
            l1SubList := []*ListNode{ { val: l1[i].val, hasVal: true} }
            inOrder, isTied := compareLists(l1SubList, l2[i].subList) 
            if !isTied { return inOrder, false }
        } else if !l1[i].hasVal && l2[i].hasVal {
            // l1 is a list, l2 is a number
            // create a sublist
            l2SubList := []*ListNode{ { val: l2[i].val, hasVal: true} }
            inOrder, isTied := compareLists(l1[i].subList, l2SubList) 
            if !isTied { return inOrder, false }
        } else if !l1[i].hasVal && !l2[i].hasVal {
            // l1 and l2 are lists
            inOrder, isTied := compareLists(l1[i].subList, l2[i].subList)
            if !isTied { return inOrder, false }
        }
    }

    if len(l1) == len(l2) {
        return false, true
    } else {
        return len(l1) < len(l2), false
    }
}

func findClosingBracket(listStr string) int {
    i := 0
    bracketCounter := 1;

    for i < len(listStr) {
        elStr := string(listStr[i])
        i++;
        if elStr == "[" { 
            bracketCounter += 1 
        } else if elStr == "]" {
            bracketCounter -= 1
            if bracketCounter == 0 { break }
        }
    }
    return i 
}

func createList(listStr string) []*ListNode {
    list := make([]*ListNode, 0)
    i := 0
    for i < len(listStr) {
        elStr := string(listStr[i])
        if elStr == "," { 
            i += 1
        } else if elStr == "[" {
            subList := createList(listStr[i+1:])
            node := ListNode{subList: subList}
            list = append(list, &node)
            i += findClosingBracket(listStr[i+1:]) + 1
        } else if strings.ContainsAny(elStr, "0123456789") {
            val, err := strconv.Atoi(elStr) 
            if err != nil { panic(err) }
            node := ListNode{ val: val, hasVal: true }
            list = append(list, &node)
            i += 1
        } else if elStr == "]" {
            return list
        } else {
        }
    }
    return list;
}

func createDataPairs(l1Str, l2Str string) ([]*ListNode, []*ListNode) {
    return createList(l1Str[1:]), createList(l2Str[1:])
}

func printList(nodes []*ListNode) {
    fmt.Print("[")
    for i, node := range nodes {
        if node.hasVal {
            fmt.Print(node.val)
        } else {
            printList(node.subList)
        }

        if i < len(nodes)-1 {
            fmt.Print(",")
        } 
    }
    fmt.Print("]")
}

func main() {
    dat, err := os.ReadFile("input.txt")
    if err != nil { panic(err) }
    sum := 0
    pairs := strings.Split(string(dat), "\n\n")
    for i, pair := range pairs {
        lists := strings.Split(pair, "\n")
        l1, l2 := createDataPairs(lists[0], lists[1])
        // fmt.Println(pair)
        inOrder, isTied := compareLists(l1, l2)
        if isTied { panic("Tied result") }

        if inOrder {
            // fmt.Println("in order")
            sum += (i  + 1)
        } else {
            // fmt.Println("not in order")
        }
    }
    fmt.Println(sum)
}

