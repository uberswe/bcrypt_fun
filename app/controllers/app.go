package controllers

import (
	"github.com/revel/revel"
	"time"
	"strconv"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	var date string = time.Now().Format("2006")
	var title string = "Bcrypt.fun"
	var count string = "5"

	stringsVar := "Test1234!"
	difficultyVar := 5

	var stringsTempVar string
	var difficultyTempVar int
	c.Params.Query=c.Request.URL.Query()
	c.Params.Bind(&stringsTempVar,"strings")
	c.Params.Bind(&difficultyTempVar,"difficulty")

	if (c.Session["difficulty"] != "") {
		parsediff, err := strconv.ParseInt(c.Session["difficulty"], 10, 64)
		if !(err != nil) && parsediff > 0 && parsediff <= 14 {
			difficultyVar = int(parsediff)
		}
	}

	if (c.Session["strings"] != "") {
		stringsVar = c.Session["strings"]
	}

	if (stringsTempVar != "") {
		stringsVar = stringsTempVar
	}

	var remember string = "false"
	if (c.Session["remember"] != "") {
		remember = c.Session["remember"];
	}

	// Should be moved to new controller and all controllers inherit
	var action string = c.Action
	return c.Render(date, title, count, stringsVar, remember, difficultyVar, action)
}
