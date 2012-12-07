package main

import (
    "github.com/QLeelulu/goku"
    "github.com/QLeelulu/goku-demo/simple-intro"
    _ "github.com/QLeelulu/goku-demo/simple-intro/controllers" // notice this!! import controllers
    "log"
)

func main() {
    rt := &goku.RouteTable{Routes: gokudemo.Routes}
    middlewares := []goku.Middlewarer{}
    s := goku.CreateServer(rt, middlewares, gokudemo.Config)
    goku.Logger().Logln("Server start on", s.Addr)
    log.Fatal(s.ListenAndServe())
}
