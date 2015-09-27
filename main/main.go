package main

import (
	"crypto/rand"
	"log"

	"github.com/cagnosolutions/realdb"
)

func main() {
	disk := realdb.NewDiskStore(1)
	b := make([]byte, (1024 * 128))
	if _, err := rand.Read(b); err != nil {
		log.Fatal(err)
	}
	disk.Write(b)
}
