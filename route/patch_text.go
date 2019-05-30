package route

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/volatiletech/sqlboiler/boil"

	"github.com/pongsanti/langcontent/db/models"
	"github.com/volatiletech/null"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

const PatchMcontentTextIDRouterParam = "textID"

// PATCH /mcontent_texts/{textID}
func CreatePatchMcontentTextHandlerFunc(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Print("Mcontent: PatchMcontentTextHandlerFunc")

		ctx := req.Context()
		renderError := func(err error) {
			render.Status(req, 400)
			render.JSON(w, req, struct{ Error string }{
				err.Error(),
			})
		}

		// get text id
		textIDString := chi.URLParam(req, PatchMcontentTextIDRouterParam)
		textID, err := strconv.Atoi(textIDString)
		if err != nil {
			log.Print("Error convert mcontent id ", err)
			renderError(errMcontentTextIdInvalid)
			return
		}

		// find text by id
		text, err := models.FindMcontentText(ctx, db, textID)
		if err != nil {
			log.Print("Error find mcontent text ", err)
			renderError(errMcontentTextIdInvalid)
			return
		}

		reqData := &patchTextReq{}
		_ = render.Bind(req, reqData)

		if reqData.Title != nil {
			text.Title = *reqData.Title
		}
		if reqData.Subtitle != nil {
			text.Subtitle = null.StringFromPtr(reqData.Subtitle)
		}
		if reqData.Detail != nil {
			text.Detail = null.StringFromPtr(reqData.Detail)
		}
		if reqData.Xtext1 != nil {
			text.Xtext1 = null.StringFromPtr(reqData.Xtext1)
		}

		_, err = text.Update(ctx, db, boil.Whitelist(
			models.McontentTextColumns.Title,
			models.McontentTextColumns.Subtitle,
			models.McontentTextColumns.Detail,
			models.McontentTextColumns.Xtext1,
		))

		if err != nil {
			log.Print("Error update mcontent text ", err)
			renderError(errConnectingDatabase)
			return
		}

		render.JSON(w, req, newPatchTextRes(*text))
	}
}
