package stack

type Stack struct {
	values []int
}

// New returns a new pointer instance of Stack
func New() *Stack {
	return &Stack{
		values: make([]int, 0),
	}
}

// Push adds an element to the top of the stack
func (s *Stack) Push(v int) {
	s.values = append(s.values, v)
}

// Pop returns an element from the top of the stack and removes it from the stack
func (s *Stack) Pop() int {
	val := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]

	return val
}

// Len returns the number of stack values
func (s *Stack) Len() int {
	return len(s.values)
}

// IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}
