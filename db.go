package realdb

import (
	"log"
	"sync"
)

const mb = 1024 * 1024

type DataStore struct {
	store map[string]*Store
	sync.RWMutex
}

func NewDataStore() *DataStore {
	ds := &DataStore{
		store: make(map[string]*Store, 0),
	}
	fd := open("db/0.dat")
	defer fd.Close()
	info, err := fd.Stat()
	if err != nil {
		log.Fatal(err)
	}
	if info.Size() != 0 {
		ds.Load()
	}
	return ds
}

func (ds *DataStore) Load() {

}

func (ds *DataStore) Read(line int) {

}

func (ds *DataStore) Write() {

}
