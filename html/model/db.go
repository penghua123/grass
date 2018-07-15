package model

import (
	"database/sql"
	"fmt"
)

func Login(db *sql.DB, username, password string) bool {
	sStmt := fmt.Sprintf(`SELECT * FROM "User" WHERE "userName"= username AND "password" = password`)
	_, err := db.Exec(sStmt)
	if err != nil {
		return false
	}
	return true
}

func Registor(db *sql.DB, username, password string) error {
	sStmt := fmt.Sprintf(`INSERT INTO "User" ("userName", "password") VALUE ($1, $2)`)
	_, err := db.Exec(sStmt)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func Update(db *sql.DB, username, newPassword string) error {
	sStmt := fmt.Sprintf()

}
