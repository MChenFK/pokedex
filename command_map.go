package main

import (
	"fmt"
	"io"
	"net/http"
)

func commandMap() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}
	fmt.Printf("%s", body)
	return nil
}
