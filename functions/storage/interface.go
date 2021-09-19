package main

type Storage interface {
	Put(path, file, data string) error
	List(path string)
	Get(path, file string) ([]byte, error)
	Delete(path, file string) error
	SetPath(path string) error
	SetData(data string)
	SetFile(file string) error
	GetPath() string
	GetData() string
	GetFile() string
}
