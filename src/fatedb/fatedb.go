package main

import (
	"fmt"
	"os"
)

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

func usage() {
	fmt.Printf("%s <command> <key> <value>", os.Args[0])
	os.Exit(1)
}

func main() {
	var db = DB{}
	db.Set("5", "6")

	var args = os.Args[1:]
	var n = len(args)
	if n == 0 {
		usage()
	}

	switch args[0] {
	case "set":
		if n < 3 {
			usage()
		}
		db.Set(args[1], args[2])
	case "get":
		if n == 1 {
			usage()
		} else if n == 2 {
			fmt.Println(db.Get(args[1]))
		} else {
			fmt.Println(db.GetWithDefault(args[1], args[2]))
		}
	default:
		usage()
	}
}
