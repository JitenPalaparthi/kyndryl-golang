package interfaces

import "demo/models"

type IPerson interface {
	Add(person *models.Person) (any, error)
	Delete(id string) (any, error)
}
