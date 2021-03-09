package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-training/persistence-layer-for-users-and-recipies/db"
	"go-training/persistence-layer-for-users-and-recipies/entities"
	logging_functionality "go-training/persistence-layer-for-users-and-recipies/logging-functionality"
	"go-training/persistence-layer-for-users-and-recipies/printutils"
	"log"
	"time"
)

func main() {

	var err error
	db.DB, err =  sql.Open("mysql", "root:23sl_@/user_recipes?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.DB.Close()

	db.DB.SetConnMaxLifetime(time.Minute*5)
	db.DB.SetMaxOpenConns(10)
	db.DB.SetMaxIdleConns(10)
	db.DB.SetConnMaxIdleTime(time.Minute*3)

	var userIds []uint
	TestUsers(&userIds)
	TestRecipes(userIds)

	lc := logging_functionality.LoginController{}
	if user, err := lc.Login("twinevidence", "3zs`/Pq`P?"); err != nil {
		fmt.Println(err)
	} else {
		printutils.PrintUsers([]entities.User{user})
	}

	loggedUser := lc.GetLoggedUser()
	printutils.PrintUsers([]entities.User{loggedUser})

	lc.Logout()



}
