package main

import "fmt"

type Repository interface {
	Get(key string) string
	Set(key string, value string)
}

type dataInterface map[string]string

type mysqlRepository struct {
	Data dataInterface
}

type psqlRepository struct {
	Data dataInterface
}

func (mr *mysqlRepository) Get(key string) string {
	val, exists := mr.Data[key]
	if !exists {
		return "-1"
	}
	return val
}
func (mr *mysqlRepository) Set(key string, value string) {
	mr.Data[key] = value
}

func newMysqlRepo(data dataInterface) Repository {
	return &mysqlRepository{
		Data: data,
	}
}
func (mr *psqlRepository) Get(key string) string {
	val, exists := mr.Data[key]
	if !exists {
		return "-1"
	}
	return val
}
func (mr *psqlRepository) Set(key string, value string) {
	mr.Data[key] = value
}

func newPsqlRepo(data dataInterface) Repository {
	return &psqlRepository{
		Data: data,
	}
}

func LearnRepositoryPattern() {
	// Initialize MySQL repository
	mySqlData := make(dataInterface)
	sqlRepo := newMysqlRepo(mySqlData)
	sqlRepo.Set("name", "rajesh")
	fmt.Println("MySQL Repository - Name:", sqlRepo.Get("name"))

	// Initialize PostgreSQL repository
	psqlData := make(dataInterface)
	psqlRepo := newPsqlRepo(psqlData)
	psqlRepo.Set("name", "suresh")
	fmt.Println("PostgreSQL Repository - Name:", psqlRepo.Get("name"))
}
