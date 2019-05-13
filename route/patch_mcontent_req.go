package route

import "net/http"

type patchMcontentReq postMcontentReq

func (m *patchMcontentReq) Bind(r *http.Request) error {
	return nil
}
