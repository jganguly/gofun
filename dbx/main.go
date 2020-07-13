package main

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

type StructTab struct {
	Co1  string			// uppercase for json.Marshall to work
	Co2  int
	Co3	 float64
	Co4  mysql.NullTime
}


func main() {

	sql := "SELECT Co1, Co2, Co3, Co4 from tpl.tab"
	db := dbConn()
	ch := make(chan []StructTab, 10)
	var rows []StructTab
	dbSelectx(db, sql, ch)
	select {
	case rs := <-ch: // get the result from the channel
		rows = rs
	case <-time.After(5 * time.Second):
		fmt.Println("timeout")
	}
	fmt.Println(rows)
}

func dbConn() (db *sqlx.DB) {
	dbDriver := "mysql"
	db, err := sqlx.Open(dbDriver, "root:root@tcp(127.0.0.1:3306)/tpl")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func dbSelectx(db *sqlx.DB, sql string, ch chan []StructTab) {
	row := StructTab{}
	rs := make([]StructTab, 0)

	// Query
	rows, err := db.Queryx(sql)
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		err = rows.StructScan(&row)
		rs = append(rs,row)
	}
	ch <- rs
	close(ch)
}


func dbCrudx(db *sqlx.DB, sql string) {
	fmt.Println(sql)
	_, err := db.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	db.Exec(sql)
}