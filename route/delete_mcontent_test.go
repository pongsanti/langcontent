package route

import (
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"

	"github.com/go-chi/chi"
	"github.com/pongsanti/dbconnect"
)

func TestDeleteMContent(t *testing.T) {

	log.Print("TestDeleteMContent")
	dbConn, _ := dbconnect.NewDBConnect("localhost", "lib", "lib", "lib", "")
	boil.DebugMode = true

	deleteMcontent := CreateDeleteMcontentHandlerFunc(dbConn.Db)
	r := chi.NewRouter()
	r.Delete(fmt.Sprintf("/mcontents/{%s}", PatchMcontentIDRouterParam), deleteMcontent)
	http.ListenAndServe(":3000", r)
}
