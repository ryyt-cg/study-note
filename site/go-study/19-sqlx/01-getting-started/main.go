package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = db.Exec(`CREATE TABLE place (id integer primary key AUTOINCREMENT, name text NOT NULL, place text)`)
	if err != nil {
		log.Fatalln(err)
	}

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO place (name, place) VALUES (?, ?)", "bobby", "maroubra")
	tx.MustExec("INSERT INTO place (name, place) VALUES (?, ?)", "sally", "maroubra")
	tx.MustExec("INSERT INTO place (name, place) VALUES (?, ?)", "bobby", "bondi")
	tx.MustExec("INSERT INTO place (name, place) VALUES (?, ?)", "sally", "newtown")
	tx.MustExec("INSERT INTO place (name, place) VALUES (?, ?)", "bobby", "redfern")
	tx.MustExec("INSERT INTO place (name, place) VALUES (?, ?)", "sally", "redfern")
	tx.Commit()

	rows, err := db.Queryx("SELECT * FROM place")
	if err != nil {
		log.Fatalln(err)
	}

	// NullString is used to allow for NULL values into the table columns where the column are not defined as NOT NULL.
	// This is useful when the column is not required to have a value.
	// There are other types like NullInt64, NullFloat64, NullBool, NullTime, etc.
	for rows.Next() {
		var p struct {
			ID    int            `db:"id"`
			Name  sql.NullString `db:"name"`
			Place string         `db:"place"`
		}
		err = rows.StructScan(&p)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", p)
	}
}
