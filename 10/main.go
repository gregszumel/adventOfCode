package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


type CPU struct {
    cycle int
    Xval int
    criticalValues []int
    CRTScreen []string
}

func (c *CPU) checkForCriticalValues() {
    if c.cycle == 20 || ((c.cycle - 20) % 40 == 0) {
        c.criticalValues = append(c.criticalValues, c.Xval * c.cycle)
    }
}

func (c *CPU) updateCRT() {
    val := c.Xval 
    position := (c.cycle - 1)% 40

    if val-1 == position || val == position || val + 1 == position {
        c.CRTScreen = append(c.CRTScreen, "#")
    } else {
        c.CRTScreen = append(c.CRTScreen, ".")
    }
}

func (c *CPU) printCRT() {
    counter := 0
    for _, val := range c.CRTScreen {
        fmt.Print(val)
        counter += 1
        if counter == 40 {
            counter = 0
            fmt.Print("\n")

        }
    }
    fmt.Print("\n")

}

func (c *CPU) incrementCycle() {
    c.cycle = c.cycle + 1
    c.checkForCriticalValues()
    c.updateCRT()
}

func (c *CPU) addOp(amt int) {
    c.incrementCycle()
    c.incrementCycle()
    c.Xval += amt 
}

func (c *CPU) NoOp() {
    c.incrementCycle()
}

func (c *CPU) sumCriticalSignals()int {
    sum := 0
    for _, sig := range c.criticalValues {
        sum += sig
    }
    return sum
}


func main() {
    cpu := CPU{Xval: 1}

    dat, err := os.ReadFile("input")
    if err != nil { panic(err) }
    
    instructions := strings.Split(string(dat), "\n")
    for _, instruction := range instructions {
        if instruction == "" {

        } else if instruction == "noop" { 
            cpu.NoOp()
        } else {
            amtStr := strings.Split(instruction, " ")[1]
            amt, err := strconv.Atoi(amtStr)
            if err != nil {panic(err)}
            cpu.addOp(amt)
        }
    }
    fmt.Println(cpu.sumCriticalSignals())
    cpu.printCRT()
}
