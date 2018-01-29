package controllers

import (
	"github.com/revel/revel"
	"net/url"
	"time"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"strconv"
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
	var paramRemember bool
	var difficultyVar int

	c.Params.Bind(&paramStrings, "strings") // Sets the number of passwords
	c.Params.Bind(&paramRemember, "remember") // Store values in session cookie
	c.Params.Bind(&difficultyVar,"difficulty")

	if difficultyVar < 0 {
		difficultyVar = 1
	}

	if difficultyVar >= 14 {
		difficultyVar = 14
	}

	parsediff := strconv.FormatInt(int64(difficultyVar), 10)

	if (paramRemember) {
		c.Session["strings"] = paramStrings;
		c.Session["remember"] = boolToString(paramRemember);
		c.Session["difficulty"] = parsediff;
	} else {
		delete(c.Session, "strings")
		delete(c.Session, "remember")
		delete(c.Session, "difficulty")
	}

	data := make(map[string]interface{})

	stringArray := strings.Split(paramStrings,"\n")
	hashes := []Hash{}
	limit := 0
	for _, str := range stringArray {
		hash, _ := HashPassword(str, difficultyVar)
		hashes = append(hashes, Hash{Hash:hash})
		limit++
		// Limit to 20 hashes
		if limit >= 20 {
			break
		}
		if difficultyVar >= 10 {
			break;
		}
	}

	data["href"] = "https://bcrypt.fun" + "/api/v1/hashes"
	data["hashes"] = hashes
	data["count"] = len(hashes)
	return c.RenderJSON(data)
	//return c.RenderXML(data)
}

func HashPassword(password string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

func boolToString(b bool) string {
	return strconv.FormatBool(b)
}

