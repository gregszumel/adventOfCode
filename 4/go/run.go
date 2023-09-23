package main

import (
	"os"
	"strings"
    "strconv"
)

func preprocessString(inp string) (int64, int64, int64, int64) {
    nums := strings.Split(inp, ",")
    leftNums := strings.Split(nums[0], "-")
    rightNums := strings.Split(nums[1], "-")
    leftMin, _ := strconv.ParseInt(leftNums[0], 10, 0)
    leftMax, _ := strconv.ParseInt(leftNums[1], 10, 0)
    rightMin, _ := strconv.ParseInt(rightNums[0], 10, 0)
    rightMax, _ := strconv.ParseInt(rightNums[1], 10, 0)
    return leftMin, leftMax, rightMin, rightMax
}

func noOverlap(inp string) bool {
    leftMin, leftMax, rightMin, rightMax := preprocessString(inp)

    if leftMin > rightMax || rightMin > leftMax {
        return true
    }

    return false
}

func isContained(inp string) bool {
    leftMin, leftMax, rightMin, rightMax := preprocessString(inp)

    if leftMin <= rightMin && leftMax >= rightMax {
        return true
    } else if rightMin <= leftMin && rightMax >= leftMax {
        return true
    } else {
        return false
    }
}

func main() {
    dat, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }

    numFullyContained := 0;
    anyOverlap := 0;


    ranges := strings.Split(string(dat), "\n")
    for _, v := range ranges {
        if len(v) == 0 {
            continue
        }

        print(v)
        if isContained(v) {
            numFullyContained += 1
            print(" | fully contained")
        }

        if !noOverlap(v) {
            print(" | some overlap")
            anyOverlap += 1
        }
        print("\n")
    }
    println("part 1:", numFullyContained)
    println("part 2:", anyOverlap)
}
