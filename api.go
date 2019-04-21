package main

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Hash struct {
	Hash string ` json:"hash" xml:"hash" `
}

type Params struct {
	url.Values
}

func Hashes(w http.ResponseWriter, r *http.Request) {

	// TODO rate limit this?

	session, err := store.Get(r, cookieName)

	if err != nil {
		log.Println("Session error: ", err.Error())
	}

	var paramStrings string
	var paramRemember string
	var difficultyVar int
	var remember = false

	paramStrings = r.FormValue("strings")   // Sets the number of passwords
	paramRemember = r.FormValue("remember") // Store values in session cookie

	if paramRemember == "on" {
		remember = true
	}

	i64Tmp, err := strconv.ParseInt(r.FormValue("difficulty"), 10, 64)

	if err != nil {
		difficultyVar = 0
	} else {
		for a := 0; a <= int(i64Tmp); a++ {
			difficultyVar = a
		}
	}

	if difficultyVar < 0 {
		difficultyVar = 1
	}

	if difficultyVar >= 14 {
		difficultyVar = 14
	}

	parsediff := strconv.FormatInt(int64(difficultyVar), 10)

	if remember {
		session.Values["strings"] = paramStrings
		session.Values["remember"] = boolToString(remember)
		session.Values["difficulty"] = parsediff
	} else {
		delete(session.Values, "strings")
		delete(session.Values, "remember")
		delete(session.Values, "difficulty")
	}

	data := make(map[string]interface{})

	stringArray := strings.Split(paramStrings, "\n")
	var hashes []Hash
	limit := 0
	for _, str := range stringArray {
		str = strings.TrimSuffix(str, "\n")
		str = strings.TrimSuffix(str, "\r")
		hash, _ := HashPassword(str, difficultyVar)
		hashes = append(hashes, Hash{Hash: hash})
		limit++
		// Limit to 20 hashes
		if limit >= 20 {
			break
		}
		if difficultyVar >= 10 {
			break
		}
	}

	err = session.Save(r, w)

	if err != nil {
		log.Printf("Session save error: %v\n", err)
	}

	data["href"] = siteUrl + "/api/v1/hashes"
	data["hashes"] = hashes
	data["count"] = len(hashes)

	payload, err := json.Marshal(data)
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError, err)
		return
	}

	log.Printf("Request received on: %s\n", r.URL.Path)

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func HashPassword(password string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

func boolToString(b bool) string {
	return strconv.FormatBool(b)
}
