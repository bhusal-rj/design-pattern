package main

import "fmt"

func main() {

	LearnRepositoryPattern()

	//learning singleton pattern
	_ = LearnSingletonPattern("Rajesh Bhusal")
	_ = LearnSingletonPattern("Bhusal Rajesh")
	val3 := LearnSingletonPattern("This is demo singleton pattern")
	fmt.Println(val3)
}
