package fatedb

type InMemoryDB struct{}

var memDB = make(map[interface{}]interface{})

func (db InMemoryDB) Get(key interface{}) interface{} {
	return memDB[key]
}

func (db InMemoryDB) GetWithDefault(key, defaultValue interface{}) interface{} {
	var value, exist = memDB[key]
	if exist {
		return value
	} else {
		return defaultValue
	}
}

func (db InMemoryDB) Set(key, value interface{}) {
	memDB[key] = value
}

func (db InMemoryDB) Del(key interface{}) {
	delete(memDB, key)
}
