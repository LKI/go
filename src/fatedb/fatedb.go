package fatedb

type IFateDB interface {
	Set(key, value interface{})
	Get(key interface{}) interface{}
	GetWithDefault(key, defaultValue interface{}) interface{}
	Del(key interface{})
}

func DB(filePath string) IFateDB {
	if filePath == "" {
		return InMemoryDB{}
	} else {
		return JsonFileDB{filePath}
	}
}
