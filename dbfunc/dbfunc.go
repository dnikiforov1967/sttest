package dbfunc

import (
    "../errhand"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type Product struct {
    id int64
    Name string `json:"name"`
    Product_id string `json:"product_id"`
    Category string `json:"category"`
    Quanto bool `json:"quanto"`
    CreationDate string `json:"creationDate"`
    ExpirationDate string `json:"expirationDate"`
    Terms TermsStruct `json:"terms"`
}

type TermsStruct struct {
    Events []Event `json:"events"`
}

type Event struct {
    EventType string `json:"type"`
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
    rows, err := db.Query("select eventType from events where parent_id = $1", prod.id);
    if err != nil {
        return err
    }
    defer rows.Close()
    prod.Terms = TermsStruct{}
    for rows.Next(){
        event := Event{}
        err := rows.Scan(&event.EventType)
        if err != nil {
            return err
        }
        prod.Terms.Events = append(prod.Terms.Events, event)
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
    
    //Cycle to insert referenced data
    for _, event := range prod.Terms.Events {
        _, err := tx.Exec("insert into events (parent_id, eventType) values ($1, $2)",
        prod.id, event.EventType);
        if err != nil{
            tx.Rollback()
            return err
        }
    }

    err = tx.Commit()
    if err != nil{
        panic(err)
    }
    return nil
}

func (prod *Product) UpdateProduct(origId string) error {
    
    db, err := openLocalDb();
    if err != nil {
        return err
    }

    defer db.Close()
    tx, err := openTrans(db)
    if err != nil {
        return err
    }

    _, err = tx.Exec("update products set name=$1, product_id=$2, category=$3, quanto=$4 where product_id=$5", 
        prod.Name, prod.Product_id, prod.Category, prod.Quanto, origId);
    if err != nil{
        tx.Rollback()
        return err
    }

    err = tx.Commit()
    if err != nil{
        return err
    }
    return nil
}

func (prod *Product) DeleteProductByProductId() error {
    
    db, err := openLocalDb()
    if err != nil{
        return err
    }

    defer db.Close()

    tx, err := openTrans(db)
    if err != nil{
        return err
    }

    _, err = tx.Exec("delete from events where parent_id in (select id from products where product_id = $1)", 
        prod.Product_id);
    if err != nil{
        tx.Rollback()
        return err
    }
    _, err = tx.Exec("delete from products where product_id = $1", 
        prod.Product_id);
    if err != nil{
        tx.Rollback()
        return err
    }
    err = tx.Commit()
    if err != nil{
        return err
    }
    return nil
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