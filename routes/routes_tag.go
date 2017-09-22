package routes

import (
	"ireul.com/bastion/models"
	"ireul.com/com"
	"ireul.com/web"
)

// TagList list all tags
func TagList(ctx *web.Context, r APIRender, db *models.DB) {
	// all tags from servers
	ss := []models.Server{}
	tags := []string{models.ServerTagDefault}
	db.Find(&ss)
	for _, s := range ss {
		for _, t := range s.Tags {
			tags = append(tags, t)
		}
	}
	// all tags from grants
	gt := []models.Grant{}
	db.Select("DISTINCT tag").Find(&gt)
	for _, g := range gt {
		tags = append(tags, g.Tag)
	}

	r.Success("tags", com.CompactSliceStr(tags))
}
