package store

type Store interface {
	Request() RequestRepository
}
