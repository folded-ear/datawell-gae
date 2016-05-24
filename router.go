package datawell

import (
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
	"net/http"
	"strings"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/{path:.*}", handler)
	http.Handle("/", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	path := vars["path"]
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	if strings.Contains(path, `penis`) {
		fmt.Fprint(w, `<p>Penis is Funny!</p>`)
	}
	if u == nil {
		url, _ := user.LoginURL(ctx, path)
		fmt.Fprintf(w, `<p><a href="%s">Sign in or register</a> for '%s'</p>`, url, path)
		return
	}
	url, _ := user.LogoutURL(ctx, path)
	fmt.Fprintf(w, `<p>Welcome, %s, to '%s'! (<a href="%s">sign out</a>)</p>`, u, path, url)
}
