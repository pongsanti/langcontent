package route

import (
	"net/http"
	"time"
)

type postMcontentReq struct {
	Label       string
	ContentType string
	StartAt     *time.Time
	EndAt       *time.Time
	Status      *string
	Xtime1      *time.Time
	Xbool1      *bool
}

func (m *postMcontentReq) Bind(r *http.Request) error {
	if m.ContentType == "" {
		return errContentTypeRequired
	}
	return nil
}
