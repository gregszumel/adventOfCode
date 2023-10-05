package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main () {
    dat, _ := os.ReadFile("input")
    sensors, maxX, minX := parseInput(dat)
    cantContainCounter := 0
    i := minX
    y := 2000000
    // find First notInRage element 

    for true {
        if inRadiusOfAny(sensors, i, y) {
            i--;
        } else {
            break;
        }
    }

    for true {
        inRadius := inRadiusOfAny(sensors, i, y)
        isOnAny := isOnAny(sensors, i, y)
        if inRadius && !isOnAny {
            cantContainCounter += 1
        } else if i > maxX {
            break
        } else {
        }
        i++
    }

    fmt.Println(cantContainCounter)
}

func parseInput(dat []byte) (sensors []Sensor, maxX, minX int) {
    datStr := strings.TrimRight(string(dat), "\n")
    for _, line := range strings.Split(datStr, "\n") {
        line = line[10:]
        sensorBeacon := strings.Split(line, ": closest beacon is at ")
        sensor, beacon := sensorBeacon[0], sensorBeacon[1]
        x, y := parseXY(sensor)
        bx, by := parseXY(beacon)
        maxX = max(x, bx, maxX) 
        minX = min(y, by, minX)
        sensors = append(sensors, initSensor(x, y, bx, by))
    }
    return sensors, maxX, minX
}

func parseXY(xyStr string) (x int, y int) {
    xySplit := strings.Split(xyStr, ", ")
    x, _ = strconv.Atoi(xySplit[0][2:])
    y, _ = strconv.Atoi(xySplit[1][2:])
    return x, y
}


type Sensor struct {
    x, y, bx, by, d int
}

func (s Sensor) inRadius(x, y int) bool {
    distance := Abs(s.x - x) + Abs(s.y - y)
    return s.d >= distance
}

func inRadiusOfAny(sensors []Sensor, x, y int) bool {
    for _, sensor := range sensors {
        if sensor.inRadius(x, y) {
            return true
        }
    }
    return false
}

func isOnAny(sensors []Sensor, x, y int) bool {
    for _, sensor := range sensors {
        if sensor.bx == x && sensor.by == y {
            return true
        } else if sensor.x == x && sensor.y == y {
            return true
        }
    }
    return false
}


func initSensor(x, y, bx, by int) Sensor {
    distance := Abs(x - bx) + Abs(y - by)
    return Sensor{x: x, y: y, bx: bx, by: by, d: distance}
}

func Abs(x int) int {
    if x < 0 { 
        return -x
    } else {
        return x
    }
}

