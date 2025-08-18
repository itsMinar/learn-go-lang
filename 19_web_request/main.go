package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/posts"

func main() {
	fmt.Println("Web Request in Go")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Println("response: ", response)

	defer response.Body.Close()

	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	dataString := string(dataBytes)
	fmt.Println("dataString: ", dataString)
}
