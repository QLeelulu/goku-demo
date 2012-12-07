package gokudemo

import (
    "github.com/QLeelulu/goku"
)

// routes
var Routes []*goku.Route = []*goku.Route{
    &goku.Route{
        Name:     "static",
        IsStatic: true,
        Pattern:  "/public/(.*)",
    },
    &goku.Route{
        Name:    "default",
        Pattern: "/{controller}/{action}",
        Default: map[string]string{"controller": "home", "action": "index"},
    },
}
