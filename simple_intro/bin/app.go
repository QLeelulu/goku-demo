package main

import (
    "github.com/QLeelulu/goku"
    "github.com/QLeelulu/goku-demo/simple_intro"
    _ "github.com/QLeelulu/goku-demo/simple_intro/controllers" // notice this!! import controllers
    "log"
)

func main() {
    rt := &goku.RouteTable{Routes: simple_intro.Routes}
    middlewares := []goku.Middlewarer{}
    s := goku.CreateServer(rt, middlewares, simple_intro.Config)
    goku.Logger().Logln("Server start on", s.Addr)
    log.Fatal(s.ListenAndServe())
}
