package stack

type Stack struct {
	values []int
}

// NewStack returns a new pointer instance of Stack
func NewStack() *Stack {
	return &Stack{
		values: make([]int, 0),
	}
}

// Push adds an element to the top of the stack
func (s *Stack) Push(n int) {
	s.values = append(s.values, n)
}

// Pop removes an element from the top of the stack
func (s *Stack) Pop() int {
	res := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]

	return res
}

// Len returns the number of stack values
func (s *Stack) Len() int {
	return len(s.values)
}

// IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}
