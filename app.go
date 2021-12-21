package bcrypt

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type IndexPageData struct {
	Date        string
	Title       string
	Count       string
	Difficulty  int64
	Strings     string
	Remember    string
	Action      string
	MoreStyles  []string
	MoreScripts []string
}

func Index(c *gin.Context) {

	session, err := store.Get(c.Request, cookieName)

	if err != nil {
		log.Println("Session error: ", err.Error())
	}

	data := new(IndexPageData)
	if c.Request.URL.Path != "/" {
		errorHandler(c, http.StatusNotFound)
		return
	}
	data.Date = time.Now().Format("2006")
	data.Title = siteName
	data.Count = "5"

	data.Strings = "Test1234!"
	data.Difficulty = 5

	var stringsTempVar string
	var difficultyTempVar int64
	stringsTempVar = mux.Vars(c.Request)["strings"]

	difficultyTempVar, _ = strconv.ParseInt(mux.Vars(c.Request)["difficulty"], 10, 64)

	if difficultyTempVar > 0 && difficultyTempVar <= 14 {
		data.Difficulty = difficultyTempVar
	} else if session.Values["difficulty"] != "" {
		if str, ok := session.Values["difficulty"].(string); ok {
			parsediff, err := strconv.ParseInt(str, 10, 64)
			if !(err != nil) && parsediff > 0 && parsediff <= 14 {
				data.Difficulty = int64(parsediff)
			}
		}
	}

	if session.Values["strings"] != "" {
		if str, ok := session.Values["strings"].(string); ok {
			data.Strings = str
		}
	}

	strings := c.Request.URL.Query().Get("strings")

	if len(strings) != 0 {
		data.Strings = strings
	}

	if stringsTempVar != "" {
		data.Strings = stringsTempVar
	}

	data.Remember = "false"
	if session.Values["remember"] != "" {
		if str, ok := session.Values["remember"].(string); ok {
			data.Remember = str
		}
	}

	err = session.Save(c.Request, c.Writer)

	if err != nil {
		log.Printf("Session save error: %v\n", err)
	}

	c.HTML(http.StatusOK, "app/index.html", data)
}
