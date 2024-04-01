package helpers

import "os"

func SaveToFile(filename string, data string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte(data + "\n"))
	if err != nil {
		return err
	}
	return nil
}
