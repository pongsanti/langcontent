package route

import (
	"github.com/pongsanti/langcontent"
	"github.com/pongsanti/langcontent/db/models"
)

type postMcontentRes struct {
	Content langcontent.McontentRes
}

func newPostMcontentRes(mcont models.Mcontent) postMcontentRes {
	return postMcontentRes{
		Content: langcontent.NewMcontentRes(mcont),
	}
}
