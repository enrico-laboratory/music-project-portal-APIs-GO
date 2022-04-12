package parser

import "time"

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
