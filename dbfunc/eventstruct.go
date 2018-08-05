package dbfunc

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

//Event representation
type EventStruct struct {
    id int64
    parent_id int64
    EventType string `json:"type"`
    Terminal bool `json:"terminal"`
    ExecutionStruct `json:"execution"`
    CashDirectionStruct `json:"cashDirection"`
}

//Method InsertEvent of Event structure accepts opened transaction as the parameter
//and executes event record insertion into database
func (event *EventStruct) InsertEvent(tx *sql.Tx) error {
    result, err := tx.Exec("insert into events (parent_id, eventType, terminal, kind, origin, execType, path, cashType, paymentType, method, algorithmId) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
        event.parent_id, event.EventType, event.Terminal, event.Kind, event.Origin, event.ExecType,
				 event.Path, event.CashType, event.PaymentType, event.Method, event.AlgorithmId);
    if err != nil{
        tx.Rollback()
        return err
    }
    event.id, _ = result.LastInsertId();
    return nil
}
