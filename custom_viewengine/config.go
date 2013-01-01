package custom_viewengine

import (
    "github.com/QLeelulu/goku"
    "path"
    "runtime"
    "time"
)

var Config *goku.ServerConfig = &goku.ServerConfig{
    Addr:           ":8080",
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
    //RootDir:        _, filename, _, _ := runtime.Caller(1),
    StaticPath: "static", // static content 
    ViewPath:   "views",

    LogLevel: goku.LOG_LEVEL_LOG,
    Debug:    true,
}

func init() {
    // set the RootDir as current dir.
    _, filename, _, _ := runtime.Caller(1)
    Config.RootDir = path.Dir(filename)

    // change view engine
    theme := "blue"
    viewDir := path.Join(Config.RootDir, Config.ViewPath)
    layout := "layout"
    extName := ".html"
    useCache := !Config.Debug
    Config.ViewEnginer = CreateThemeViewEngine(theme, viewDir, layout, extName, useCache)
}
