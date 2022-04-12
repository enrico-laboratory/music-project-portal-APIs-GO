package main

import (
	"example.com/music-project-portal-APIs-GO/notionApi"
	"example.com/music-project-portal-APIs-GO/parser"
	"log"
	"os"
)

func main() {

	c, err := notionApi.NewClient(os.Getenv("NOTION_TOKEN"))
	if err != nil {
		log.Fatal("Error:", err)
	}
	var musicProjectsQueryFilter notionApi.MusicProjectsQueryFilter
	musicProjectsQueryFilter.Filter.Property = "Status"
	musicProjectsQueryFilter.Filter.Select.Equals = "On Going"

	// MUSIC PROJECT DATABASE
	response := c.Database.Query(MusicProjectDatabaseId, musicProjectsQueryFilter)
	var musicProjectModel parser.MusicProjectModel
	musicProjectParsed := musicProjectModel.GetMusicProjectParsed(response)
	for _, record := range musicProjectParsed {
		log.Print(record.Id, record.Title, record.Description)
	}

	// LOCATIONS DATABASE
	//response := c.Database.Query(LocationDatabaseId, nil)
	//var locationModel parser.LocationModels
	//mappedLocation := parser.Parser(locationModel, response)
	////log.Println(mappedLocation)
	//for _, location := range mappedLocation.Results {
	//	log.Print(location.Properties.Location.Title[0].PlainText)
	//}

	////Convert the body to type string
	//bodyString, _ := ioutil.ReadAll(response.Body)
	//sb := string(bodyString)
	//log.Println(sb)

	//
	//var database Database
	//
	//err := json.Unmarshal(body, &database)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	////fmt.Println(string(body))
	//fmt.Println(json)

}
