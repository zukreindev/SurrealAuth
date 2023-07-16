package database

import (
	"FiberAuthWithSurrealDb/Util"
	"github.com/surrealdb/surrealdb.go"
)

var db *surrealdb.DB

func init() {
	var err error
	db, err = surrealdb.New(util.GetConfig("database", "url"))
	if err != nil {
		panic(err)
	}

	if _, err = db.Signin(map[string]interface{}{
		"user": util.GetConfig("database", "user"),
		"pass": util.GetConfig("database", "pass"),
	}); err != nil {
		panic(err)
	}
}

func Connect() {
	var err error
	db, err= surrealdb.New(util.GetConfig("database", "url"))
	if err != nil {
		panic(err)
	}

	if _, err = db.Signin(map[string]interface{}{
		"user": util.GetConfig("database", "user"),
		"pass": util.GetConfig("database", "pass"),
	}); err != nil {
		panic(err)
	}
	util.Log("Database", "Connected to database")
}

func Close() {
	db.Close()
	util.Log("Database", "Disconnected from database")
}
