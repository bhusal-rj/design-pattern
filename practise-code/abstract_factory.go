package main

import "fmt"

// Animal is the type for our abstract factory
type Animal interface {
	Says()
	LikesWater() bool
}

// create a concrete factory for dogs
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
func (df *DogFactory) New() Animal {
	return &Dog{}
}

type CatFactory struct {
}

func (cf *CatFactory) New() Animal {
	return &Cat{}
}

func main() {
	//Create one each of a DogFactory and a CatFactory
	dogFactory := DogFactory{}
	catFactory := CatFactory{}

	//call the new method to create a dog and a cat
	dog := dogFactory.New()
	cat := catFactory.New()

	dog.Says()
	cat.Says()
}
