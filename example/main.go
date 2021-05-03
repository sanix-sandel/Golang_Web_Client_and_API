package main

import "github.com/web_client/rest"

func main() {
	if err := rest.Exec(); err != nil {
		panic(err)
	}
}
