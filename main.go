package main

import (
	"fmt"
	"log"

	stack "github.com/algorithms/April_04"
	queue "github.com/algorithms/April_07"
)

func main() {
    // stackExample()
    queueExample()
}

func queueExample() {
    q := queue.NewQueue[int](5)
    fmt.Println(q)
    err := q.Enqueue(3)
    if err != nil {
        log.Fatalf("error in Enqueu(3) : %s", err)
    }
    err = q.Enqueue(15)
    if err != nil {
        log.Fatalf("error in Enqueu(3) : %s", err)
    }
    err = q.Enqueue(18)
    if err != nil {
        log.Fatalf("error in Enqueu(3) : %s", err)
    }
    err = q.Enqueue(16)
    if err != nil {
        log.Fatalf("error in Enqueu(3) : %s", err)
    }
    err = q.Enqueue(19)
    if err != nil {
        log.Fatalf("error in Enqueu(3) : %s", err)
    }
    err = q.Enqueue(45)
    if err != nil {
        log.Printf("error in Enqueue 45 : %s",err)
    }
    fmt.Println(q)
    err = q.Dequeue()
    if err != nil {
        log.Printf("1st deque : %s",err)
    }
    fmt.Println(q)
    err = q.Dequeue()
    if err != nil {
        log.Printf("2nd deque : %s",err)
    }
    fmt.Println(q)
    err = q.Dequeue()
    if err != nil {
        log.Printf("3rd deque : %s",err)
    }
    fmt.Println(q)
    err = q.Dequeue()
    if err != nil {
        log.Printf("4th deque : %s",err)
    }
    fmt.Println(q)
    err = q.Dequeue()
    if err != nil {
        log.Printf("5th deque : %s",err)
    }
    fmt.Println(q)
    err = q.Dequeue()
    if err != nil {
        log.Printf("6th deque : %s",err)
    }
    fmt.Println(q)
}

func stackExample() {
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
