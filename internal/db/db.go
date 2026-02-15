package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/joseph0x45/goutils"
	"github.com/joseph0x45/sad"
)

type Conn struct {
	db *sqlx.DB
}

func (c *Conn) Close() {
	c.db.Close()
}

func GetConn() *Conn {
	dbPath := goutils.Setup()
	db, err := sad.OpenDBConnection(sad.DBConnectionOptions{
		Reset:             false,
		EnableForeignKeys: true,
		DatabasePath:      dbPath,
	}, migrations)
	if err != nil {
		panic(err)
	}
	return &Conn{db}
}
