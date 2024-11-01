package models

import "database/sql"

// Repository is the database repository.
type Repository interface {
	AllDogBreeds() ([]*DogBreed, error)
	GetBreedByName(b string) (*DogBreed, error)
	GetDogOfMonthByID(b int) (*DogOfMonth, error)
}

// wrapper for *sql.DB
type mysqlRepository struct {
	//store the pool of database connection.
	DB *sql.DB
}

// factory method for creating a Repository
func newMysqlRepository(conn *sql.DB) Repository {
	return &mysqlRepository{
		DB: conn,
	}
}

type testRepository struct {
	DB *sql.DB
}

func newTestRepository(conn *sql.DB) Repository {
	return &testRepository{
		DB: conn,
	}
}
