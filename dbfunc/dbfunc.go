package dbfunc

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type Product struct {
    Id int64
    Name string
    Product_id string
    Category string
    Quanto bool
    CreationDate string
    ExpirationDate string
}

func openLocalDb() *sql.DB {
    db, err := sql.Open("sqlite3", "sttest.sqlt")
    if err != nil {
        panic(err)
    }
    return db;
}

func openTrans(db *sql.DB) *sql.Tx {
    tx, err := db.Begin()
    if err != nil {
        panic(err)
    }
    return tx;
}

func (prod *Product) InsertProduct() {
    
    db := openLocalDb();

    defer db.Close()
    var quanto int = 0;
    if prod.Quanto {
        quanto = 1
    }
    tx := openTrans(db)

    result, err := tx.Exec("insert into products (name, product_id, category, quanto, creationDate, expirationDate) values ($1, $2, $3, $4, $5, $6)", 
        prod.Name, prod.Product_id, prod.Category, quanto, prod.CreationDate, prod.ExpirationDate);
    if err != nil{
        tx.Rollback()
        panic(err)
    }
    prod.Id, _ = result.LastInsertId();
    err = tx.Commit()
    if err != nil{
        panic(err)
    }
}

func (prod *Product) DeleteProductByProductId() {
    
    db := openLocalDb()

    defer db.Close()

    tx := openTrans(db)

    _, err := tx.Exec("delete from products where product_id = $1", 
        prod.Product_id);
    if err != nil{
        tx.Rollback()
        panic(err)
    }
    err = tx.Commit()
    if err != nil{
        panic(err)
    }
}

func (prod *Product) DeleteProduct() {
    
    db := openLocalDb()

    defer db.Close()

    tx := openTrans(db)

    _, err := tx.Exec("delete from products where id = $1", 
        prod.Id);
    if err != nil{
        tx.Rollback()
        panic(err)
    }
    err = tx.Commit()
    if err != nil{
        panic(err)
    }
}