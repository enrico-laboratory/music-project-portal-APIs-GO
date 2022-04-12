package notionApi

type MusicProjectsQueryFilter struct {
	Filter struct {
		Property string `json:"property"`
		Select   struct {
			Equals string `json:"equals"`
		} `json:"select"`
	} `json:"filter"`
}
