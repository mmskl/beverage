package DB

import (
    "gorm.io/gorm"
    "fmt"
)

type Schema struct {}


func (s *Schema) upgrade() {
    fmt.Println("Schema.Upgrade()")
    config := gorm.Config{
        DisableForeignKeyConstraintWhenMigrating: true,
    }

    SetCustomConnectConfig(&config)
    c := Connect()
    c.db.AutoMigrate(&Entry{})

}

func DoUpgrade() {
    s := Schema{}
    s.upgrade()
}
