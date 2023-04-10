package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/santi1234567/blockprint-indexer/indexer"
)

func main() {
	indexer, err := indexer.New("https://api.blockprint.sigp.io/")
	if err != nil {
		log.Fatal(err)
	}
	indexer.GetSyncStatus()

}
