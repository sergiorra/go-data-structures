package queue

type Queue struct {
	values []int
}

// New returns a new pointer instance of Queue
func New() *Queue {
	return &Queue{
		values: make([]int, 0),
	}
}

// Push adds an element to the top of the queue
func (q *Queue) Push(v int) {
	q.values = append(q.values, v)
}

// Pop returns an element from the top of the queue and removes it from the queue
func (q *Queue) Pop() int {
	val := q.values[0]
	q.values = q.values[1:]

	return val
}

// Len returns the number of queue values
func (q *Queue) Len() int {
	return len(q.values)
}

// IsEmpty returns true if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}
