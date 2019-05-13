package route

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"

	"github.com/go-chi/chi"
	"github.com/pongsanti/dbconnect"
)

func TestPatchMContentText(t *testing.T) {

	dbConn, _ := dbconnect.NewDBConnect("localhost", "lib", "lib", "lib", "")
	boil.DebugMode = true

	patchMcontent := CreatePatchMcontentTextHandlerFunc(dbConn.Db)
	r := chi.NewRouter()
	r.Patch(fmt.Sprintf("/mcontent_texts/{%s}", PatchMcontentTextIDRouterParam), patchMcontent)
	http.ListenAndServe(":3000", r)
}
