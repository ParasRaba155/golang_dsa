package main

import (
	"fmt"
	"log"

	stack "github.com/algorithms/April_04"
	queue "github.com/algorithms/April_07"
	search "github.com/algorithms/April_15"
)

type customString string

func main() {
	// stackExample()
	// queueExample()
	searchExample()
}

func searchExample() {
	l1 := []int8{0, -5, -4, -8, -2, 45, 64}
	r1 := search.LinearSearch(l1, 45)
	log.Printf("r1 = %d", r1)

	l2 := []customString{"hello", "hi"}
	r2 := search.LinearSearch(l2, "hi")
	log.Printf("r2 = %d", r2)

	l3 := []int8{-5, -5, -4, -2, -1, 5, 89}
	r3 := search.BinarySearch(l3, -5)
	log.Printf("r3 = %d", r3)

	l4 := []string{"abc", "b", "cut", "dr", "ef"}
	r4 := search.BinarySearch(l4, "ff")
	r42 := search.BinarySearch(l4, "dr")
	log.Printf("r4 = %d", r4)
	log.Printf("r42 = %d", r42)

	l5 := []int{10, 12, 13, 16, 18, 19, 20, 21, 22, 23, 24, 33, 35, 42, 47}
	r5 := search.InterpolationSearch(l5, 20, len(l5)-1, 0)
	log.Printf("r5 = %d", r5)
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
		log.Printf("error in Enqueue 45 : %s", err)
	}
	fmt.Println(q)
	err = q.Dequeue()
	if err != nil {
		log.Printf("1st deque : %s", err)
	}
	fmt.Println(q)
	err = q.Dequeue()
	if err != nil {
		log.Printf("2nd deque : %s", err)
	}
	fmt.Println(q)
	err = q.Dequeue()
	if err != nil {
		log.Printf("3rd deque : %s", err)
	}
	fmt.Println(q)
	err = q.Dequeue()
	if err != nil {
		log.Printf("4th deque : %s", err)
	}
	fmt.Println(q)
	err = q.Dequeue()
	if err != nil {
		log.Printf("5th deque : %s", err)
	}
	fmt.Println(q)
	err = q.Dequeue()
	if err != nil {
		log.Printf("6th deque : %s", err)
	}
	fmt.Println(q)
	err = q.Enqueue(18)
	if err != nil {
		log.Fatalf("error in Enqueu(18) : %s", err)
	}
	fmt.Println(q)
}

func stackExample() {
	s := stack.Stack[int]{Capacity: 3}
	fmt.Printf("s = %+v\n", s)
	err := s.Push(32)
	fmt.Printf("s = %+v\n", s)
	if err != nil {
		log.Fatalf("error in pushing : %s", err)
	}
	err = s.Push(12)
	fmt.Printf("s = %+v\n", s)
	if err != nil {
		log.Fatalf("error in pushing : %s", err)
	}
	err = s.Push(45)
	if err != nil {
		log.Fatalf("error in pushing : %s", err)
	}
	fmt.Printf("s = %+v\n", s)
	fmt.Printf("s is empty = %v\n", s.IsEmpty())
	fmt.Printf("s is full = %v\n", s.IsFull())
	// err = s.Push(12)
	if err != nil {
		log.Fatalf("error in pushing : %s", err)
	}
	err = s.Pop()
	fmt.Printf("s = %+v\n", s)
	if err != nil {
		log.Fatalf("error in pushing : %s", err)
	}
	err = s.Pop()
	fmt.Printf("s = %+v\n", s)
	if err != nil {
		log.Fatalf("error in pushing : %s", err)
	}
	err = s.Pop()
	fmt.Printf("s = %+v\n", s)
	if err != nil {
		log.Fatalf("error in pushing : %s", err)
	}
	err = s.Pop()
	fmt.Printf("s = %+v\n", s)
	if err != nil {
		log.Fatalf("error in pushing : %s", err)
	}
}
