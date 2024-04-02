package main

import "fmt"

//import "fmt"

func main() {
	//websites := []string{"http://google.com", "http://aws.com"}
	websites := map[string]string{
		"Google": "http://google.com",
		"AWS":    "http://aws.com",
	}
	fmt.Println("AWS provider:", websites["AWS"])
	websites["GCP"] = "https://gcp.com"
	fmt.Println(websites)
	delete(websites, "AWS")
	fmt.Println(websites)
}
