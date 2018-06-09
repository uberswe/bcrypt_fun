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
	"flag"
)

var sessionExpiry = 3600 * 24 * 365  // 365 days
var cookieName = "BCRYPTFUN"
var siteUrl = "https://bcrypt.fun"
var siteName = "Bcrypt.fun"
var host = ":8005"
var tmpl = ParseTemplates()

var store *sessions.CookieStore

func init() {
	key := make([]byte, 32)

	_, err := rand.Read(key)
	if err != nil {
		panic("Could not generate key")
	}

	store = sessions.NewCookieStore(key)

	store.Options = &sessions.Options{
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   sessionExpiry,
		HttpOnly: true,
	}

	rand.Seed(time.Now().UnixNano())
}

func main() {
	flag.StringVar(&cookieName, "cookiename", cookieName, "The name to be used for cookies")
	flag.StringVar(&siteUrl, "siteurl", siteUrl, "The site url")
	flag.StringVar(&siteName, "sitename", siteName, "The name of the site")
	flag.StringVar(&host, "host", host, "The host and port (localhost:8005)")
	flag.IntVar(&sessionExpiry, "sessionecpiry", sessionExpiry, "Time in seconds that sessions should last (3600 * 24 * 365)")

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