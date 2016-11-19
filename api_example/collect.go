package main

import(
        "fmt"
        "os"
        "os/exec"
)
func main() {
	c := exec.Command("curl", "localhost:8080/posts")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		fmt.Println(err)
	} 
}
