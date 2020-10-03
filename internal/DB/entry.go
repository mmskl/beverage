package DB

import "gorm.io/gorm"

type EntryModel struct{
    gorm.Model
    Id int `gorm:"primary_key"`
    Name string
    Url bool
}
