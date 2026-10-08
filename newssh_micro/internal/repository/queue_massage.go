package repository

type QueueMassage[T any] interface {
	Push(T) error
	Pop() <-chan T
	Close()
}