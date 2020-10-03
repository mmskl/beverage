package DB

import (
    "fmt"
    "log"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

const db_path = "../../data/beverage.db"

type ConnectionType struct {
    db *gorm.DB
}

var Connection ConnectionType

func Connect(config *gorm.Config) *ConnectionType {
    fmt.Println("connecting to DB")

    if (!fileExists(db_path)) {
        db_path
    }

    db, err := gorm.Open(sqlite.Open(db_path), config)
    if err != nil {
        log.Fatal("Failed to init db:", err)
    }

    Connection := &ConnectionType{db: db}

    return Connection
}




// type manager struct {
//     db *gorm.DB
// }

// var Mgr manager




// func (mgr *manager) AddEntry(entry *entry.Entry) (err error) {
//     mgr.db.Create(entry)
//     if errs := mgr.db.GetErrors(); len(errs) > 0 {
//         err = errs[0]
//     }
//     return
// }




