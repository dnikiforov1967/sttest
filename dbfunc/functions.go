package dbfunc

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func openLocalDb() (*sql.DB, error) {
    db, err := sql.Open("sqlite3", "sttest.sqlt")
    if err != nil {
        return nil, err
    }
    return db, nil;
}

func openTrans(db *sql.DB) (*sql.Tx, error) {
    tx, err := db.Begin()
    if err != nil {
        return nil, err
    }
    return tx, nil;
}
