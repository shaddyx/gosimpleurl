package main

import (
	"fmt"

	"./parserlib/simpleurl"
)

func main() {
	a := simpleurl.NewSimpleUrl()
	data, err := a.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
	}
	if a.Proxy == "" {
		fmt.Println("Niiiiiiiiiiiiiiiiil")
	}
	fmt.Println(data)
}
