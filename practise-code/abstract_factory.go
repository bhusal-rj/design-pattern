package main

import "fmt"

// Animal is the type for our abstract factory
type Animal interface {
	Says()
	LikesWater() bool
}

// create a concrete factory for dogs as it returns the object to the dog
type Dog struct{}

func (d *Dog) Says() {
	fmt.Println("Woof!!!")
}

func (d *Dog) LikesWater() bool {
	return true
}

// create a concrete factory for cats
type Cat struct {
}

func (c *Cat) Says() {
	fmt.Println("Meow!!")
}

func (d *Cat) LikesWater() bool {
	return false
}

type AnimalFactory interface {
	New() Animal
}

type DogFactory struct {
}

// we can return both value and pointer in the interface
// This is animal factory which is responsible for returning the animal
func (df *DogFactory) New() Animal {
	return &Dog{}
}

type CatFactory struct {
}

func (cf *CatFactory) New() Animal {
	return &Cat{}
}

func abstractFactory() {
	//Create one each of a DogFactory and a CatFactory

	//call to repository pattern
	//call the adapter
	//this is an instance of DogFactory struct
	dogFactory := DogFactory{}

	//this is cat object
	catFactory := CatFactory{}

	//call the new method to create a dog and a cat

	dog := dogFactory.New()
	cat := catFactory.New()

	dog.Says()
	cat.Says()
}
