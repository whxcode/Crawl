package model

import (
	_"database/sql"
	_"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var dbx,errx = sqlx.Open(`mysql`,`root@(127.0.0.1)/bolg`)

func init() {
	if errx != nil {
		log.Fatalln(`初始化失败:`,errx)
		os.Exit(1)
	}
}

func GetDbx() *sqlx.DB {
	return dbx
}