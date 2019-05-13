package route

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/volatiletech/null"

	"github.com/go-chi/chi"

	"github.com/go-chi/render"
	"github.com/pongsanti/langcontent/db/models"
)

const McontentIDRouterParam = "contentID"

func CreatePostMcontentTextHandlerFunc(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Print("Mcontent: PostMcontentTextHandlerFunc")

		ctx := req.Context()
		renderError := func(err error) {
			render.JSON(w, req, struct{ Error string }{
				err.Error(),
			})
		}

		// get mcontent id
		mcontIDString := chi.URLParam(req, McontentIDRouterParam)
		mcontID, err := strconv.Atoi(mcontIDString)
		if err != nil {
			log.Print("Error convert mcontent id ", err)
			renderError(errMcontentIdInvalid)
			return
		}

		// request binding
		reqData := &postTextReq{}
		err = render.Bind(req, reqData)
		if err != nil {
			renderError(err)
			return
		}

		text := models.McontentText{
			Lang:     reqData.Lang,
			Title:    reqData.Title,
			Subtitle: null.StringFromPtr(reqData.Subtitle),
			Detail:   null.StringFromPtr(reqData.Detail),
			Xtext1:   null.StringFromPtr(reqData.Xtext1),
		}

		// find mcontent
		mcont, err := models.FindMcontent(ctx, db, mcontID)
		if err != nil {
			renderError(errMcontentIdInvalid)
			return
		}

		// associate text to mcontent
		err = mcont.AddMcontentTexts(ctx, db, true, &text)

		render.JSON(w, req, newPostTextRes(text))
	}
}
