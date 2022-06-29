package model

type User struct {
	Object    string `json:"object"`
	Id        string `json:"id"`
	UserType  string `json:"type"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
}

type Results struct {
	Object string  `json:"object"`
	Page   []*Page `json:"results"`
}
