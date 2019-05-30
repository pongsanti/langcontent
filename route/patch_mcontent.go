package route

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/pongsanti/langcontent/db/models"
)

const PatchMcontentIDRouterParam = "contentID"

// PATCH /mcontents/{contentID}
func CreatePatchMcontentHandlerFunc(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Print("Mcontent: PatchMcontentHandlerFunc")

		ctx := req.Context()
		renderError := func(err error) {
			render.JSON(w, req, struct{ Error string }{
				err.Error(),
			})
		}

		// get mcontent id
		mcontIDString := chi.URLParam(req, PatchMcontentIDRouterParam)
		mcontID, err := strconv.Atoi(mcontIDString)
		if err != nil {
			log.Print("Error convert mcontent id ", err)
			renderError(errMcontentIdInvalid)
			return
		}

		// find mcontent
		mcont, err := models.Mcontents(
			models.McontentWhere.ID.EQ(mcontID),
			qm.Load(models.McontentRels.McontentTexts,
				qm.OrderBy("lang asc, id desc")),
		).One(ctx, db)
		if err != nil {
			renderError(errMcontentIdInvalid)
			return
		}

		// request binding
		reqData := &patchMcontentReq{}
		err = render.Bind(req, reqData)
		if err != nil {
			renderError(err)
			return
		}

		// update model
		if reqData.Label != nil {
			mcont.Label = *reqData.Label
		}
		if reqData.StartAt != nil {
			mcont.StartAt = null.TimeFromPtr(reqData.StartAt)
		}
		if reqData.EndAt != nil {
			mcont.EndAt = null.TimeFromPtr(reqData.EndAt)
		}
		if reqData.Status != nil {
			mcont.Status = null.StringFromPtr(reqData.Status)
		}
		if reqData.Xtime1 != nil {
			mcont.Xtime1 = null.TimeFromPtr(reqData.Xtime1)
		}
		if reqData.Xbool1 != nil {
			mcont.Xbool1 = null.BoolFromPtr(reqData.Xbool1)
		}
		if reqData.DeletedAt.Set {
			if !reqData.DeletedAt.Valid {
				mcont.DeletedAt = null.TimeFromPtr(nil)
			} else {
				mcont.DeletedAt = null.TimeFrom(reqData.DeletedAt.Value)
			}
		}

		_, err = mcont.Update(ctx, db, boil.Whitelist(
			models.McontentColumns.Label,
			models.McontentColumns.StartAt,
			models.McontentColumns.EndAt,
			models.McontentColumns.Status,
			models.McontentColumns.Xtime1,
			models.McontentColumns.Xbool1,
			models.McontentColumns.DeletedAt,
		))
		if err != nil {
			log.Print("Error updating an mcontent ", err)
			renderError(errConnectingDatabase)
			return
		}

		render.JSON(w, req, newPatchMcontentRes(*mcont))
	}
}
