package queue

type QueueChan[T any] struct {
	ch chan T
}

func New[T any](size int) *QueueChan[T] {
	return &QueueChan[T]{
		ch: make(chan T, size),
	}
}

func (q *QueueChan[T]) Push(value T) error {
	q.ch <- value
	return nil
}

func (q *QueueChan[T]) Pop() <-chan T {
	return q.ch
}

func (q *QueueChan[T]) Close() {
	close(q.ch)
}