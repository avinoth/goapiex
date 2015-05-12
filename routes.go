package main

import "net/http"

type Route struct {
    Name                 string
    Method               string
    Pattern               string
    HandlerFunc       http.HandlerFunc
}

type Routes []Route

var routes = Routes {
    Route{
        "Home",
        "GET",
        "/",
        home,
    },
    Route{
        "Index",
        "GET",
        "/todos",
        index,
    },
    Route{
        "Show",
        "GET",
        "/todos/{todoId}",
        show,
    },
    Route{
        "Create",
        "POST",
        "/todos",
        create,
    },
}
