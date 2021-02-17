package DB

import (
    "log"
    _ "fmt"
    "os"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

const db_path = "./data/beverage.db"

type ConnectionType struct {
    db *gorm.DB
}

var Connection ConnectionType

var ConnectionConfig *gorm.Config



func SetCustomConnectConfig(config *gorm.Config) {
    ConnectionConfig = config
}

func Connect() *ConnectionType {

    var err error
    var db *gorm.DB

    if (ConnectionConfig != nil) {
        db, err = gorm.Open(sqlite.Open(db_path), ConnectionConfig)
    } else {
        config := gorm.Config{}
        db, err = gorm.Open(sqlite.Open(db_path), &config)
    }

    if (!fileExists(db_path)) {
        // @TODO move to init.sql file
        // db_path
    }

    if err != nil {
        log.Fatal("Failed to init db:", err)
    }

    Connection := ConnectionType{db: db}

    return &Connection
}

func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}



func (c *ConnectionType) AddEntry(entry *Entry) *ConnectionType {
    c.db.Create(entry)
    return c
}


func (c *ConnectionType) ListEntries() []Entry {

	entries := []Entry{}
    c.db.Find(&entries)

    return entries
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




