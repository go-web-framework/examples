// gflux/api allows programmers to create a REST API easily.
// FUTURE FEATURE : This tool automatically creates html pages to get and set

package main

import (
	"fmt"
	"github.com/go-web-framework/gflux/api"
	"net/http"
)

type Post struct {
	api.Model
	Title string
	Author string
}

type User struct {
	api.Model
	Name string
}

func main(){
	// create new REST API that uses sqlite3 database database.db as the datastore
	a := api.New("sqlite3", "database.db")
	
	// create a new resource posts which is a new table in database.db
	rposts := api.NewResource("posts", &Post{})
	
	// create resource users which allows GET and POST requests
	rusers := api.NewResource("users", &User{}).allow("GET", "POST")
	
	// OPTIONAL: override GET handler. Same can be done for POST handler.
	// Default handler only handles application/json and application/xml
	rusers.Handlers["GET"] = func(user User, w http.responseWriter, accepts []string){
		// IMPLEMENTATION that writes to w. accepts looks like ["html", "json"]
	}
	
	a.Serve()
}
	
