package handlers

import (
	"github.com/BlackRRR/notion-control/model"
)

type Handlers struct {
	Page     *model.Page
	Block    *model.Block
	Database *model.Database
	Results  *model.Results
}

func InitHandler() *Handlers {
	rs := model.Results{Page: []*model.Page{{
		Object:       "",
		Id:           "",
		CreatedTime:  "",
		CreatedBy:    &model.User{},
		LastEditTime: "",
		LastEditBy:   &model.User{},
		Archived:     false,
		Properties: &model.Properties{
			Assign: &model.Assign{
				ID:   "",
				Type: "",
				People: []*model.User{{
					Object:    "",
					Id:        "",
					UserType:  "",
					Name:      "",
					AvatarUrl: ""}},
			},
			BotLang: &model.BotLang{},
			Date: &model.Date{
				ID:   "",
				Type: "",
				Date: &model.Dates{},
			},
			Status: &model.Status{
				ID:     "",
				Type:   "",
				Select: &model.Select{},
			},
			Bot: &model.Bot{},
			Name: &model.Name{
				ID:   "",
				Type: "",
				Title: []*model.RichText{{
					Type: "",
					Text: &model.Text{
						Content: "",
						Link:    nil,
					},
					Annotations: &model.Annotations{},
					PlainText:   "",
					Href:        nil,
				}},
			},
		},
		Parent: &model.ParentObject{
			DBParent:        &model.DatabaseParent{},
			PageParent:      &model.PageParent{},
			WorkSpaceParent: &model.WorkSpaceParent{},
		},
		Url: "",
	}}}
	handler := &Handlers{
		Results: &rs,
	}

	return handler

}
