package main

import (
	"fmt"
	"os"
)

func main() {
	// Outputs your 'visual' env var (in my case 'vim' (set in ~/.bashrc))
	fmt.Println(os.Getenv("VISUAL"))

	// Create your own environment variable
	os.Setenv("SITE", "GoLangCode")
	fmt.Println(os.Getenv("SITE"))

	// Equals empty string if doesn't exist
	fmt.Println(os.Getenv("missing"))

}
