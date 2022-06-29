package model

type Block struct {
	Object       string     `json:"object"`
	ID           string     `json:"id"`
	BType        string     `json:"type"`
	CreatedTime  string     `json:"created_time"`
	CreatedBy    *User      `json:"created_by"`
	LastEditTime string     `json:"last_edit_time"`
	LastEditBy   *User      `json:"last_edit_by"`
	Archived     bool       `json:"archived"`
	HasChildren  bool       `json:"has_children"`
	Paragraph    *Paragraph `json:"paragraph"`
}

type Paragraph struct {
	RichText []*RichText `json:"rich_text"`
}

type RichText struct {
	Type        string       `json:"type"`
	Text        *Text        `json:"text"`
	Annotations *Annotations `json:"annotations"`
	PlainText   string       `json:"plain_text"`
	Href        interface{}  `json:"href"`
}

type Text struct {
	Content string      `json:"content"`
	Link    interface{} `json:"link"`
}

type Annotations struct {
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Strikethrough bool   `json:"strikethrough"`
	Underline     bool   `json:"underline"`
	Code          bool   `json:"code"`
	Color         string `json:"color"`
}
