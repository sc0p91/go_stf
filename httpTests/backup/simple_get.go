package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Simply defines the URL
	url := "https://jsonplaceholder.typicode.com/users"

	// Creates the Request
	req, _ := http.NewRequest("GET", url, nil)

	// Adding optional Header to Request
	req.Header.Add("cache-control", "no-cache")

	// Actually does the Request
	res, _ := http.DefaultClient.Do(req)

	// Closes the Request after finishing
	defer res.Body.Close()

	// Collects all data from response
	body, _ := ioutil.ReadAll(res.Body)

	// Prints the response data
	fmt.Println(string(body))
}
