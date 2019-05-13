package langcontent

import (
	"time"

	"github.com/pongsanti/langcontent/db/models"
)

type McontentRes struct {
	ID          int
	CreatedAt   *time.Time
	ContentType string
	StartAt     *time.Time
	EndAt       *time.Time
	Status      *string
	Xtime1      *time.Time
	Texts       []McontentTextRes
}

func NewMcontentRes(mcont models.Mcontent) McontentRes {
	texts := NewMcontentTextReses(&mcont)

	return McontentRes{
		ID:          mcont.ID,
		CreatedAt:   mcont.CreatedAt.Ptr(),
		ContentType: mcont.ContentType,
		StartAt:     mcont.StartAt.Ptr(),
		EndAt:       mcont.EndAt.Ptr(),
		Status:      mcont.Status.Ptr(),
		Xtime1:      mcont.Xtime1.Ptr(),
		Texts:       texts,
	}
}

func NewMcontentTextReses(mcont *models.Mcontent) []McontentTextRes {
	if mcont == nil {
		return nil
	}
	texts := McontentTexts(mcont)
	if texts == nil {
		return nil
	}

	var reses []McontentTextRes
	for _, text := range texts {
		reses = append(reses, NewMcontentTextRest(text))
	}
	return reses
}

type McontentTextRes struct {
	ID        int
	CreatedAt *time.Time
	Lang      string
	Title     string
	Subtitle  *string
	Detail    *string
	Xtext1    *string
}

func NewMcontentTextRest(mcontText *models.McontentText) McontentTextRes {
	return McontentTextRes{
		ID:        mcontText.ID,
		CreatedAt: mcontText.CreatedAt.Ptr(),
		Lang:      mcontText.Lang,
		Title:     mcontText.Title,
		Subtitle:  mcontText.Subtitle.Ptr(),
		Detail:    mcontText.Detail.Ptr(),
		Xtext1:    mcontText.Xtext1.Ptr(),
	}
}
