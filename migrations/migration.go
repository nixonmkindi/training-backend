package migrations

import (
	"fmt"
	"strings"

	"github.com/tzdit/sample_api/package/log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Up() {
	m, err := migrate.New(
		"file://migrations",
		"postgres://postgres:postgres@localhost:5432/aim?sslmode=disable")
	if err != nil {
		log.Errorf("error creating new migration: %v", err)
	}
	if err := m.Up(); err != nil {
		if strings.Contains(err.Error(), "no change") {
			fmt.Println("You have got the latest version of the database. Nothing to migrate....")
			return
		}
		log.Errorf("error during up migration:%v", err)
	}
	fmt.Println("Successfully completed database up migration")
}
func Down() {
	m, err := migrate.New(
		"file://migrations",
		"postgres://postgres:postgres@localhost:5432/aim?sslmode=disable")
	if err != nil {
		log.Errorf("error creating new migration: %v", err)
	}
	if err := m.Down(); err != nil {
		if strings.Contains(err.Error(), "no change") {
			fmt.Println("You have got the latest version of the database. Nothing to migrate....")
			return
		}
		log.Errorf("error during down migration:%v", err)
	}
	fmt.Println("Successfully completed database down migration")

}
