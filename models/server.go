package models

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"

	"ireul.com/com"
	"ireul.com/orm"
)

// ServerTagDefault is the default server tag applied to all server automatically
const ServerTagDefault = "default"

// ServerNameRegexp 用户登录名正则表达式
var ServerNameRegexp = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]{2,23}$`)

// ServerTagRegexp regexp for a single tag
var ServerTagRegexp = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]{0,23}$`)

// ServerAddressMaxLen max length of address
const ServerAddressMaxLen = 36

// Server 代表一个受管理的远端服务器
type Server struct {
	orm.Model
	Name        string   `json:"name" orm:"unique_index"`
	Address     string   `json:"address" orm:"index"`
	Port        uint     `json:"port"`
	Fingerprint string   `json:"fingerprint" orm:"index"`
	Tag         string   `json:"-" orm:"index"`
	Tags        []string `json:"tags" orm:"-"`
	Desc        string   `json:"desc" orm:"type:text"`
	Token       string   `json:"token" orm:"unique_index"`
}

// BeforeSave save Tag from Tags, and append port 22 if necessary
func (s *Server) BeforeSave() error {
	// create Tag from Tags
	s.Tags = com.CompactSliceStr(append(s.Tags, ServerTagDefault))
	s.Tag = "," + strings.Join(s.Tags, ",") + ","
	// assign port 22 by default
	if s.Port == 0 {
		s.Port = 22
	}
	// assign random token
	if len(s.Token) == 0 {
		bytes := make([]byte, 16)
		rand.Read(bytes)
		s.Token = hex.EncodeToString(bytes)
	}
	return nil
}

// AfterFind set Tags from Tag
func (s *Server) AfterFind() error {
	// recover Tags from Tag
	s.Tags = com.CompactSliceStr(strings.Split(s.Tag, ","))
	return nil
}

// AuditableName implements types.Auditable
func (s Server) AuditableName() string {
	return fmt.Sprintf("Server(%d, %s)", s.ID, s.Name)
}

// AuditableDetail implements types.Auditable
func (s Server) AuditableDetail() string {
	return fmt.Sprintf("%s:%d, %s", s.Address, s.Port, s.Tag)
}
