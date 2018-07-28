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

func (prod *Product) InsertProduct() {
    
    db, err := sql.Open("sqlite3", "sttest.sqlt")
    if err != nil {
        panic(err)
    }
    defer db.Close()
    var quanto int = 0;
    if (prod.Quanto) {
        quanto = 1
    }
    result, err := db.Exec("insert into products (name, product_id, category, quanto, creationDate, expirationDate) values ($1, $2, $3, $4, $5, $6)", 
        prod.Name, prod.Product_id, prod.Category, quanto, prod.CreationDate, prod.ExpirationDate);
    if err != nil{
        panic(err)
    }
    prod.Id, _ = result.LastInsertId();

}
