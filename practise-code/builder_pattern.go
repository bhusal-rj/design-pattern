package main

import "fmt"

type Address struct {
	street  string
	number  int32
	city    string
	country string
}

func CreateAddress() *Address {
	return &Address{}
}

func (a *Address) SetStreet(streetName string) *Address {
	a.street = streetName
	return a
}

func (a *Address) SetNumber(number int32) *Address {
	a.number = number
	return a
}

func (a *Address) SetCity(city string) *Address {
	a.city = city
	return a
}

func (a *Address) SetCountry(country string) *Address {
	a.country = country
	return a
}

func builderPattern() {
	// Method chaining
	myAddress := CreateAddress().
		SetCountry("Nepal").
		SetCity("Kathmandu").
		SetNumber(5).SetStreet("Pulchowk")
	fmt.Println(myAddress)
}
