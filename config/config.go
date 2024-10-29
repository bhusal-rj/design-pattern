package config

import (
	"database/sql"
	"sync"

	"github.com/bhusal-rj/design-pattern/models"
)

type Application struct {
	Models *models.Models
}

var instance *Application
var once sync.Once
var db *sql.DB

func New(pool *sql.DB) *Application {
	db = pool
	return GetInstance()
}

func GetInstance() *Application {
	// if and only do is called only once
	// all the instance of the application services will use this method to get the instance.
	// if no instance is pre created new instance will be created else existing instance will be returned
	once.Do(func() {
		instance = &Application{
			Models: models.New(db),
		}
	})
	return instance
}
