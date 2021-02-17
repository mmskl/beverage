package DB

import (
    "gorm.io/gorm"
)

type Schema struct {}


func (s *Schema) upgrade() {
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
