package models_test

import (
	"demo/models"
	"testing"
)

func TestValidateSuccess(t *testing.T) {
	p1 := models.Person{}
	p1.Email = "Jitenp@Gmail.com"
	p1.ID = 1
	p1.Name = "Jiten"
	if !p1.Validate() {
		t.Fail()
	}
}

func TestValidateFail(t *testing.T) {
	p1 := models.Person{}
	p1.ID = 1
	p1.Name = "Jiten"
	//	time.Sleep(time.Second * 11)
	if p1.Validate() {
		t.Fail()
	}
}

func TestToStringSuccess(t *testing.T) {
	p1 := models.Person{}
	p1.Email = "Jiten.Palaparthi@gmail.com"
	p1.ID = 101
	p1.Name = "Jiten"
	p1.Status = "active"
	p1.LastUpdated = 1711980290 //time.Now().Unix()
	str, err := p1.ToString()
	if err != nil {
		t.Fail()
	}
	if str != `{"id":101,"name":"Jiten","email":"Jiten.Palaparthi@gmail.com","status":"active","lastUpdated":1711980290}` {
		t.Fail()
	}

}

func TestToStringFail(t *testing.T) {
	p1 := models.Person{}
	p1.Email = "Jiten.Palaparthi@gmail.com"
	p1.ID = 101
	p1.Name = "Jiten"
	p1.Status = "active"
	p1.LastUpdated = 1711980290 //time.Now().Unix()
	str, err := p1.ToString()
	if err != nil {
		t.Fail()
	}
	if str == `{"id":101,"name":"Jiten","email":"Jiten.Palaparthi@gmail.com","status":"active","lastUpdated":"1711980290"}` {
		t.Fail()
	}

}
