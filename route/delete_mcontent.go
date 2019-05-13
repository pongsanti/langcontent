package route

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/pongsanti/langcontent/db/models"
)

// DELETE /mcontents/{id}
func CreateDeleteMcontentHandlerFunc(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Print("Mcontent: DeleteMcontentHandlerFunc")

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
		).One(ctx, db)
		if err != nil {
			log.Print("Error select a mcontent ", err)
			renderError(errMcontentIdInvalid)
			return
		}

		// delete
		_, err = mcont.Delete(ctx, db)
		if err != nil {
			log.Print("Error delete mcontent ", err)
			renderError(errConnectingDatabase)
			return
		}

		render.JSON(w, req, struct{ Result string }{"done"})
	}
}
