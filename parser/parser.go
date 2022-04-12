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

func (m MusicProjectModel) GetMusicProjectParsed(res *http.Response) []MusicProjectModelParsed {

	mObject := getDatabaseObject(m, res)

	var musicProjectDatabaseParsed []MusicProjectModelParsed

	for _, record := range mObject.Results {
		musicProjectDatabaseParsedRecord := MusicProjectModelParsed{
			Id:          record.ID,
			Status:      record.Properties.Status.Select.Name,
			Description: record.Properties.Description.RichText[0].PlainText,
			ChoirRollup: record.Properties.ChoirRollup.Rollup.Array[0].Title[0].PlainText,
			Published:   record.Properties.Published.Checkbox,
			Title:       record.Properties.Title.Title[0].PlainText,
		}
		musicProjectDatabaseParsed = append(musicProjectDatabaseParsed, musicProjectDatabaseParsedRecord)
	}
	return musicProjectDatabaseParsed
}
