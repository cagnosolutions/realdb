package realdb

import (
	"fmt"
	"log"
	"os"
	"path"
)

type DiskStore struct {
	size, count int
}

func NewDiskStore(maxmb int) *DiskStore {
	return &DiskStore{
		size:  (1024 * 1024) * maxmb,
		count: 0,
	}
}

func (d *DiskStore) path() string {
	return fmt.Sprintf("db/%d.dat", d.count)
}

func (d *DiskStore) Write(data []byte) {
	fd := open(d.path())
	defer fd.Close()
	info, err := fd.Stat()
	if err != nil {
		log.Panic(err)
	}
	if info.Size()+int64(len(data)) <= int64(d.size) {
		if _, err := fd.Write(data); err != nil {
			log.Fatal(err)
		}
		return
	}
	d.count++
	d.Write(data)
}

func (d *DiskStore) WriteLine(data []byte) {
	// stuff
}

func (d *DiskStore) Read(data []byte) {
	fd := open(d.path())
	defer fd.Close()
	info, err := fd.Stat()
	if err != nil {
		log.Fatal(err)
	}
	if info.Size()+int64(len(data)) <= int64(d.size) {
		if _, err := fd.Write(data); err != nil {
			log.Fatal(err)
		}
		return
	}
	d.count++
	d.Write(data)
}

func open(filepath string) *os.File {
	dir, file := path.Split(filepath)
	if dir != "" {
		if _, err := os.Stat(dir); err != nil && os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				log.Fatal(err)
			}
		}
	}
	if file != "" {
		if _, err := os.Stat(filepath); err != nil && os.IsNotExist(err) {
			if _, err := os.Create(filepath); err != nil {
				log.Fatal(err)
			}
		}
	}
	fd, err := os.OpenFile(filepath, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if err := fd.Sync(); err != nil {
		log.Fatal(err)
	}
	return fd
}
