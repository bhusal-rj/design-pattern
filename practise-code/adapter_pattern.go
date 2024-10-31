package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type ToDo struct {
	UserID    int    `json:"userId" xml:"userId"`
	ID        int    `json:"id" xml:"id"`
	Title     string `json:"title" xml:"title"`
	Completed bool   `json:"completed" xml:"completed"`
}

type DataInterface interface {
	GetData() (*ToDo, error)
}
type RemoteService struct {
	Remote DataInterface
}

func (rs *RemoteService) CallRemoteService() (*ToDo, error) {
	return rs.Remote.GetData()
}

// JSONBackend satisfies the DataInteface
// Remote in RemoteService is eligible to be DataInterface
type JSONBackend struct{}

func (jb *JSONBackend) GetData() (*ToDo, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	todoData := &ToDo{}
	err = json.Unmarshal(body, todoData)
	if err != nil {
		fmt.Println(err, "Error")
		return nil, err
	}
	return todoData, nil

}

type XMLBackend struct{}

func (xb *XMLBackend) GetData() (*ToDo, error) {
	xmlFile := `
	<?xml version="1.0" encoding="UTF-8" ?>
	<root>
	<userID>1</userID>
	<title>delectus aut autem</title>
	<id>1</id>
	<completed>false</completed>
	</root>
	`

	var todo ToDo
	_ = xml.Unmarshal([]byte(xmlFile), &todo)
	return &todo, nil
}

func adapter() {
	//No adapter
	todo := getRemoteData()

	if todo == nil {
		fmt.Println("There has been an error")
		return
	}
	fmt.Println("TODO without adapter", todo.ID, todo.UserID, todo.Title)

	//With adapter,using JSON
	jsonBackend := &JSONBackend{}
	jsonAdapter := &RemoteService{Remote: jsonBackend}
	tdFromJSON, _ := jsonAdapter.CallRemoteService()
	fmt.Println("FROM JSON Adapter", tdFromJSON.Title)

	//With adapter,using XML
	xmlBackend := &XMLBackend{}
	xmlAdapter := &RemoteService{Remote: xmlBackend}
	tdFromXML, _ := xmlAdapter.CallRemoteService()
	fmt.Println("From XML Adapter", tdFromXML.Title)
}

func getRemoteData() *ToDo {
	// Get the request header capture if the response has any error
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	todoData := &ToDo{}
	err = json.Unmarshal(body, todoData)
	if err != nil {
		fmt.Println(err, "Error")
		return nil
	}
	return todoData

}
