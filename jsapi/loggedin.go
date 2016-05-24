package jsapi

import (
	"github.com/ant0ine/go-json-rest/rest"
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
	"net/http"
)

type LoggedInMiddleware struct{}

func (mw *LoggedInMiddleware) MiddlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {

	return func(w rest.ResponseWriter, r *rest.Request) {
		ctx := appengine.NewContext(r.Request)
		u := user.Current(ctx)
		if u == nil {
			url, _ := user.LoginURL(ctx, PATH_PREFIX+r.URL.Path)
			w.WriteHeader(http.StatusUnauthorized)
			w.WriteJson(map[string]string{
				"Error":    "Login required",
				"LoginUrl": url,
			})
			return
		}
		handler(w, r)
	}
}
