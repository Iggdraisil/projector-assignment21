package main

import (
	//  "bytes"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	// "github.com/lib/pq"
		_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	} else {
		fmt.Println("successful...")
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
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

	router := gin.New()
	router.Use(gin.Recovery())
        db.Exec("create table if not exists user (    id serial primary key,    name varchar(100),    birthdate DATE,    email varchar(100),    state varchar(100),    city varchar(100))")
	router.POST("/person", func(c *gin.Context) {
		address := faker.GetRealAddress()
		_, err := db.Exec("insert into user (`name`, `birthdate`, `email`, `state`, `city`) values (?,?,?,?,?)",
			faker.Name(), faker.Date(), faker.Email(), address.State, address.City,
		)
		if err != nil {
			c.JSON(500, gin.H{
				"user":  nil,
				"count": 0,
			})
			fmt.Println(err)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"count": 1,
			})
		}
	})
	router.GET("/person", func(c *gin.Context) {
		_, err := db.Exec("select * from user order by RAND() LIMIT 1")
		if err != nil {
			c.JSON(500, gin.H{
				"user":  nil,
				"count": 0,
			})
			fmt.Println(err)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"count": 1,
			})
		}
	})
	err = router.Run(":3000")
	if err != nil {
		panic(err)
	}
}