package common

type Queue[T any] struct {
	q []T
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		q: []T{},
	}
}

func (q *Queue[T]) Pop() (T, bool) {
	var out T
	if len(q.q) == 0 {
		return out, false
	}

	out = q.q[0]

	if len(q.q) == 1 {
		q.q = []T{}
	} else {
		q.q = q.q[1:]
	}
	return out, true
}

func (q *Queue[T]) Push(e T) {
	q.q = append(q.q, e)
}
