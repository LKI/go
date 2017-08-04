package fatedb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type JsonFileDB struct {
	filePath string
}

func (db JsonFileDB) Get(key interface{}) interface{} {
	var data = db.read()
	return data[key]
}

func (db JsonFileDB) GetWithDefault(key, defaultValue interface{}) interface{} {
	var data = db.read()
	var value, exist = data[key]
	if exist {
		return value
	} else {
		return defaultValue
	}
}

func (db JsonFileDB) Set(key, value interface{}) {
	var data = db.read()
	data[key] = value
	db.save(data)
}

func (db JsonFileDB) Del(key interface{}) {
	var data = db.read()
	delete(data, key)
	db.save(data)
}

func (db JsonFileDB) read() map[interface{}]interface{} {
	var data = make(map[interface{}]interface{})
	if _, err := os.Stat(db.filePath); os.IsNotExist(err) {
		return data
	}
	var content, err = ioutil.ReadFile(db.filePath)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	json.Unmarshal(content, &data)
	return data
}

func (db JsonFileDB) save(data map[interface{}]interface{}) {
	var content, err = json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	var file, _ = os.Create(db.filePath)
	defer file.Close()
	file.Write(content)
}
