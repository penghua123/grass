package main

import (
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "huap"
    password = ""
    dbname   = "test"
)

type Teacher struct {
    ID  int
    Age int
    Name string
    Site string
}

func main() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        panic(err)
    }

    fmt.Println("Successfully connected!")

    sqlStatement := `INSERT INTO teacher (id, age, name, site)  VALUES ($1, $2, $3, $4)  RETURNING id`
    id := 1
    err = db.QueryRow(sqlStatement, 3, 46, "Tom", "Shanghai").Scan(&id)

    //if err != nil {
    //    panic(err)
    //}
    fmt.Println("New record ID is:", id)
}
