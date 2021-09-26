package main

type Storage interface {
	Put() error
	List()
	Get() ([]byte, error)
	Delete() error
}
