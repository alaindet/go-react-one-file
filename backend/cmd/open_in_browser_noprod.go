//go:build !prod

package main

import "fmt"

func OpenURLInBrowser(url string) error {
	fmt.Println("Skip opening browser in development mode")
	return nil
}
