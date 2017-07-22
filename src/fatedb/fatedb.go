package main

import "fmt"

var inMemoryDB = make(map[interface{}]interface{})

type DB struct {
}

func (db DB) Get(key interface{}) interface{} {
	return inMemoryDB[key]
}

func (db DB) GetWithDefault(key, defaultValue interface{}) interface{} {
	var value, exist = inMemoryDB[key]
	if exist {
		return value
	} else {
		return defaultValue
	}
}

func (db DB) Set(key, value interface{}) {
	inMemoryDB[key] = value
}

func (db DB) Del(key interface{}) {
	delete(inMemoryDB, key)
}

func main() {
	var db = DB{}
	db.Set("5", "6")

	fmt.Println(db.Get("5"))
	fmt.Println(db.Get(6))
	fmt.Println(db.GetWithDefault("7", "100"))

	var anotherDB = DB{}
	anotherDB.Set(6, "7")
	fmt.Println(anotherDB.Get("5"))
	fmt.Println(db.Get(6))
	fmt.Println(anotherDB.GetWithDefault("7", "100"))
}
