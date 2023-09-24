package main
import "fmt"

type Node struct {
    val string;
    next *Node
}

type Stack struct {
    head *Node
    size int
}


func (s *Stack) Size()int  {
    return s.size
}

func (s *Stack) IsEmpty()bool {
    return s.size == 0
}

func (s *Stack) Peek()(val string, success bool) {
    if s.IsEmpty() {
        fmt.Println("StackEmptyException")
        return "", false
    }
    return s.head.val, true
}

func (s *Stack) Push(val string) {
    s.head = &Node{val: val, next: s.head}
    s.size++
}

func (s *Stack) Pop()(val string, success bool) {
    if s.IsEmpty() {
        fmt.Println("StackEmptyException")
        return "", false
    }
    ret := s.head.val
    s.head = s.head.next
    s.size--
    return ret, true
}

func (s *Stack) Print() {
    temp := s.head
    fmt.Print("Values stored in stack are: ")
    for temp != nil {
        fmt.Print(temp.val, " ")
        temp = temp.next
    }
    fmt.Println()
}
