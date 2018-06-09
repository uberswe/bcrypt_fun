package main

import (
	"time"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

type IndexPageData struct {
	Date string
	Title string
	Count string
	Difficulty int64
	Strings string
	Remember string
	Action string
	MoreStyles []string
	MoreScripts []string
}

func Index(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, cookieName)

	if err != nil {
		log.Println("Session error: ", err.Error())
	}

	data := new(IndexPageData)
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound, err)
		return
	}
	data.Date = time.Now().Format("2006")
	data.Title = siteName
	data.Count = "5"

	data.Strings = "Test1234!"
	data.Difficulty = 5

	var stringsTempVar string
	var difficultyTempVar int64
	stringsTempVar = mux.Vars(r)["strings"]

	difficultyTempVar, _ = strconv.ParseInt(mux.Vars(r)["difficulty"], 10, 64)

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

	if stringsTempVar != "" {
		data.Strings = stringsTempVar
	}

	data.Remember = "false"
	if session.Values["remember"] != "" {
		if str, ok := session.Values["remember"].(string); ok {
			data.Remember = str
		}
	}

	log.Printf("Request received on: %s\n", r.URL.Path)

	err = session.Save(r, w)

	if err != nil {
		log.Printf("Session save error: %v\n", err)
	}

	err = tmpl.ExecuteTemplate(w, "index.html", data)

	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError, err)
		return
	}
}
