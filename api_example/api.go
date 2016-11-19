package main

import(
	"github.com/go-web-framework/gflux/api"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"encoding/json"
	"bytes"
)

type Post struct {
	Body string
}

func main(){

	newPost2 := Post{Body: "what"}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(newPost2)
	fmt.Println(b)

	a := api.New("sqlite3", "api.db")
	a.NewResource("posts", &Post{})
	a.DB.Insert(newPost2, "posts")
	a.Serve()
}
