package main

import(
        "fmt"
        "os"
        "os/exec"
				"flag"
)
func main() {
	var Id string
	flag.Parse()
	args := flag.Args()
	if (len(args) > 1 || len(args) <= 0){
		fmt.Println("incorrect args")
		return
	}
	Id = args[0]
	c := exec.Command("curl", "localhost:8080/posts" + "/" + Id)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		fmt.Println(err)
	} 
}
