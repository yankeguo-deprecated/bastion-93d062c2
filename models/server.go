package models

import (
	"strings"

	"ireul.com/com"
	"ireul.com/orm"
)

// Server 代表一个受管理的远端服务器
type Server struct {
	orm.Model
	Name        string   `json:"name" orm:"unique_index"`
	Address     string   `json:"address" orm:"index"`
	Port        uint     `json:"port"`
	Fingerprint string   `json:"fingerprint" orm:"index"`
	Tag         string   `json:"-" orm:"index"`
	Tags        []string `json:"tags" orm:"-"`
}

// BeforeSave save Tag from Tags, and append port 22 if necessary
func (s *Server) BeforeSave() error {
	s.Tag = strings.Join(s.Tags, ",") + ","
	if s.Port == 0 {
		s.Port = 22
	}
	return nil
}

// AfterFind set Tags from Tag
func (s *Server) AfterFind() error {
	s.Tags = com.CompactSliceStr(strings.Split(s.Tag, ","))
	return nil
}
