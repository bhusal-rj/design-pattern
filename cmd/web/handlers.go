package main

import (
	"fmt"
	"net/http"

	"github.com/bhusal-rj/design-pattern/pets"
	"github.com/go-chi/chi/v5"
	"github.com/tsawler/toolbox"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.gohtml", nil)
}

func (app *application) ShowPage(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	app.render(w, fmt.Sprintf("%s.page.gohtml", page), nil)
}

func (app *application) CreateDogFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("dog"))
}

func (app *application) CreateCatFromFactory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creting cat from factory")
	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("cat"))
}

func (app *application) CreateDogFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dog, err := pets.NewPetFromAbstractFactory("dog")
	if err != nil {
		t.ErrorJSON(w, err)
		return
	}
	_ = t.WriteJSON(w, http.StatusAccepted, dog)
}

func (app *application) CreateCatFromAbstractFactory(w http.ResponseWriter, r *http.Request) {

	var t toolbox.Tools
	cat, err := pets.NewPetFromAbstractFactory("cat")
	if err != nil {
		t.ErrorJSON(w, err)
		return
	}
	_ = t.WriteJSON(w, http.StatusAccepted, cat)

}
func (app *application) TestPatterns(w http.ResponseWriter, r *http.Request) {
	app.render(w, "test.page.gohtml", nil)
}

func (app *application) GetAllDogBreedsJSON(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dogBreeds, err := app.App.Models.DogBreed.All()

	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = t.WriteJSON(w, http.StatusOK, dogBreeds)

}

func (app *application) CreateDogWithBuilder(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	//create the dog with builder pattern
	newPet, err := pets.NewPetBuilder().
		SetSpecies("dog").
		SetBreed("mixed breed").
		SetWeight(15).
		SetDescription("Mixed breed").
		SetColor("Black").
		SetAgeEstimated(20).
		Build()

	if err != nil {
		t.ErrorJSON(w, err)
		return
	}

	t.WriteJSON(w, http.StatusOK, newPet)

}

func (app *application) CreateCatWithBuilder(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	newPet, err := pets.NewPetBuilder().
		SetSpecies("cat").
		SetBreed("mixed breed").
		SetWeight(15).
		SetDescription("Mixed cat breed").
		SetColor("White").
		SetAgeEstimated(20).
		Build()

	if err != nil {
		t.ErrorJSON(w, err)
		return
	}

	t.WriteJSON(w, http.StatusOK, newPet)

}

// func (app *application) AnimalFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
// 	//Setup the toolbox
// 	var t toolbox.Tools
// 	// Get species from URL itself.

// 	// Create a pet from abstract factory.

// 	// Write the result as JSON
// }
