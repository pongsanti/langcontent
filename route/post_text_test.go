package route

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"

	"github.com/go-chi/chi"
	"github.com/pongsanti/dbconnect"
)

func TestPostText(t *testing.T) {
	dbConn, _ := dbconnect.NewDBConnect("localhost", "lib", "lib", "lib", "")
	boil.DebugMode = true

	postText := CreatePostMcontentTextHandlerFunc(dbConn.Db)
	r := chi.NewRouter()
	r.Post(fmt.Sprintf("/mcontents/{%s}/texts", McontentIDRouterParam), postText)
	http.ListenAndServe(":3000", r)
}
