package model

type Database struct {
	Object       string        `json:"object"`
	Id           string        `json:"id"`
	CreateTime   string        `json:"create_time"`
	LastEditTime string        `json:"last_edited_time"`
	LastEditBy   *User         `json:"last_edit_by"`
	Title        []*RichText   `json:"title"`
	Description  []*RichText   `json:"description"`
	Icon         string        `json:"icon"`
	Cover        string        `json:"cover"`
	Properties   *Properties   `json:"Properties"`
	Parent       *ParentObject `json:"parent"`
	Url          string        `json:"url"`
	Archived     bool          `json:"archived"`
	IsInline     bool          `json:"is_inline"`
}
