package database

import (
	"fmt"
	"io/ioutil"
	"strings"

	log "github.com/sirupsen/logrus"
)

// MakeMigrations executes the sql migration files
func MakeMigrations() {
	file, err := ioutil.ReadFile("database/migrations.sql")

	if err != nil {
		log.Error("Error while opening migrations file: ", err)
	}

	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		fmt.Println("Pute")
		_, err := DB.Exec(request)
		if err != nil {
			log.Error("Error during migrations: ", err)
		}
	}

}
