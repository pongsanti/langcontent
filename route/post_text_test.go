package route

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-chi/chi"
	"github.com/pongsanti/dbconnect"
)

func TestPostText(t *testing.T) {
	dbConn, _ := dbconnect.NewDBConnect("localhost", "lib", "lib", "lib", "")
	postText := CreatePostMcontentTextHandlerFunc(dbConn.Db)
	r := chi.NewRouter()
	r.Post(fmt.Sprintf("/mcontents/{%s}/texts", McontentIDRouterParam), postText)
	http.ListenAndServe(":3000", r)
}
