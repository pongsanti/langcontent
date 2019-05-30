package route

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/go-chi/render"
	"github.com/pongsanti/langcontent/db/models"
)

// POST /mcontents
func CreatePostMcontentHandlerFunc(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Print("Mcontent: PostMcontentHandlerFunc")

		ctx := req.Context()
		renderError := func(err error) {
			render.Status(req, 400)
			render.JSON(w, req, struct{ Error string }{
				err.Error(),
			})
		}

		// request binding
		reqData := &postMcontentReq{}
		err := render.Bind(req, reqData)
		if err != nil {
			renderError(err)
			return
		}

		mcont := models.Mcontent{
			Label:       reqData.Label,
			ContentType: reqData.ContentType,
			StartAt:     null.TimeFromPtr(reqData.StartAt),
			EndAt:       null.TimeFromPtr((reqData.EndAt)),
			Status:      null.StringFromPtr(reqData.Status),
			Xtime1:      null.TimeFromPtr(reqData.Xtime1),
			Xbool1:      null.BoolFromPtr(reqData.Xbool1),
		}

		// new tx
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			log.Print("Error creating tx ", err)
			renderError(errConnectingDatabase)
			return
		}

		err = mcont.Insert(ctx, tx, boil.Infer())
		if err != nil {
			log.Print("Error insert mcontent ", err)
			tx.Rollback()
			renderError(errConnectingDatabase)
			return
		}

		err = tx.Commit()
		if err != nil {
			log.Print("Error commit tx ", err)
			renderError(errConnectingDatabase)
			return
		}

		render.JSON(w, req, newPostMcontentRes(mcont))
	}
}
