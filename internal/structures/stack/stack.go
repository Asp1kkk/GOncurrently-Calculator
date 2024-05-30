package stack

type node[T any] struct {
	Value    T
	Previous *node[T]
}

type Stack[T any] struct {
	Top *node[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		Top: nil,
	}
}

func (s *Stack[T]) Push(value T) {
	new := &node[T]{
		Value:    value,
		Previous: s.Top,
	}

	s.Top = new
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.Top == nil {
		var nul T
		return nul, false
	}

	value := s.Top.Value
	s.Top = s.Top.Previous

	return value, true
}
