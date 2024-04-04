package database

import (
	"demo/models"

	"gorm.io/gorm"
)

type Person struct {
	DB *gorm.DB
}

func (p *Person) Add(person *models.Person) (any, error) {
	p.DB.AutoMigrate(&models.Person{}) // This code going to create a table if table does not exist
	tx := p.DB.Save(person)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return person.ID, nil
}

func (p *Person) Delete(id string) (any, error) {
	tx := p.DB.Delete(&models.Person{}, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx.RowsAffected, nil
}

func (p *Person) Get(id string) (*models.Person, error) {
	person := new(models.Person)
	tx := p.DB.First(person, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return person, nil
}
