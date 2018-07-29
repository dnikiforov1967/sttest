package dbfunc

import (
    "../errhand"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type Product struct {
    id int64
    Name string `json:"name,omitempty"`
    Product_id string `json:"product_id,omitempty"`
    Category string `json:"category,omitempty"`
    Quanto bool `json:"quanto,omitempty"`
    CreationDate string `json:"creationDate,omitempty"`
    ExpirationDate string `json:"expirationDate,omitempty"`
}

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

func (prod *Product) FetchProductByProductId() (error) {
    db, err := openLocalDb();
    if err != nil {
        return err
    }
    defer db.Close()
    row := db.QueryRow("select id, name, product_id, category, quanto, creationDate, expirationDate from products where product_id = $1",
    prod.Product_id);
    err = row.Scan(&prod.id, &prod.Name, &prod.Product_id, &prod.Category, &prod.Quanto, &prod.CreationDate, &prod.ExpirationDate)
    if err != nil {
        if err != sql.ErrNoRows {
            return err
        } else {
            return errhand.ErrProdNotFound
        }
    } 
    return nil
}

func (prod *Product) InsertProduct() error {
    
    db, err := openLocalDb();
    if err != nil {
        return err
    }

    defer db.Close()
    tx, err := openTrans(db)
    if err != nil {
        return err
    }

    result, err := tx.Exec("insert into products (name, product_id, category, quanto, creationDate, expirationDate) values ($1, $2, $3, $4, $5, $6)", 
        prod.Name, prod.Product_id, prod.Category, prod.Quanto, prod.CreationDate, prod.ExpirationDate);
    if err != nil{
        tx.Rollback()
        return err
    }
    prod.id, _ = result.LastInsertId();
    err = tx.Commit()
    if err != nil{
        panic(err)
    }
    return nil
}

func (prod *Product) UpdateProduct() {
    
    db, _ := openLocalDb();

    defer db.Close()
    tx, _ := openTrans(db)

    _, err := tx.Exec("update products set name=$1, product_id=$2, category=$3, quanto=$4 where id=$5", 
        prod.Name, prod.Product_id, prod.Category, prod.Quanto, prod.id);
    if err != nil{
        tx.Rollback()
        panic(err)
    }

    err = tx.Commit()
    if err != nil{
        panic(err)
    }
}

func (prod *Product) DeleteProductByProductId() {
    
    db, _ := openLocalDb()

    defer db.Close()

    tx, _ := openTrans(db)

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
    
    db, _ := openLocalDb()

    defer db.Close()

    tx, _ := openTrans(db)

    _, err := tx.Exec("delete from products where id = $1", 
        prod.id);
    if err != nil{
        tx.Rollback()
        panic(err)
    }
    err = tx.Commit()
    if err != nil{
        panic(err)
    }
}