package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestConnection(t *testing.T) {
	// Create the Request (URL, Method and Headers)
	url := "https://idpe.post.ch/auth/sam"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("cache-control", "no-cache")

	// Get the login page and close the connection afterwards
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error("Failed to connect")
		os.Exit(1)
	}
	defer res.Body.Close()

	// Check if the Response Status is OK
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Error(err)
		}

		t.Log("Answer from ", res.TLS.ServerName)
		t.Log("Status: ", res.Status)

		_ = bodyBytes
		// --> Outputs the whole body
		//bodyString := string(bodyBytes)
		//t.Log(bodyString)
	}
}
