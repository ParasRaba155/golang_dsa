package main

import (
	"fmt"
	"log"

	stack "github.com/algorithms/April_04"
)

func main() {
    s := stack.Stack[int]{Capacity: 3}
    fmt.Printf("s = %+v\n",s)
    err := s.Push(32)
    fmt.Printf("s = %+v\n",s)
    if err != nil {
        log.Fatalf("error in pushing : %s",err)
    }
    err = s.Push(12)
    fmt.Printf("s = %+v\n",s)
    if err != nil {
        log.Fatalf("error in pushing : %s",err)
    }
    err = s.Push(45)
    if err != nil {
        log.Fatalf("error in pushing : %s",err)
    }
    fmt.Printf("s = %+v\n",s)
    fmt.Printf("s is empty = %v\n",s.IsEmpty())
    fmt.Printf("s is full = %v\n",s.IsFull())
    // err = s.Push(12)
    if err != nil {
        log.Fatalf("error in pushing : %s",err)
    }
    err = s.Pop()
    fmt.Printf("s = %+v\n",s)
    if err != nil {
        log.Fatalf("error in pushing : %s",err)
    }
    err = s.Pop()
    fmt.Printf("s = %+v\n",s)
    if err != nil {
        log.Fatalf("error in pushing : %s",err)
    }
    err = s.Pop()
    fmt.Printf("s = %+v\n",s)
    if err != nil {
        log.Fatalf("error in pushing : %s",err)
    }
    err = s.Pop()
    fmt.Printf("s = %+v\n",s)
    if err != nil {
        log.Fatalf("error in pushing : %s",err)
    }
}
