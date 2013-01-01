package main

import (
    "github.com/QLeelulu/goku"
    "github.com/QLeelulu/goku-demo/custom_viewengine"
    _ "github.com/QLeelulu/goku-demo/custom_viewengine/controllers" // notice this!! import controllers
    "log"
)

func main() {
    rt := &goku.RouteTable{Routes: custom_viewengine.Routes}
    middlewares := []goku.Middlewarer{}
    s := goku.CreateServer(rt, middlewares, custom_viewengine.Config)
    goku.Logger().Logln("Server start on", s.Addr)
    log.Fatal(s.ListenAndServe())
}
