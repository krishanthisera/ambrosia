package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PsqlConnect struct {
	Host     string
	Port     int64
	User     string
	Password string
	Dbname   string
}

func DbConnect(dbString *PsqlConnect) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbString.Host, dbString.Port, dbString.User, dbString.Password, dbString.Dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db

}

func DbInsert(dbCon *sql.DB, topic Topic) {
	sqlStatement := fmt.Sprintf("INSERT INTO stories (name, description, creation, last_edit, stories) VALUES ('%s', '%s', '%s', '%s', '%s')",
		topic.Name, topic.Description, topic.Creation, topic.LastEdit, topic.Stories)
	_, err := dbCon.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	dbCon.Close()
}
