package route

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"

	"github.com/go-chi/chi"
	"github.com/pongsanti/dbconnect"
)

func TestPatchMContent(t *testing.T) {

	dbConn, _ := dbconnect.NewDBConnect("localhost", "lib", "lib", "lib", "")
	boil.DebugMode = true

	patchMcontent := CreatePatchMcontentHandlerFunc(dbConn.Db)
	r := chi.NewRouter()
	r.Patch(fmt.Sprintf("/mcontents/{%s}", PatchMcontentIDRouterParam), patchMcontent)
	http.ListenAndServe(":3000", r)
}
