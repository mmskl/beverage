package DB

import "gorm.io/gorm"

type Entry struct{
    gorm.Model
    Id int `gorm:"primary_key"`
    Name string
    Url string
}


// func (mgr *manager) AddEntry(entry *entry.Entry) (err error) {
//     mgr.db.Create(entry)
//     if errs := mgr.db.GetErrors(); len(errs) > 0 {
//         err = errs[0]
//     }
//     return
// }
