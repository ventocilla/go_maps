package main

import "fmt"

//import "fmt"

func main() {
	//websites := []string{"http://google.com", "http://aws.com"}
	websites := map[string]string{
		"Google": "http://google.com",
		"AWS":    "http://aws.com",
	}
	fmt.Println(websites)
}
