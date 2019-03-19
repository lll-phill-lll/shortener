package storage

import (
	"database/sql"
	"fmt"
	"github.com/lll-phill-lll/shortener/pkg/task"
)

// var DB = make(map[string]string)

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
	DB   *sql.DB
	Name string
}

func (pdb *PostgresDB) createStatementForDBInitialization() string {
	statement := "CREATE TABLE IF NOT EXISTS links (hash varchar PRIMARY KEY, link varchar)"
	return statement
}

func (pdb *PostgresDB) Init() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	err = db.Ping()
	pdb.DB = db

	_, err = db.Exec(pdb.createStatementForDBInitialization())
	if err != nil {
		return err
	}
	pdb.DB = db
	return nil
}

func (pdb *PostgresDB) Save(task task.Task) error {
	_, err := pdb.DB.Exec("INSERT INTO links(hash, link) VALUES($1, $2)", task.Hash, task.URL)
	if err != nil {
		return err
	}

	return nil
}

func (pdb *PostgresDB) Load(hash string) (task.Task, error) {
	var URL string
	rows, err := pdb.DB.Query("SELECT link FROM links WHERE hash = $1", hash)
	defer rows.Close()
	if err != nil {
		return task.Task{}, err
	}
	for rows.Next() {
		err = rows.Scan(&URL)
		if err != nil {
			return task.Task{}, err
		}
		break
	}
	return task.Task{Hash: hash, URL: URL}, nil
}
