package route

import (
	"github.com/pongsanti/langcontent"
	"github.com/pongsanti/langcontent/db/models"
)

type postTextRes struct {
	ContentText langcontent.McontentTextRes
}

func newPostTextRes(mcontText models.McontentText) postTextRes {
	return postTextRes{
		ContentText: langcontent.NewMcontentTextRes(&mcontText),
	}
}
