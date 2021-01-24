package configs

import "sync"

func CreateConfigDB() *ConfigDB {
	return &ConfigDB{make(configs, 0), make(dictApi, 0), sync.Mutex{}}
}