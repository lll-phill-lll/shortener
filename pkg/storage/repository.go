package storage

import (
	"database/sql"
	"fmt"
	"github.com/lll-phill-lll/shortener/pkg/task"
)

var DB = make(map[string]string)

type DataBase interface {
	Save(task.Task) error
	Load(string) (task.Task, error)
	Init() error
}


const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "shortener"
)

type PostgresDB struct {
	DB *sql.DB
}

func (pbd PostgresDB) Init() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	statement := "CREATE TABLE IF NOT EXISTS %s (%s int, %s varchar, %s varchar, %s varchar)"
	query := fmt.Sprintf(statement,
			"user",
			"id",
			"name",
			"email",
			"address",
		)
	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
	pbd.DB = db
	return nil
}

func (pbd PostgresDB) Save(task task.Task) error {
	return nil
}

func (pbd PostgresDB) Load(hash string) (task.Task, error) {
	return task.Task{}, nil

}
