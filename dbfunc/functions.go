package dbfunc

import (
	"github.com/dnikiforov1967/sttest/config"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

//Function opens db connection
func openLocalDb() (*sql.DB, error) {
    db, err := sql.Open("sqlite3", config.Database)
    if err != nil {
        return nil, err
    }
    return db, nil;
}

//Function begins new transaction
func openTrans(db *sql.DB) (*sql.Tx, error) {
    tx, err := db.Begin()
    if err != nil {
        return nil, err
    }
    return tx, nil;
}
