package route

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/volatiletech/null"

	"github.com/go-chi/chi"

	"github.com/go-chi/render"
	"github.com/pongsanti/langcontent/db/models"
)

const McontentIDRouterParam = "contentID"

// POST /mcontents/{contentID}/texts
func CreatePostMcontentTextHandlerFunc(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Print("Mcontent: PostMcontentTextHandlerFunc")

		ctx := req.Context()
		renderError := func(err error) {
			render.Status(req, 400)
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

		// new tx
		tx, err := db.Begin()
		if err != nil {
			log.Print("Error create new tx ", err)
			renderError(errConnectingDatabase)
			return
		}

		// delete existing texts of the language we will insert
		_, err = models.McontentTexts(
			models.McontentTextWhere.McontentID.EQ(mcontID),
			models.McontentTextWhere.Lang.EQ(reqData.Lang),
		).DeleteAll(ctx, tx)
		if err != nil {
			log.Print("Error delete texts ", err)
			tx.Rollback()
			renderError(errConnectingDatabase)
			return
		}

		// find mcontent
		mcont, err := models.Mcontents(
			models.McontentWhere.ID.EQ(mcontID),
			models.McontentWhere.DeletedAt.IsNull(),
			qm.Load(models.McontentRels.McontentTexts),
		).One(ctx, tx)
		if err != nil {
			log.Print("Error find mcont ", err)
			tx.Rollback()
			renderError(errMcontentIdInvalid)
			return
		}

		// associate text to mcontent
		err = mcont.AddMcontentTexts(ctx, db, true, &text)
		if err != nil {
			log.Print("Error associate text with mcontent ", err)
			tx.Rollback()
			renderError(errConnectingDatabase)
			return
		}

		// commit tx
		if err = tx.Commit(); err != nil {
			log.Print("Error commit tx ", err)
			tx.Rollback()
			renderError(errConnectingDatabase)
			return
		}

		render.JSON(w, req, newPostTextRes(text))
	}
}
