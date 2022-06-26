package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var instance *db

type db struct {
	Population map[string]int
}

func initDb() *db {
	database := &db{}
	database.Population = make(map[string]int, 10)
	database.Population["India"] = 1300
	database.Population["China"] = 1400
	return database
}

func GetDatabase() *db {
	once.Do(func() {
		data := initDb()
		instance = data
	})
	return instance
}

func main() {
	data := GetDatabase()
	fmt.Printf("%p\n", data)
	pata := GetDatabase()
	fmt.Printf("%p\n", pata)
}
