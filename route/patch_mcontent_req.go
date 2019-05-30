package route

import (
	"net/http"
	"time"

	"github.com/pongsanti/langcontent/share"
)

type patchMcontentReq struct {
	Label       *string
	ContentType string
	StartAt     *time.Time
	EndAt       *time.Time
	Status      *string
	Xtime1      *time.Time
	DeletedAt   share.JSONTime
}

func (m *patchMcontentReq) Bind(r *http.Request) error {
	return nil
}
