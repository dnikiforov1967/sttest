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
    id int64
    parent_id int64
    EventType string `json:"type"`
	Terminal bool `json:"terminal"`
	ExecutionStruct `json:"execution"`
	CashDirection `json:"cashDirection"`
}

type ExecutionStruct struct {
	OnStruct `json:"on"`
	Origin string `json:"origin"`
	ExecType string `json:"type"`
}

type OnStruct struct {
	Kind string `json:"kind"`
}

type CashDirection struct {
	Path string `json:"path"`
	CashType string `json:"type"`
	Payment `json:"payment"`
}

type Payment struct {
	PaymentType string `json:"type"`
	Method string `json:"method"`
	Algorithm string `json:"algorithm"`
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

func (prod *Product) findIdByProductId(tx *sql.Tx) error {
    row := tx.QueryRow("select id from products where product_id = $1",
    prod.Product_id);
    err := row.Scan(&prod.id)
    if err != nil {
        if err != sql.ErrNoRows {
            return err
        } else {
            return errhand.ErrProdNotFound
        }
    }
    return nil
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
    rows, err := db.Query("select eventType, terminal, kind, origin, execType, path, cashType, paymentType, method, algorithm from events where parent_id = $1", prod.id);
    if err != nil {
        return err
    }
    defer rows.Close()
    prod.Terms = TermsStruct{}
    for rows.Next(){
        event := Event{}
        err := rows.Scan(&event.EventType, &event.Terminal, &event.Kind, &event.Origin,
				&event.ExecType, &event.Path, &event.CashType, &event.PaymentType, &event.Method, &event.Algorithm);
        if err != nil {
            return err
        }
        prod.Terms.Events = append(prod.Terms.Events, event)
    }
    return nil
}

func (event *Event) InsertEvent(tx *sql.Tx) error {
    result, err := tx.Exec("insert into events (parent_id, eventType, terminal, kind, origin, execType, path, cashType, paymentType, method, algorithm) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
        event.parent_id, event.EventType, event.Terminal, event.Kind, event.Origin, event.ExecType,
				 event.Path, event.CashType, event.PaymentType, event.Method, event.Algorithm);
    if err != nil{
        tx.Rollback()
        return err
    }
    event.id, _ = result.LastInsertId();
    return nil
}

func (prod *Product) InsertEvents(tx *sql.Tx) error {
    //Cycle to insert referenced data
    for _, event := range prod.Terms.Events {
        event.parent_id = prod.id
        err := event.InsertEvent(tx)
        if err != nil{
            tx.Rollback()
            return err
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
    
    //Cycle to insert referenced data
    err = prod.InsertEvents(tx)
    if err != nil{
        return err
    }

    err = tx.Commit()
    if err != nil{
        return err
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

    result, err := tx.Exec("update products set name=$1, product_id=$2, category=$3, quanto=$4 where product_id=$5", 
        prod.Name, prod.Product_id, prod.Category, prod.Quanto, origId);
    if err != nil{
        tx.Rollback()
        return err
    }
    count, err := result.RowsAffected()
    if err != nil{
        tx.Rollback()
        return err
    }
    if count == 0 {
        tx.Rollback()
        return errhand.ErrProdNotFound
    }

    //Find productId
    err = prod.findIdByProductId(tx)
    if err != nil {
        tx.Rollback()
        return err
    }

    //Remove old events
    _, err = tx.Exec("delete from events where parent_id=$1", prod.id);
    if err != nil{
        tx.Rollback()
        return err
    }
    //Cycle to insert referenced data
    err = prod.InsertEvents(tx)
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
    result, err := tx.Exec("delete from products where product_id = $1", 
        prod.Product_id);
    if err != nil{
        tx.Rollback()
        return err
    }
    count, err := result.RowsAffected()
    if err != nil{
        tx.Rollback()
        return err
    }
    if count == 0 {
        tx.Rollback()
        return errhand.ErrProdNotFound
    }
    err = tx.Commit()
    if err != nil{
        return err
    }
    return nil
}
