package model

type Page struct {
	Object       string        `json:"object"`
	Id           string        `json:"id"`
	CreatedTime  string        `json:"created_time"`
	CreatedBy    *User         `json:"created_by"`
	LastEditTime string        `json:"last_edit_time"`
	LastEditBy   *User         `json:"last_edit_by"`
	Archived     bool          `json:"archived"`
	Properties   *Properties   `json:"Properties"`
	Parent       *ParentObject `json:"parent"`
	Url          string        `json:"url"`
}

type Properties struct {
	Assign  *Assign  `json:"Assign"`
	BotLang *BotLang `json:"Bot Lang"`
	Date    *Date    `json:"Date"`
	Status  *Status  `json:"Status"`
	Bot     *Bot     `json:"Bot"`
	Name    *Name    `json:"Name"`
}

type Assign struct {
	ID     string  `json:"id"`
	Type   string  `json:"type"`
	People []*User `json:"people"`
}

type BotLang struct {
	ID          string         `json:"id"`
	Type        string         `json:"type"`
	MultiSelect []*MultiSelect `json:"multi_select"`
}

type Date struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Date *Dates `json:"date"`
}

type Dates struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type Status struct {
	ID     string  `json:"id"`
	Type   string  `json:"type"`
	Select *Select `json:"select"`
}

type Select struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Bot struct {
	ID          string         `json:"id"`
	Type        string         `json:"type"`
	MultiSelect []*MultiSelect `json:"multi_select"`
}

type Name struct {
	ID    string      `json:"id"`
	Type  string      `json:"type"`
	Title []*RichText `json:"title"`
}

type MultiSelect struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type DbProperties struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type ParentObject struct {
	DBParent        *DatabaseParent  `json:"db_parent"`
	PageParent      *PageParent      `json:"page_parent"`
	WorkSpaceParent *WorkSpaceParent `json:"work_space_parent"`
}

type DatabaseParent struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type PageParent struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type WorkSpaceParent struct {
	Type      string `json:"type"`
	WorkSpace bool   `json:"work_space"`
}
