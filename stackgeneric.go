package generics

type StackErr string

func (e StackErr) Error() string {
	return string(e)
}

const (
	EmptyStackErr = StackErr("Empty Stack")
)

type Stack[T any] struct {
	elements []T
}

func (sp *Stack[T]) Push(e T) {
	sp.elements = append(sp.elements, e)
}

func (sp *Stack[T]) Pop() (T, error) {
	if len(sp.elements) == 0 {
		var zero T
		return zero, EmptyStackErr
	}
	e := sp.elements[len(sp.elements) - 1]
	sp.elements = sp.elements[:len(sp.elements) - 1]
	return e, nil
}

func (sp *Stack[T]) IsEmpty() bool {
	return len(sp.elements) == 0
}

func (sp *Stack[T]) GetElements() []T {
	return sp.elements
}
