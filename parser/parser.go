package parser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func getDatabaseObject[K Model](databaseModel K, res *http.Response) K {
	bodyByte, _ := ioutil.ReadAll(res.Body)

	err := json.Unmarshal(bodyByte, &databaseModel)
	if err != nil {
		log.Fatal(err)
		return databaseModel
	}
	return databaseModel
}
