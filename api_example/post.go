package main

import(
        "fmt"
        "os"
        "os/exec"
				"flag"
)

type Post struct {
	Body string
}

func main() {
	var Id string
	flag.Parse()
	args := flag.Args()
	if (len(args) > 1 || len(args) <= 0){
		fmt.Println("incorrect args")
		return
	}
	Id = args[0]
	
	jsonPost := "{\"Body\":\"" + Id + "\"}"

	c := exec.Command("curl","-H", "Content-Type: application/json", "-X", "POST", "-d", jsonPost, "localhost:8080/posts")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		fmt.Println(err)
	} 
}
