package pets

import "github.com/bhusal-rj/design-pattern/models"

func NewPet(species string) *models.Pet {
	pet := &models.Pet{
		Species:     species,
		Breed:       "",
		MinWeight:   0,
		MaxWeight:   0,
		Description: "no description entered",
		LifeSpan:    0,
	}

	return pet
}
