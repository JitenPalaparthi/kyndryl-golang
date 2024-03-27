package person

import (
	"errors"
	"fmt"
	"os"
)

type Person struct {
	Name, Email string
	Age         uint8
}

func (p Person) Add(file string) (bool, error) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return false, err // Returning error from this
	}
	data := fmt.Sprintf(`{"Name":"%s","Age":%d,"Email:"%s"}`, p.Name, p.Age, p.Email)
	//data = fmt.Sprintf("{""Name"\":""%s"",""Age"":%d,""Email:""%s""}""", p.Name, p.Age, p.Email)

	n, err := f.WriteString(data + "\n")
	if err != nil {
		return false, err // Returning error from this
	}
	if n > 0 {
		return true, nil
	}
	return false, errors.New("no data has been added to the file")
	//return false, fmt.Errorf("no data has been added to the file")
}
