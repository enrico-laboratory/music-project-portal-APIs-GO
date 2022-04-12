package main

import (
	"context"
	"example.com/music-project-portal-APIs-GO/notionApi"
	"example.com/music-project-portal-APIs-GO/parser"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func main() {

	// NOTION CLIENT
	c, err := notionApi.NewClient(os.Getenv("NOTION_TOKEN"))
	if err != nil {
		log.Fatal("Error:", err)
	}

	// MUSIC PROJECT FILTERS
	var musicProjectsQueryFilter notionApi.MusicProjectsQueryFilter
	musicProjectsQueryFilter.Filter.Property = "Status"
	musicProjectsQueryFilter.Filter.Select.Equals = "On Going"

	// MUSIC PROJECT DATABASE
	response := c.Database.Query(MusicProjectDatabaseId, musicProjectsQueryFilter)
	var musicProjectModel parser.MusicProjectModel
	musicProjectParsed := musicProjectModel.GetMusicProjectParsed(response)
	//for _, record := range musicProjectParsed {
	//	log.Print(record.Id, record.Title, record.Description)
	//}

	//// LOCATIONS DATABASE
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

	// MONGODB CLIENT
	client, err := parser.NewClient()
	if err != nil {
		log.Fatal("Damn!!")
	}

	// MONGODB UPDATE
	musicProjectCollection := client.Database("music_projects_portal").Collection("music_project")
	opts := options.Update().SetUpsert(true)

	// TODO fix error: cannot transform type []parser.MusicProjectModelParsed to a BSON Document:
	// TODO WriteArray can only write a Array while positioned on a Element or Value but is positioned on a TopLevel
	many, err := musicProjectCollection.UpdateMany(context.TODO(), musicProjectParsed, opts)
	if err != nil {
		log.Print(err)
	}

	log.Print(many.UpsertedCount)

}
