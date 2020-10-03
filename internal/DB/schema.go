package DB

import (
    "gorm.io/gorm"
    "fmt"
)

type Schema struct {}


func (s *Schema) Upgrade() {
    fmt.Println("Schema.Upgrade()")
    config := gorm.Config{
        DisableForeignKeyConstraintWhenMigrating: true,
    }

    c := Connect(&config)
    c.db.AutoMigrate(&EntryModel{})

}

func DoUpgrade() {
    s := Schema{}
    s.Upgrade()
}
