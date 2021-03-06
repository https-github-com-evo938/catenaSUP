// Copyright (c) 2018 Catena cyber
// Author Philippe Antoine <p.antoine@catenacyber.fr>
// Initializes an empty database for catenaSUP

package main

import (
	"log"
	"os"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Expects name of database file as argument")
	}
	db, err := sql.Open("sqlite3", os.Args[1])
	if err != nil {
		log.Fatalf("failed to open database %v", err)
	}
	defer db.Close()

	//main table is users table
	stmt, err := db.Prepare("CREATE TABLE users (user TEXT UNIQUE, hashpass BLOB, salt BLOB)")
	if err != nil {
		log.Fatalf("failed to prepare create request %v", err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatalf("failed to execute create request %v", err)
	}

	//meta table gives metadata information
	stmt, err = db.Prepare("CREATE TABLE meta (version TEXT, hashfun TEXT)")
	if err != nil {
		log.Fatalf("failed to prepare create request %v", err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatalf("failed to execute create request %v", err)
	}
	stmt, err = db.Prepare("INSERT INTO meta VALUES ('v1', 'argon2')")
	if err != nil {
		log.Fatalf("failed to prepare insert request %v", err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatalf("failed to execute insert request %v", err)
	}
}
