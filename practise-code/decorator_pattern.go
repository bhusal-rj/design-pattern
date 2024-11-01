package main

import "fmt"

type DB interface {
	Store(string) error
}

type Store struct{}

func (s *Store) Store(value string) error {
	fmt.Println("Storing into db", value)
	return nil
}

func myExecuteFunc(db DB) ExecuteFn {
	return func(s string) {
		fmt.Println("My Execute function")
		db.Store(s)
	}
	//access to DB
}
func decorator() {
	s := &Store{}
	Execute(myExecuteFunc(s))
}

type ExecuteFn func(string)

func Execute(fn ExecuteFn) {
	fn("FOO")
}
