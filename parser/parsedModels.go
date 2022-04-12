package parser

type ModelParsed interface {
	MusicProjectModelParsed
}

type MusicProjectModelParsed struct {
	Id          string `bson:"_id" json:"id"`
	Title       string `bson:"title" json:"title"`
	Status      string `bson:"status" json:"status"`
	Description string `bson:"description" json:"description"`
	Published   bool   `bson:"published" json:"published"`
	ChoirRollup string `bson:"choirRollup" json:"choirRollup"`

	Repertoire string `bson:"repertoire" json:"repertoire"`
	//Tasks      string
	//Cast       string
}
