package controllers

import (
	"github.com/revel/revel"
	"time"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	var date string = time.Now().Format("2006")
	var title string = "Bcrypt.fun"
	var count string = "5"
	if (c.Session["count"] != "") {
		count = c.Session["count"]
	}
	var length string = "25"
	if (c.Session["length"] != "") {
		length = c.Session["length"]
	}
	var capital string = "true"
	if (c.Session["capital"] != "") {
		capital = c.Session["capital"];
	}
	var lower string = "true"
	if (c.Session["lower"] != "") {
		lower = c.Session["lower"];
	}
	var special string = "true"
	if (c.Session["special"] != "") {
		special = c.Session["special"];
	}
	var spaces string = "false"
	if (c.Session["spaces"] != "") {
		spaces = c.Session["spaces"];
	}
	var numbers string = "true"
	if (c.Session["numbers"] != "") {
		numbers = c.Session["numbers"];
	}
	var highlight string = "false"
	if (c.Session["highlight"] != "") {
		highlight = c.Session["highlight"];
	}
	var remember string = "false"
	if (c.Session["remember"] != "") {
		remember = c.Session["remember"];
	}

	// Should be moved to new controller and all controllers inherit
	var action string = c.Action
	return c.Render(date, title, count, length, capital, lower, special, spaces, numbers, highlight, remember, action)
}
