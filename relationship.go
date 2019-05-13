package langcontent

import "github.com/pongsanti/langcontent/db/models"

func McontentTexts(mcont *models.Mcontent) models.McontentTextSlice {
	if mcont == nil {
		return nil
	}

	mcontR := mcont.R
	if mcontR == nil {
		return nil
	}

	return mcontR.McontentTexts
}
