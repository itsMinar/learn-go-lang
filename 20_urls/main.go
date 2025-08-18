package main

import (
	"fmt"
	"net/url"
)

const my_url string = "https://lco.dev:3000/learn?coursename=golang&paymentid=12345"

func main() {
	fmt.Println("Welcome to handling URLs in Go!")
	fmt.Println("URL: ", my_url)

	// parsing the URL
	parsedUrl, err := url.Parse(my_url)
	if err != nil {
		panic(err)
	}

	// fmt.Println("scheme: ", parsedUrl.Scheme)
	// fmt.Println("host: ", parsedUrl.Host)
	// fmt.Println("path: ", parsedUrl.Path)
	// fmt.Println("port: ", parsedUrl.Port())
	// fmt.Println("raw query: ", parsedUrl.RawQuery)

	qparams := parsedUrl.Query()
	fmt.Println("query params: ", qparams["coursename"])

	for _, value := range qparams {
		fmt.Println("Param is: ", value)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcass",
		RawPath: "user=Minar",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("Another URL: ", anotherUrl)

}
