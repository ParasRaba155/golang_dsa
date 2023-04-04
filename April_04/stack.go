package stack

import (
	"errors"
)

var (
	ErrStackCapicityFull = errors.New("capacity of stack reached")
	ErrStackEmpty        = errors.New("stack is empty")
)

// Stack is the custom data type that allows us to `Push`, `Pop`,`Peek`
type Stack[C any] struct {
	Elements []C // Elements holds the actual element in the Stack
	Capacity int // Capacity of the stack
}

// IsEmpty checks if the stack is empty or not
func (s *Stack[C]) IsEmpty() bool {
	return len(s.Elements) == 0
}

// IsFull checks if the stack is full or not
func (s *Stack[C]) IsFull() bool {
	return len(s.Elements) >= s.Capacity
}

// Push pushes the element to the last in stack
func (s *Stack[C]) Push(e C) error {
	if s.IsFull() {
		return ErrStackCapicityFull
	}
	s.Elements = append(s.Elements, e)
	return nil
}

// Peek returns the last element in the stack
func (s *Stack[C]) Peek() C {
	if s.IsEmpty() {
        var c C
        return c // return the zero value of the generic
	}
    return s.Elements[len(s.Elements)-1]
}

// Pop pops the last element from stack if it exist otherwise throws the error
func (s *Stack[C]) Pop() error {
	if s.IsEmpty() {
		return ErrStackEmpty
	}
	s.Elements = s.Elements[0 : len(s.Elements)-1]
	return nil
}
