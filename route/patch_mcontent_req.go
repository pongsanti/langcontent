package route

import (
	"net/http"
	"time"
)

type patchMcontentReq struct {
	Label       *string
	ContentType string
	StartAt     *time.Time
	EndAt       *time.Time
	Status      *string
	Xtime1      *time.Time
}

func (m *patchMcontentReq) Bind(r *http.Request) error {
	return nil
}
