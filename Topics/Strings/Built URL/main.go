package main

import "fmt"

func main() {
	// DO NOT modify the code block below!
	var protocol, domain, path string
	fmt.Scanln(&protocol, &domain, &path)

	// Concatenate the 'protocol', 'domain', and 'path' variables into the 'completeURL' variable below:
	// Use the format: protocol://domain/path
	completeURL := protocol + "://" + domain + "/" + path

	fmt.Println(completeURL) // DO NOT delete, this prints the completeURL!
}
