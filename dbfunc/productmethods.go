package dbfunc

import (
	"time"
    "github.com/dnikiforov1967/sttest/errhand"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

//Method searches Product record in database which has specified product_id value
//and returns its database-specific primary key value
func (prod *Product) findIdByProductId(tx *sql.Tx) error {
    row := tx.QueryRow("select id from products where product_id = $1",
    prod.Product_id);
    err := row.Scan(&prod.id)
    if err != nil {
        if err != sql.ErrNoRows {
            return err
        } else {
            return errhand.ProdNotFound
        }
    }
    return nil
}

//Method searches Product record in database which has specified product_id value
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
            return errhand.ProdNotFound
        }
    }
    rows, err := db.Query("select eventType, terminal, kind, origin, execType, path, cashType, paymentType, method, algorithmId from events where parent_id = $1", prod.id);
    if err != nil {
        return err
    }
    defer rows.Close()
    prod.Terms = TermsStruct{}
    for rows.Next(){
        event := Event{}
        err := rows.Scan(&event.EventType, &event.Terminal, &event.Kind, &event.Origin,
				&event.ExecType, &event.Path, &event.CashType, &event.PaymentType, &event.Method, &event.AlgorithmId);
        if err != nil {
            return err
        }
        prod.Terms.Events = append(prod.Terms.Events, event)
    }
    return nil
}

//Method saves in the database all Event objects what pertain to this Product object
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

//Method saves Product object in the database
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

	//Creation date should be calculable
	prod.CreationDate = time.Now().Format("2006-01-02")

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

//Method updates Product object in the database
//Method accepts product_id value of Product to be updated as the parameter
//This is because product_id can be modified
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
        return errhand.ProdNotFound
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

//Method deleted Product from database by product_id value
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
        return errhand.ProdNotFound
    }
    err = tx.Commit()
    if err != nil{
        return err
    }
    return nil
}