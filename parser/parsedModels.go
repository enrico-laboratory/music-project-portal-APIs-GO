package parser

type ModelParsed interface {
	MusicProjectModelParsed
}

type MusicProjectModelParsed struct {
	Id          string
	Status      string
	Description string
	ChoirRollup string
	Published   bool
	Title       string
}
