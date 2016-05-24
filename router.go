package datawell

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/csrf"
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
	"net/http"
	"strings"
	"encoding/json"
	"io/ioutil"
)

type Keys struct {
	csrf string
}

func init() {
	keys := Keys{}
	keysBytes, _ := ioutil.ReadFile("keys.json")
	json.Unmarshal(keysBytes, keys)

	r := mux.NewRouter()
	r.HandleFunc("/{path:.*}", handler)
	http.Handle("/", csrf.Protect([]byte(keys.csrf))(r))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	w.Header().Set("X-CSRF-Token", csrf.Token(r))
	vars := mux.Vars(r)
	path := vars["path"]
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	if u == nil {
		url, _ := user.LoginURL(ctx, path)
		fmt.Fprintf(w, `<p><a href="%s">Sign in or register</a> for '%s'</p>`, url, path)
		return
	}
	if strings.Contains(path, `penis`) {
		fmt.Fprint(w, `<p>Penis is Funny!</p>`)
	}
	url, _ := user.LogoutURL(ctx, path)
	fmt.Fprintf(w, `<p>Welcome, %s, to '%s'! (<a href="%s">sign out</a>)</p>`, u, path, url)
}
