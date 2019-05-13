package route

import "net/http"

type postTextReq struct {
	Lang     string
	Title    string
	Subtitle *string
	Detail   *string
	Xtext1   *string
}

func (p *postTextReq) Bind(r *http.Request) error {
	if p.Lang == "" {
		return errLanguageRequired
	}

	if p.Title == "" {
		return errTitleRequired
	}
	return nil
}
