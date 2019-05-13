package route

import "net/http"

type patchTextReq struct {
	Title    *string
	Subtitle *string
	Detail   *string
	Xtext1   *string
}

func (p *patchTextReq) Bind(r *http.Request) error {
	return nil
}
