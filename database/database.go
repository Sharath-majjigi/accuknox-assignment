package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var(
	DBConn *gorm.DB
)

const DNS string = "admin:Sharath2001@tcp(database-2.cidhjwutum25.ap-southeast-1.rds.amazonaws.com:3306)/notesdb?parseTime=true"
