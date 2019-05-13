package route

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi"
	"github.com/pongsanti/dbconnect"
)

func TestPostMContent(t *testing.T) {

	dbConn, _ := dbconnect.NewDBConnect("localhost", "lib", "lib", "lib", "")
	postMcontent := CreatePostMcontentHandlerFunc(dbConn.Db)
	r := chi.NewRouter()
	r.Post("/mcontents", postMcontent)
	http.ListenAndServe(":3000", r)
}
