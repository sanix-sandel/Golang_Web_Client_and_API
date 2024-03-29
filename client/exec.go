package client

import (
	"fmt"
	"net/http"
)

func DoOps(c *http.Client) error {
	resp, err := c.Get("http://www.google.com")

	if err != nil {
		return err
	}
	fmt.Println("Resulsts of DoOps:", resp.StatusCode)
	return nil
}

func DefaultGetGolang() error {
	resp, err := http.Get("https://www.golang.org")
	if err != nil {
		return err
	}
	fmt.Println("Results of DefaultGetGolang:", resp.StatusCode)
	return nil
}
