package integration_test

import (
	"demo/models"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"testing"
)

func Up() {

	// 1. Ensure application is up and running
	// 2. Ensure all related components are up and running

	/*
		docker run --name pg -d -p 5432:5432 -e POSTGRES_USER=app -e POSTGRES_PASSWORD=app@123 -e POSTGRES_DB=demodb postgres
	*/
	// also need to check whether docker is installed or not
	cmd := exec.Command("docker", "run", "--name", "pg", "-d", "-p", "5432:5432", "-e", "POSTGRES_USER=app", "-e", "POSTGRES_PASSWORD=app@123", "-e", "POSTGRES_DB=demodb", "postgres")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	// cmd1 := exec.Command("./app")
	// err = cmd1.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func Down() {
	cmd := exec.Command("docker", "rm", "-f", "pg")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func TestSuccess(t *testing.T) {
	//Up()
	//defer Down()

	cmd := exec.Command("curl", "--location", "localhost:8086/person/add", "--header", `Content-Type: application/json`, `--data-raw`, `{
		"name":"Jiten P",
		"email":"Jiten.Palaparthi@gmail.com",
		"status":"active"
	}`)

	output, err := cmd.Output()
	fmt.Println("Record Inserted:", string(output))

	if err != nil {
		t.Fail()
	}

	cmd1 := exec.Command("curl", "--location", `localhost:8086/person/`+string(output))

	output2, err := cmd1.Output()
	if err != nil {
		t.Fail()
	}
	person := models.Person{}
	err = json.Unmarshal(output2, &person)
	if err != nil {
		t.Fail()
	}
	fmt.Println(person)
	if person.Email != "Jiten.Palaparthi@gmail.com" || person.Name != "Jiten P" {
		t.Fail()
	}

	// ensure that the application is up and running
	// make a call to the application
	//

}
