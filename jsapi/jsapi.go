package jsapi

import (
	"github.com/ant0ine/go-json-rest/rest"
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
	"log"
	"net/http"
	"strings"
)

const PATH_PREFIX = "/jsapi"

func init() {
	api := rest.NewApi()
	api.Use(
		&LoggedInMiddleware{},
		&rest.RecoverMiddleware{
			EnableResponseStackTrace: true,
		},
		&rest.ContentTypeCheckerMiddleware{},
	)
	router, err := rest.MakeRouter(
		rest.Get(   "/current_user",    getCurrentUser),
		rest.Get(   "/events",          getEvents),
		rest.Post(  "/events",          postEvents),
		rest.Get(   "/events/:id",      getEvent),
		rest.Patch( "/events/:id",      patchEvent),
		rest.Delete("/events/:id",      deleteEvent),
		rest.Get(   "/tags",            getTags),
		rest.Get(   "/tagsets",         getTagsets),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	http.Handle("/", http.StripPrefix(PATH_PREFIX, api.MakeHandler()))
}

func getTags(w rest.ResponseWriter, r *rest.Request) {
	panic("not implemented")
}

func getTagsets(w rest.ResponseWriter, r *rest.Request) {
	panic("not implemented")
}

func deleteEvent(w rest.ResponseWriter, r *rest.Request) {
	panic("not implemented")
}

func patchEvent(w rest.ResponseWriter, r *rest.Request) {
	panic("not implemented")
}

func getEvent(w rest.ResponseWriter, r *rest.Request) {
	panic("not implemented")
}

func postEvents(w rest.ResponseWriter, r *rest.Request) {
	panic("not implemented")
}

func getEvents(w rest.ResponseWriter, r *rest.Request) {
	panic("not implemented")
}

func getCurrentUser(w rest.ResponseWriter, r *rest.Request) {
	ctx := appengine.NewContext(r.Request)
	u := user.Current(ctx)
	logoutUrl, _ := user.LogoutURL(ctx, r.Referer())
	w.WriteJson(RecordResponse{
		Data: Data{
			Type: "user",
			Id: u.ID,
			Attributes: map[string]interface{}{
				"email": u.Email,
				"name": strings.Split(u.Email, "@")[0],
			},
		},
		Meta: map[string]interface{}{
			"logout_url": logoutUrl,
		},
	})
}
