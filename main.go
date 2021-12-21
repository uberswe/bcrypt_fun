package bcrypt

import (
	"embed"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"html/template"
	"io/fs"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

//go:embed assets/*
var static embed.FS

//go:embed views/*
var views embed.FS

var sessionExpiry = 3600 * 24 * 365 // 365 days
var cookieName = "BCRYPTFUN"
var siteUrl = "https://bcrypt.fun"
var siteName = "Bcrypt.fun"
var host = ":80"

var store *sessions.CookieStore

func init() {
	runtime.GOMAXPROCS(1)
	rand.Seed(time.Now().UTC().UnixNano())
	key := make([]byte, 32)

	_, err := rand.Read(key)
	if err != nil {
		panic("Could not generate key")
	}

	store = sessions.NewCookieStore(key)

	store.Options = &sessions.Options{
		Domain:   "bcrypt.fun",
		Path:     "/",
		MaxAge:   sessionExpiry,
		HttpOnly: true,
	}

	rand.Seed(time.Now().UnixNano())
}

func Run() {
	flag.Parse()
	flag.StringVar(&cookieName, "cookiename", cookieName, "The name to be used for cookies")
	flag.StringVar(&siteUrl, "siteurl", siteUrl, "The site url")
	flag.StringVar(&siteName, "sitename", siteName, "The name of the site")
	flag.StringVar(&host, "host", host, "The host and port (localhost:8005)")
	flag.IntVar(&sessionExpiry, "sessionexpiry", sessionExpiry, "Time in seconds that sessions should last (3600 * 24 * 365)")

	if envPort := os.Getenv("PORT"); envPort != "" {
		host = envPort
	}

	r := gin.Default()

	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)

	assets, err := fs.Sub(static, "assets")
	if err != nil {
		panic(err)
	}
	r.StaticFS("/assets", http.FS(assets))

	r.GET("/api/v1/hashes", RedirectToIndex)
	r.POST("/api/v1/hashes", Hashes)
	r.Any("/", Index)
	r.NoRoute(errorHandlerNoRoute)

	log.Printf("Running on %s\n", host)
	err = http.ListenAndServe(host, r)
	if err != nil {
		return
	}
}

func RedirectToIndex(c *gin.Context) {
	http.Redirect(c.Writer, c.Request, "/", 301)
}

func errorHandlerNoRoute(c *gin.Context) {
	errorHandler(c, http.StatusNotFound)
}

func errorHandler(c *gin.Context, status int) {
	if status == http.StatusNotFound {
		c.HTML(http.StatusNotFound, "errors/404.html", nil)
	} else {
		c.HTML(http.StatusInternalServerError, "errors/500.html", nil)
	}
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	err := fs.WalkDir(views, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		f, err := views.Open(path)
		if err != nil {
			return err
		}
		h, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		name := strings.Replace(path, "views/", "", 1)
		log.Println(name)
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return err
		}
		return nil
	})
	return t, err
}
