package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type OperationFn func(int) int;
type TestFn func(int) bool;

type Monkey struct {
    items []int;
    operation OperationFn;
    test TestFn;
    trueTestMonkey int
    falseTestMonkey int
    divisibleBy int;
    inspectedCount int
}


func createOpFn(opStr string, rhs string) OperationFn {

    // define op, which is add or multiply
    op := func(x, y int) int {
        if opStr == "*" { return x * y }
        return x + y
    }

    // define the operationFn, which takes old int as input. 
    // if rhs is "old", then just call op with old as both inputs
    // otherewise, convert the rhs into a value and return op on old and the val
    opFn := func(old int) int {
        if rhs == "old" {
            return op(old, old);
        } else {
            rhsVal, err := strconv.Atoi(rhs)
            if err != nil {panic(err)}
            return op(old, rhsVal)
        }
    }
    return opFn;
}

func createTestFn(divisibleByStr string) (TestFn, int) {
    divisibleBy, err := strconv.Atoi(divisibleByStr)
    if err != nil { panic(err) }

    f := func(val int) bool {
        if val % divisibleBy == 0 {
            return true
        } else {
            return false
        }
    }

    return f, divisibleBy
}


func splitAndCutAfter(raw, splitText string) string {
    items := strings.Split(raw, splitText)[1]
    items = strings.Trim(items, " ")
    return items
}


func readRawMonkey(monkeyStr string) Monkey {
    monkeyRows := strings.Split(monkeyStr, "\n")
    items := make([]int, 0);
    var opFn OperationFn
    var testFn TestFn
    var trueTestMonkey, falseTestMonkey, divisibleBy int;

    for _, row := range monkeyRows {
        if strings.Contains(row, "Starting items:") {
            itemsStr := splitAndCutAfter(row, "Starting items:")
            for _, item := range strings.Split(itemsStr, ", ") {
                item, err := strconv.Atoi(item)
                if err != nil { panic(err) }
                items = append(items, item)
            }
        } else if strings.Contains(row, "Operation:") {
            fullop := strings.Split(splitAndCutAfter(row, "Operation:"), " ")
            op, rhs := fullop[3], fullop[4]
            opFn = createOpFn(op, rhs)
        } else if strings.Contains(row, "Test:") {
            testVal := splitAndCutAfter(row, "Test: divisible by")
            testFn, divisibleBy = createTestFn(testVal)

        } else if strings.Contains(row, "If true: throw to monkey") {
            monkey := splitAndCutAfter(row, "If true: throw to monkey ")
            monkeyVal, err := strconv.Atoi(monkey)
            if err != nil { panic(err) }
            trueTestMonkey = monkeyVal
        } else if strings.Contains(row, "If false: throw to monkey") {
            monkey := splitAndCutAfter(row, "If false: throw to monkey ")
            monkeyVal, err := strconv.Atoi(monkey)
            if err != nil { panic(err) }
            falseTestMonkey = monkeyVal
        }
    }
    return Monkey{ 
        items: items, operation: opFn, test: testFn, 
        trueTestMonkey:trueTestMonkey, falseTestMonkey: falseTestMonkey, 
        divisibleBy: divisibleBy,
    }
}

func doMonkeyTurn(i int, monkeys []Monkey, mod int) []Monkey {
    monkey := monkeys[i]
    // fmt.Println("Monkey", i)
    for _, item := range monkey.items {
        // fmt.Println("Monkey", i, "inspects item with worry level", item)
        newItemVal := monkey.operation(item)
        // newItemVal = int(float64(newItemVal) / 3)
        // fmt.Print(newItemVal, "\n")
        newItemVal = newItemVal % mod
        // fmt.Println("Worry level goes from", item, "to", newItemVal)
        if monkey.test(newItemVal) {
            // fmt.Println("Worry level after test is divisible by target, thrown to monkey", monkey.trueTestMonkey)
            monkeys[monkey.trueTestMonkey].items = append(
                monkeys[monkey.trueTestMonkey].items, newItemVal)
        } else {
            // fmt.Println("Worry level after test is NOT divisible by target, thrown to monkey", monkey.falseTestMonkey)
            monkeys[monkey.falseTestMonkey].items = append(
                monkeys[monkey.falseTestMonkey].items, newItemVal)
        }
        monkeys[i].inspectedCount = monkeys[i].inspectedCount + 1
    }
    monkeys[i].items = make([]int, 0)
    return monkeys
}

func main () {
    monkeys := make([]Monkey, 0)
    dat, err := os.ReadFile("input")
    if err != nil { panic(err) }
    rawMonkeys := strings.Split(string(dat), "\n\n")
    mod := 1
    for _, monkey := range rawMonkeys  {
        monkey := readRawMonkey(monkey)
        monkeys = append(monkeys, monkey)
        mod = mod * monkey.divisibleBy
    }

    for round := 0; round < 10000; round++ {
        for i:=0; i < len(monkeys); i++ {
            monkeys = doMonkeyTurn(i, monkeys, mod)
        }

        fmt.Println("After round", round)
        for i:=0; i < len(monkeys); i++ {
            fmt.Println("Monkey", i, "holding:", monkeys[i].items)
        }
    }
    for i:=0; i < len(monkeys); i++ {
        fmt.Println("Monkey", i, "inspected", monkeys[i].inspectedCount, "times")
    }

}
