package driver

import (
	"database/sql"
	"fmt"
	"time"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 5 // might be changed in production
const maxIdleDbConn = 5 // might be changed in production
const maxDbLifeTime = 5 * time.Minute

func ConnectDb(dsn string) (*DB, error) {
	d, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifeTime)

	err = testDB(d)
	if err != nil {
		return nil, err
	}

	dbConn.SQL = d
	return dbConn, nil
}

func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		fmt.Println("Error!", err)
		return err
	}
	fmt.Println("*** Pinged database successfully! ***")

	return nil
}
