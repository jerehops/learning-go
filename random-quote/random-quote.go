package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
	"encoding/json"

)

type MyResponse []struct {
    Content string  `json:"content"`
}

// TODO: Use native http client. With timeouts for best practices.

func main() {
	response, err := http.Get("https://api.quotable.io/quotes/random")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := io.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
	var output MyResponse
	json.Unmarshal(responseData, &output)
	fmt.Println(output[0])
}