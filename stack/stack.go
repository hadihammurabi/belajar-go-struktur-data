package stack

type Stack struct {
	size uint
	data []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Size() uint {
	return s.size
}

func (s *Stack) Push(d int) {
	s.data = append(s.data, d)
	s.size++
}

func (s *Stack) Pop() int {
	result := s.data[s.size-1]

	s.data = s.data[:s.size-1]
	s.size--

	return result
}
