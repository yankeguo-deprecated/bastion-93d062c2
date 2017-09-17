package routes

import (
	"strings"

	"ireul.com/bastion/models"
	"ireul.com/web"
)

// ServerCreateForm the form to create a server
type ServerCreateForm struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Port    int    `json:"port"`
	Tag     string `json:"tag"`
	Desc    string `json:"desc"`
}

// ServerCreate create a new server
func ServerCreate(ctx *web.Context, db *models.DB, r APIRender, a Auth, f ServerCreateForm) {
	if !models.ServerNameRegexp.MatchString(f.Name) {
		r.Fail(ParamsInvalid, "服务器名称不合法")
		return
	}
	if len(f.Address) == 0 || len(f.Address) > models.ServerAddressMaxLen {
		r.Fail(ParamsInvalid, "服务器地址不合法")
		return
	}
	if f.Port < 0 || f.Port >= 65536 {
		f.Port = 22
	}
	if len(f.Tag) > 100 {
		r.Fail(ParamsInvalid, "服务器标签过长")
		return
	}
	tags := strings.Split(f.Tag, ",")
	for _, v := range tags {
		if !models.ServerTagRegexp.MatchString(strings.TrimSpace(v)) {
			r.Fail(ParamsInvalid, "标签 \""+v+"\"不合法")
			return
		}
	}
	if len(f.Desc) > 100 {
		r.Fail(ParamsInvalid, "服务器备注过长")
		return
	}

	// check duplicates
	s := &models.Server{}
	db.Where("name = ?", f.Name).First(s)
	if !db.NewRecord(s) {
		r.Fail(ParamsInvalid, "服务器名称不能重复")
		return
	}
	s = &models.Server{
		Name:    f.Name,
		Address: f.Address,
		Port:    uint(f.Port),
		Tags:    tags,
	}
	err := db.Create(s).Error
	if err != nil {
		r.Fail(InternalError, err.Error())
		return
	}

	r.Success("server", s)
}

// ServerList list all servers
func ServerList(ctx *web.Context, db *models.DB, r APIRender, a Auth) {
	list := []models.Server{}
	db.Find(&list)
	r.Success("servers", list)
}
