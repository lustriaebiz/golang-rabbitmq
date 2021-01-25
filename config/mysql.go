package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const username, password, database string = "root", "", "bni-autodebet"

var con = fmt.Sprintf("%v:%v@/%v", username, password, database)

func MySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", con)

	if err != nil {
		return nil, err
	}

	return db, nil
}
