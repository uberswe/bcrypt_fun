package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"html/template"
	"github.com/gorilla/sessions"
	"math/rand"
	"time"
	"path/filepath"
	"os"
	"strings"
	"log"
)

// port 8005?
// cookie prefix BCRYPTFUN
// session expires = 720h
var cookieName = "BCRYPTFUN"
var siteUrl = "https://bcrypt.fun"
var siteName = "Bcrypt.fun"
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var store = sessions.NewCookieStore([]byte(RandStringRunes(50)))
var tmpl = ParseTemplates()

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/favicon.ico").Handler(http.FileServer(http.Dir("./assets/favicon.ico")))
	r.PathPrefix("/favicon.png").Handler(http.FileServer(http.Dir("./assets/img/favicon.png")))
	r.PathPrefix("/apple-touch-icon.png").Handler(http.FileServer(http.Dir("./assets/apple-touch-icon.png")))
	r.PathPrefix("/browserconfig.xml").Handler(http.FileServer(http.Dir("./assets/browserconfig.xml")))
	r.PathPrefix("/crossdomain.xml").Handler(http.FileServer(http.Dir("./assets/crossdomain.xml")))
	r.PathPrefix("/humans.txt").Handler(http.FileServer(http.Dir("./assets/humans.txt")))
	r.PathPrefix("/robots.txt").Handler(http.FileServer(http.Dir("./assets/robots.txt")))
	r.PathPrefix("/tile.png").Handler(http.FileServer(http.Dir("./assets/tile.png")))
	r.PathPrefix("/tile-wide.png").Handler(http.FileServer(http.Dir("./assets/tile-wide.png")))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	r.HandleFunc("/api/v1/hashes", RedirectToIndex).Methods("GET")
	r.HandleFunc("/api/v1/hashes", Hashes).Methods("POST")
	r.HandleFunc("/", Index)
	log.Printf("Running on :8005\n")
	http.ListenAndServe(":8005", r)
}

func RedirectToIndex (w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", 301)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int, err error) {
	w.WriteHeader(status)

	log.Printf("Error on: %s %v - %v\n", r.URL.Path, status, err)
	if status == http.StatusNotFound {
		tmpl.ExecuteTemplate(w, "404",nil)
		tmpl.Execute(w, nil)
	} else if status == http.StatusInternalServerError {
		tmpl.ExecuteTemplate(w, "500",nil)
	}
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func ParseTemplates() *template.Template {
	log.Printf("Building template files\n")
	templ := template.New("")
	err := filepath.Walk("./views", func(path string, info os.FileInfo, err error) error {
		log.Printf("Parsing: %s\n", path)
		if strings.Contains(path, ".html") {
			_, err = templ.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}

		return err
	})

	if err != nil {
		panic(err)
	}
	log.Printf("Templates ready\n")

	return templ
}