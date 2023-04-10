package indexer

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Indexer struct {
	apiEndpoint string
	client      *http.Client
}

// postgresql://user:password@netloc:port/dbname
func New(apiEndpoint string) (*Indexer, error) {
	return &Indexer{
		apiEndpoint: apiEndpoint,
		client:      &http.Client{},
	}, nil
}

func IndexHistory(indexer *Indexer) error {

	return nil
}

// func GetLastBlock(indexer *Indexer) (int, error) {
// 	req, err := http.NewRequest("GET", indexer.apiEndpoint+"/sync/status", nil)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return 0, nil
// }

func GetSyncStatus(indexer *Indexer) (int, bool, error) {
	req, err := http.NewRequest("GET", indexer.apiEndpoint+"/sync/status", nil)
	if err != nil {
		return 0, false, err
	}

	resp, err := indexer.client.Do(req)
	if err != nil {
		return 0, false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, false, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, false, err
	}

	fmt.Println(string(body))

	return 0, false, nil
}
