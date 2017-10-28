package controllers

import (
	"github.com/revel/revel"
	"net/url"
	"time"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type Api struct {
	*revel.Controller
}

type Hash struct {
	Hash string ` json:"hash" xml:"hash" `
}

type Params struct {
	url.Values
}

func (c Api) Index() revel.Result {
	// Show some documentation here
	var action string = c.Action
	var date string = time.Now().Format("2006")
	var title string = "Bcrypt.fun"
	return c.Render(action, title, date)
}

func (c Api) Hashes() revel.Result {
	var paramStrings string
	c.Params.Bind(&paramStrings, "strings") // Sets the number of passwords

	c.Session["strings"] = paramStrings;

	data := make(map[string]interface{})

	stringArray := strings.Split(paramStrings,"\n")
	hashes := []Hash{}
	for _, str := range stringArray {
		hash, _ := HashPassword(str)
		hashes = append(hashes, Hash{Hash:hash})
	}

	data["href"] = "https://bcrypt.fun" + "/api/v1/hashes"
	data["hashes"] = hashes
	data["count"] = len(hashes)
	return c.RenderJSON(data)
	//return c.RenderXML(data)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	return string(bytes), err
}

