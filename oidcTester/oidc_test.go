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

func TestBasicAuth(t *testing.T) {

	searchPattern := "Your last login attempt failed. Please check your input."
	formData := url.Values{
		"isiwebuserid": {"testuser"},
		"isiwebpasswd": {"testpass"},
	}

	res, err := http.PostForm("https://idpe.post.ch/auth/saml", formData)
	if err != nil {
		t.Error("Failed to connect")
	}

	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	defer res.Body.Close()

	// Check if the Response Status is OK
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Error(err)
		}

		t.Log("Answer from ", res.TLS.ServerName)
		t.Log("Status: ", res.Status)

		// Check if the Login failed
		bodyString := string(bodyBytes)
		if strings.ContainsAny(searchPattern, bodyString) {
			t.Error("Failed to Authenticate")
		}
	}
}
