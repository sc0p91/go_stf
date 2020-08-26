package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestConnection(t *testing.T) {
	// Create the Request (URL, Method and Headers)
	url := "https://idpe.post.ch/auth/"
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

// -> This would be the test for a successful login...
// func TestBasicAuth(t *testing.T) {

// 	searchPattern := "failed"
// 	formData := url.Values{
// 		"isiwebuserid": {"blibli"},
// 		"isiwebpasswd": {"blublu"},
// 		"transferID":   {"ac11911a-1c864-12ac3671-17401c80c3c-001a8164"},
// 		"application":  {"6b1b55c69af175703f68ab6727514474"},
// 	}

// 	res, err := http.PostForm("https://idpe.post.ch/auth/saml", formData)
// 	if err != nil {
// 		t.Error("Failed to connect")
// 	}

// 	var result map[string]interface{}
// 	json.NewDecoder(res.Body).Decode(&result)

// 	defer res.Body.Close()

// 	// Check if the Response Status is OK
// 	if res.StatusCode == http.StatusOK {
// 		bodyBytes, err := ioutil.ReadAll(res.Body)
// 		if err != nil {
// 			t.Error(err)
// 		}

// 		t.Log("Answer from ", res.TLS.ServerName)
// 		t.Log("Status: ", res.Status)

// 		// Check if the Login failed
// 		bodyString := string(bodyBytes)
// 		t.Log(bodyString)
// 		result1 := strings.Contains(bodyString, searchPattern)
// 		result2 := strings.Index(bodyString, searchPattern)
// 		t.Log(result1, result2)
// 		if strings.Contains(bodyString, searchPattern) {
// 			t.Error("Failed to Authenticate")
// 		}
// 	}
// }
