package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var(
	DBConn *gorm.DB
)

const DNS string = "root:root@tcp(localhost:3306)/notesdb?parseTime=true"
