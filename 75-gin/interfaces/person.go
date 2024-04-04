package interfaces

import "demo/models"

type IPerson interface {
	Add(person *models.Person) (any, error)
	Get(id string) (*models.Person, error)
	Delete(id string) (any, error)
}
