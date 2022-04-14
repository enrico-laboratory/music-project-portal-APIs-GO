package parser

import "time"

type Model interface {
	MusicProjectModel | LocationModels
}

type MusicProjectModel struct {
	Object  string `json:"object"`
	Results []struct {
		Object         string    `json:"object"`
		ID             string    `json:"id"`
		CreatedTime    time.Time `json:"created_time"`
		LastEditedTime time.Time `json:"last_edited_time"`
		CreatedBy      struct {
			Object string `json:"object"`
			ID     string `json:"id"`
		} `json:"created_by"`
		LastEditedBy struct {
			Object string `json:"object"`
			ID     string `json:"id"`
		} `json:"last_edited_by"`
		Cover  interface{} `json:"cover"`
		Icon   interface{} `json:"icon"`
		Parent struct {
			Type       string `json:"type"`
			DatabaseID string `json:"database_id"`
		} `json:"parent"`
		Archived   bool `json:"archived"`
		Properties struct {
			CreatedTime struct {
				ID          string    `json:"id"`
				Type        string    `json:"type"`
				CreatedTime time.Time `json:"created_time"`
			} `json:"Created Time"`
			StartDate struct {
				ID   string `json:"id"`
				Type string `json:"type"`
				Date struct {
					Start    string      `json:"start"`
					End      interface{} `json:"end"`
					TimeZone interface{} `json:"time_zone"`
				} `json:"date"`
			} `json:"Start Date"`
			Year struct {
				ID     string `json:"id"`
				Type   string `json:"type"`
				Number int    `json:"number"`
			} `json:"Year"`
			Status struct {
				ID     string `json:"id"`
				Type   string `json:"type"`
				Select struct {
					ID    string `json:"id"`
					Name  string `json:"name"`
					Color string `json:"color"`
				} `json:"select"`
			} `json:"Status"`
			Repertoire struct {
				ID       string        `json:"id"`
				Type     string        `json:"type"`
				Relation []interface{} `json:"relation"`
			} `json:"Repertoire"`
			Cancelled struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				Checkbox bool   `json:"checkbox"`
			} `json:"Cancelled"`
			Task struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				Relation []struct {
					ID string `json:"id"`
				} `json:"relation"`
			} `json:"Task"`
			Description struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				RichText []struct {
					Type string `json:"type"`
					Text struct {
						Content string      `json:"content"`
						Link    interface{} `json:"link"`
					} `json:"text"`
					Annotations struct {
						Bold          bool   `json:"bold"`
						Italic        bool   `json:"italic"`
						Strikethrough bool   `json:"strikethrough"`
						Underline     bool   `json:"underline"`
						Code          bool   `json:"code"`
						Color         string `json:"color"`
					} `json:"annotations"`
					PlainText string      `json:"plain_text"`
					Href      interface{} `json:"href"`
				} `json:"rich_text"`
			} `json:"Description"`
			Completed struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				Checkbox bool   `json:"checkbox"`
			} `json:"Completed"`
			Notes struct {
				ID       string        `json:"id"`
				Type     string        `json:"type"`
				Relation []interface{} `json:"relation"`
			} `json:"Notes"`
			Cast struct {
				ID       string        `json:"id"`
				Type     string        `json:"type"`
				Relation []interface{} `json:"relation"`
			} `json:"Cast"`
			Choir struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				Relation []struct {
					ID string `json:"id"`
				} `json:"relation"`
			} `json:"Choir"`
			MediaVault struct {
				ID       string        `json:"id"`
				Type     string        `json:"type"`
				Relation []interface{} `json:"relation"`
			} `json:"Media Vault"`
			TaskRollup struct {
				ID     string `json:"id"`
				Type   string `json:"type"`
				Rollup struct {
					Type  string `json:"type"`
					Array []struct {
						Type  string `json:"type"`
						Title []struct {
							Type string `json:"type"`
							Text struct {
								Content string      `json:"content"`
								Link    interface{} `json:"link"`
							} `json:"text"`
							Annotations struct {
								Bold          bool   `json:"bold"`
								Italic        bool   `json:"italic"`
								Strikethrough bool   `json:"strikethrough"`
								Underline     bool   `json:"underline"`
								Code          bool   `json:"code"`
								Color         string `json:"color"`
							} `json:"annotations"`
							PlainText string      `json:"plain_text"`
							Href      interface{} `json:"href"`
						} `json:"Title"`
					} `json:"array"`
					Function string `json:"function"`
				} `json:"rollup"`
			} `json:"Task Rollup"`
			ChoirRollup struct {
				ID     string `json:"id"`
				Type   string `json:"type"`
				Rollup struct {
					Type  string `json:"type"`
					Array []struct {
						Type  string `json:"type"`
						Title []struct {
							Type string `json:"type"`
							Text struct {
								Content string      `json:"content"`
								Link    interface{} `json:"link"`
							} `json:"text"`
							Annotations struct {
								Bold          bool   `json:"bold"`
								Italic        bool   `json:"italic"`
								Strikethrough bool   `json:"strikethrough"`
								Underline     bool   `json:"underline"`
								Code          bool   `json:"code"`
								Color         string `json:"color"`
							} `json:"annotations"`
							PlainText string      `json:"plain_text"`
							Href      interface{} `json:"href"`
						} `json:"Title"`
					} `json:"array"`
					Function string `json:"function"`
				} `json:"rollup"`
			} `json:"Choir Rollup"`
			EndDate struct {
				ID   string `json:"id"`
				Type string `json:"type"`
				Date struct {
					Start    string      `json:"start"`
					End      interface{} `json:"end"`
					TimeZone interface{} `json:"time_zone"`
				} `json:"date"`
			} `json:"End Date"`
			Published struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				Checkbox bool   `json:"checkbox"`
			} `json:"Published"`
			Projects struct {
				ID       string        `json:"id"`
				Type     string        `json:"type"`
				Relation []interface{} `json:"relation"`
			} `json:"Projects"`
			Title struct {
				ID    string `json:"id"`
				Type  string `json:"type"`
				Title []struct {
					Type string `json:"type"`
					Text struct {
						Content string      `json:"content"`
						Link    interface{} `json:"link"`
					} `json:"text"`
					Annotations struct {
						Bold          bool   `json:"bold"`
						Italic        bool   `json:"italic"`
						Strikethrough bool   `json:"strikethrough"`
						Underline     bool   `json:"underline"`
						Code          bool   `json:"code"`
						Color         string `json:"color"`
					} `json:"annotations"`
					PlainText string      `json:"plain_text"`
					Href      interface{} `json:"href"`
				} `json:"Title"`
			} `json:"Title"`
		} `json:"properties"`
		URL string `json:"url"`
	} `json:"results"`
	NextCursor interface{} `json:"next_cursor"`
	HasMore    bool        `json:"has_more"`
}

type LocationModels struct {
	Object  string `json:"object"`
	Results []struct {
		Object         string    `json:"object"`
		ID             string    `json:"id"`
		CreatedTime    time.Time `json:"created_time"`
		LastEditedTime time.Time `json:"last_edited_time"`
		CreatedBy      struct {
			Object string `json:"object"`
			ID     string `json:"id"`
		} `json:"created_by"`
		LastEditedBy struct {
			Object string `json:"object"`
			ID     string `json:"id"`
		} `json:"last_edited_by"`
		Cover  interface{} `json:"cover"`
		Icon   interface{} `json:"icon"`
		Parent struct {
			Type       string `json:"type"`
			DatabaseID string `json:"database_id"`
		} `json:"parent"`
		Archived   bool `json:"archived"`
		Properties struct {
			Contact struct {
				ID       string        `json:"id"`
				Type     string        `json:"type"`
				Relation []interface{} `json:"relation"`
			} `json:"Contact"`
			Phone struct {
				ID     string `json:"id"`
				Type   string `json:"type"`
				Rollup struct {
					Type     string        `json:"type"`
					Array    []interface{} `json:"array"`
					Function string        `json:"function"`
				} `json:"rollup"`
			} `json:"Phone"`
			City struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				RichText []struct {
					Type string `json:"type"`
					Text struct {
						Content string      `json:"content"`
						Link    interface{} `json:"link"`
					} `json:"text"`
					Annotations struct {
						Bold          bool   `json:"bold"`
						Italic        bool   `json:"italic"`
						Strikethrough bool   `json:"strikethrough"`
						Underline     bool   `json:"underline"`
						Code          bool   `json:"code"`
						Color         string `json:"color"`
					} `json:"annotations"`
					PlainText string      `json:"plain_text"`
					Href      interface{} `json:"href"`
				} `json:"rich_text"`
			} `json:"City"`
			Email struct {
				ID     string `json:"id"`
				Type   string `json:"type"`
				Rollup struct {
					Type     string        `json:"type"`
					Array    []interface{} `json:"array"`
					Function string        `json:"function"`
				} `json:"rollup"`
			} `json:"Email"`
			Tasks struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				Relation []struct {
					ID string `json:"id"`
				} `json:"relation"`
			} `json:"Tasks"`
			Purpose struct {
				ID          string        `json:"id"`
				Type        string        `json:"type"`
				MultiSelect []interface{} `json:"multi_select"`
			} `json:"Purpose"`
			Address struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				RichText []struct {
					Type string `json:"type"`
					Text struct {
						Content string      `json:"content"`
						Link    interface{} `json:"link"`
					} `json:"text"`
					Annotations struct {
						Bold          bool   `json:"bold"`
						Italic        bool   `json:"italic"`
						Strikethrough bool   `json:"strikethrough"`
						Underline     bool   `json:"underline"`
						Code          bool   `json:"code"`
						Color         string `json:"color"`
					} `json:"annotations"`
					PlainText string      `json:"plain_text"`
					Href      interface{} `json:"href"`
				} `json:"rich_text"`
			} `json:"Address"`
			Location struct {
				ID    string `json:"id"`
				Type  string `json:"type"`
				Title []struct {
					Type string `json:"type"`
					Text struct {
						Content string      `json:"content"`
						Link    interface{} `json:"link"`
					} `json:"text"`
					Annotations struct {
						Bold          bool   `json:"bold"`
						Italic        bool   `json:"italic"`
						Strikethrough bool   `json:"strikethrough"`
						Underline     bool   `json:"underline"`
						Code          bool   `json:"code"`
						Color         string `json:"color"`
					} `json:"annotations"`
					PlainText string      `json:"plain_text"`
					Href      interface{} `json:"href"`
				} `json:"Title"`
			} `json:"Location"`
		} `json:"properties"`
		URL string `json:"url"`
	} `json:"results"`
	NextCursor interface{} `json:"next_cursor"`
	HasMore    bool        `json:"has_more"`
}

type Contact struct {
	Object  string `json:"object"`
	Results []struct {
		Object         string    `json:"object"`
		ID             string    `json:"id"`
		CreatedTime    time.Time `json:"created_time"`
		LastEditedTime time.Time `json:"last_edited_time"`
		CreatedBy      struct {
			Object string `json:"object"`
			ID     string `json:"id"`
		} `json:"created_by"`
		LastEditedBy struct {
			Object string `json:"object"`
			ID     string `json:"id"`
		} `json:"last_edited_by"`
		Cover struct {
			Type string `json:"type"`
			File struct {
				URL        string    `json:"url"`
				ExpiryTime time.Time `json:"expiry_time"`
			} `json:"file"`
		} `json:"cover"`
		Icon   interface{} `json:"icon"`
		Parent struct {
			Type       string `json:"type"`
			DatabaseID string `json:"database_id"`
		} `json:"parent"`
		Archived   bool `json:"archived"`
		Properties struct {
			CreatedTime struct {
				ID          string    `json:"id"`
				Type        string    `json:"type"`
				CreatedTime time.Time `json:"created_time"`
			} `json:"Created Time"`
			StartDate struct {
				ID   string      `json:"id"`
				Type string      `json:"type"`
				Date interface{} `json:"date"`
			} `json:"Start Date"`
			Year struct {
				ID     string `json:"id"`
				Type   string `json:"type"`
				Number int    `json:"number"`
			} `json:"Year"`
			Status struct {
				ID     string `json:"id"`
				Type   string `json:"type"`
				Select struct {
					ID    string `json:"id"`
					Name  string `json:"name"`
					Color string `json:"color"`
				} `json:"select"`
			} `json:"Status"`
			Repertoire struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				Relation []struct {
					ID string `json:"id"`
				} `json:"relation"`
			} `json:"Repertoire"`
			Cancelled struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				Checkbox bool   `json:"checkbox"`
			} `json:"Cancelled"`
			Task struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				Relation []struct {
					ID string `json:"id"`
				} `json:"relation"`
			} `json:"Task"`
			Description struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				RichText []struct {
					Type string `json:"type"`
					Text struct {
						Content string      `json:"content"`
						Link    interface{} `json:"link"`
					} `json:"text"`
					Annotations struct {
						Bold          bool   `json:"bold"`
						Italic        bool   `json:"italic"`
						Strikethrough bool   `json:"strikethrough"`
						Underline     bool   `json:"underline"`
						Code          bool   `json:"code"`
						Color         string `json:"color"`
					} `json:"annotations"`
					PlainText string      `json:"plain_text"`
					Href      interface{} `json:"href"`
				} `json:"rich_text"`
			} `json:"Description"`
			Completed struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				Checkbox bool   `json:"checkbox"`
			} `json:"Completed"`
			Notes struct {
				ID       string        `json:"id"`
				Type     string        `json:"type"`
				Relation []interface{} `json:"relation"`
			} `json:"Notes"`
			Cast struct {
				ID       string        `json:"id"`
				Type     string        `json:"type"`
				Relation []interface{} `json:"relation"`
			} `json:"Cast"`
			Choir struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				Relation []struct {
					ID string `json:"id"`
				} `json:"relation"`
			} `json:"Choir"`
			MediaVault struct {
				ID       string        `json:"id"`
				Type     string        `json:"type"`
				Relation []interface{} `json:"relation"`
			} `json:"Media Vault"`
			TaskRollup struct {
				ID     string `json:"id"`
				Type   string `json:"type"`
				Rollup struct {
					Type  string `json:"type"`
					Array []struct {
						Type  string `json:"type"`
						Title []struct {
							Type string `json:"type"`
							Text struct {
								Content string      `json:"content"`
								Link    interface{} `json:"link"`
							} `json:"text"`
							Annotations struct {
								Bold          bool   `json:"bold"`
								Italic        bool   `json:"italic"`
								Strikethrough bool   `json:"strikethrough"`
								Underline     bool   `json:"underline"`
								Code          bool   `json:"code"`
								Color         string `json:"color"`
							} `json:"annotations"`
							PlainText string      `json:"plain_text"`
							Href      interface{} `json:"href"`
						} `json:"title"`
					} `json:"array"`
					Function string `json:"function"`
				} `json:"rollup"`
			} `json:"Task Rollup"`
			ChoirRollup struct {
				ID     string `json:"id"`
				Type   string `json:"type"`
				Rollup struct {
					Type  string `json:"type"`
					Array []struct {
						Type  string `json:"type"`
						Title []struct {
							Type string `json:"type"`
							Text struct {
								Content string      `json:"content"`
								Link    interface{} `json:"link"`
							} `json:"text"`
							Annotations struct {
								Bold          bool   `json:"bold"`
								Italic        bool   `json:"italic"`
								Strikethrough bool   `json:"strikethrough"`
								Underline     bool   `json:"underline"`
								Code          bool   `json:"code"`
								Color         string `json:"color"`
							} `json:"annotations"`
							PlainText string      `json:"plain_text"`
							Href      interface{} `json:"href"`
						} `json:"title"`
					} `json:"array"`
					Function string `json:"function"`
				} `json:"rollup"`
			} `json:"Choir Rollup"`
			EndDate struct {
				ID   string      `json:"id"`
				Type string      `json:"type"`
				Date interface{} `json:"date"`
			} `json:"End Date"`
			Published struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				Checkbox bool   `json:"checkbox"`
			} `json:"Published"`
			Projects struct {
				ID       string        `json:"id"`
				Type     string        `json:"type"`
				Relation []interface{} `json:"relation"`
			} `json:"Projects"`
			Title struct {
				ID    string `json:"id"`
				Type  string `json:"type"`
				Title []struct {
					Type string `json:"type"`
					Text struct {
						Content string      `json:"content"`
						Link    interface{} `json:"link"`
					} `json:"text"`
					Annotations struct {
						Bold          bool   `json:"bold"`
						Italic        bool   `json:"italic"`
						Strikethrough bool   `json:"strikethrough"`
						Underline     bool   `json:"underline"`
						Code          bool   `json:"code"`
						Color         string `json:"color"`
					} `json:"annotations"`
					PlainText string      `json:"plain_text"`
					Href      interface{} `json:"href"`
				} `json:"title"`
			} `json:"Title"`
		} `json:"properties"`
		URL string `json:"url"`
	} `json:"results"`
	NextCursor interface{} `json:"next_cursor"`
	HasMore    bool        `json:"has_more"`
}
