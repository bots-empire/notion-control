package model

var AdminSettings *Admin

type Admin struct {
	AdminID map[int64]*AdminUser `json:"admin_id"`
}

type AdminUser struct {
	Name string `json:"name"`
}
