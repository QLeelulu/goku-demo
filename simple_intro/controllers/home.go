package controllers

import (
    "github.com/QLeelulu/goku"
)

// home controller
var _ = goku.Controller("home").
// index action
Get("index", func (ctx *goku.HttpContext) goku.ActionResulter {
	ctx.ViewData["Message"] = "Hello World"
    return ctx.View(nil)
})
