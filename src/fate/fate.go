package main

import (
	"fatedb"
	"fmt"
)

func main() {
	var db = fatedb.DB("C:\\Code\\fate.json")
	db.Set("123", "234")
	fmt.Println(db.Get("123"))
}
