package models

import "encoding/json"

type Person struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	LastUpdated int64  `json:"lastUpdated"`
}

func (p *Person) Validate() bool {
	if p.Name == "" || p.Email == "" || p.ID == 0 {
		return false
	}
	return true
}

func (p *Person) ToString() (string, error) {
	bytes, err := json.Marshal(p)
	return string(bytes), err
}

func (p *Person) ToByte() ([]byte, error) {
	return json.Marshal(p)
}
