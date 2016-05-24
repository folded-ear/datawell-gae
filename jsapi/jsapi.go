package jsapi

import (
	"github.com/ant0ine/go-json-rest/rest"
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
	"log"
	"net/http"
)

const PATH_PREFIX = "/jsapi"

func init() {
	api := rest.NewApi()
	api.Use(
		&LoggedInMiddleware{},
		&rest.RecoverMiddleware{
			EnableResponseStackTrace: true,
		},
		&rest.JsonIndentMiddleware{},
		&rest.ContentTypeCheckerMiddleware{},
	)
	router, err := rest.MakeRouter(
		rest.Get("/#path", func(w rest.ResponseWriter, r *rest.Request) {
			path := r.PathParam("path")
			ctx := appengine.NewContext(r.Request)
			u := user.Current(ctx)
			url, _ := user.LogoutURL(ctx, PATH_PREFIX+r.URL.Path)
			w.WriteJson(map[string]interface{}{
				"user":      u,
				"LogoutUrl": url,
				"path":      path,
			})
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	http.Handle("/", http.StripPrefix(PATH_PREFIX, api.MakeHandler()))
}
