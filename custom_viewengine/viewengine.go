package custom_viewengine

import (
    "github.com/QLeelulu/goku"
)

type ThemeViewEngine struct {
    goku.DefaultViewEngine
}

func CreateThemeViewEngine(theme, viewDir, layout, extName string, useCache bool) *ThemeViewEngine {
    ve := &ThemeViewEngine{}
    ve.DefaultViewEngine = *goku.CreateDefaultViewEngine(viewDir, layout, extName, useCache)
    // default location paths
    ve.ViewLocationFormats = []string{
        theme + "/{1}/{0}",
        theme + "/shared/{0}",
    }
    ve.LayoutLocationFormats = []string{
        theme + "/{1}/{0}",
        theme + "/shared/{0}",
    }
    return ve
}
