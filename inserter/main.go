package main

import (
	//  "bytes"
	"database/sql"
	"fmt"
        "time"
	"github.com/go-faker/faker/v4"
	// "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	} else {
		fmt.Println("successful...")
	}
}

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/mydb")
	checkErr(err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	err = db.Ping()
	checkErr(err)
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(5000)
        db.Exec("create table if not exists user (    id serial primary key,    name varchar(100),    birthdate DATE,    email varchar(100),    state varchar(100),    city varchar(100))")
	for {
		address := faker.GetRealAddress()
		db.Exec("insert into user (`name`, `birthdate`, `email`, `state`, `city`) values (?,?,?,?,?)",
			faker.Name(), faker.Date(), faker.Email(), address.State, address.City,
		)
		time.Sleep(time.Second/20)

	}
}
