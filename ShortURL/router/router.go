package router

import (
	handler "ShortURL/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

/*
	Create the routes for the API. The API supports three URLs:
		1- GET "/" => Shows a description for the API
		2- GET "/{shorturl}" => If the shortUrl exists in the backend database, redirect to the long url that corresponds to func init() {
		3- Post "/Create" => Takes a post request with http body of {
																	shorturl: "short Link"
																	longurl:  "original long link"
																	}
		 Causes the API to create a mapping between the short url and the long url in the backend database
*/

func CreateRoutes(LS *handler.LinkShortnerAPI) Routes {
	return Routes{
		Route{
			"UrlRoot",
			"GET",
			"/",
			LS.UrlRoot,
		},
		Route{
			"UrlShow",
			"GET",
			"/{shorturl}",
			LS.UrlShow,
		},
		Route{
			"UrlCreate",
			"POST",
			"/Create",
			LS.UrlCreate,
		},
	}
}

func NewLinkShortenerRouter(routes Routes) *mux.Router {
	// Если StrictSlash установлен на true, то при роутере "/path/"
	// автоматически будет редиректить на "/path"
	router := mux.NewRouter().StrictSlash(true)
	// указываем всю необходимую информаци для правильной работы роутера
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
