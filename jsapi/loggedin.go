package jsapi

import (
	"github.com/ant0ine/go-json-rest/rest"
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
	"net/http"
	"strconv"
)

type LoggedInMiddleware struct{}

func (mw *LoggedInMiddleware) MiddlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {

	return func(w rest.ResponseWriter, r *rest.Request) {
		ctx := appengine.NewContext(r.Request)
		u := user.Current(ctx)
		if u == nil {
			url, _ := user.LoginURL(ctx, r.Referer())
			w.WriteHeader(http.StatusUnauthorized)
			w.WriteJson(ErrorResponse{
				Errors: []Error{
					Error{
						Status: strconv.Itoa(http.StatusUnauthorized),
						Title: "Unauthorized",
					},
				},
				Meta: map[string]interface{}{
					"login_url": url,
				},
			})
			return
		}
		handler(w, r)
	}
}
