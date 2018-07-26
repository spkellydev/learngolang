package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spkellydev/learngolang/src/approutes"
	"github.com/spkellydev/learngolang/src/approutes/server"
)

func TestHandler(t *testing.T) {
	// Form a new HTTP request using the testing suite
	// args
	// 1 > method
	// 2 > route
	// 3 ? requestbody : nil
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Use Go's httptest library to create an http recorder
	// target of the http request (mini broswer)
	recorder := httptest.NewRecorder()

	// Create an HTTP handler from our handler function
	// defined in main.go
	hf := http.HandlerFunc(server.NonStaticResourceHandler)

	// Serve the HTTP request to our recorder
	// executes the handler that we want to test
	hf.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// TOOD -- render static content and check output in context

}

// func TestRouter(t *testing.T) {
// 	// call new router
// 	r := newRouter()
// 	mockServer := httptest.NewServer(r)

// 	resp, err := http.Get(mockServer.URL + "/")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		t.Errorf("Status should be ok, got %d", resp.StatusCode)
// 	}

// 	defer resp.Body.Close()
// 	b, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// convert the bytes to a string
// 	respString := string(b)

// 	expected := `<!DOCTYPE html>
// 		<html lang="en">

// 		<head>
// 		  <meta charset="UTF-8">
// 		  <meta name="viewport" content="width=device-width, initial-scale=1.0">
// 		  <meta http-equiv="X-UA-Compatible" content="ie=edge">
// 		  <link rel="stylesheet" href="./assets/css/style.css">
// 		  <title>Learn GoLang | Go Lang: The Language of the Future</title>
// 		</head>

// 		<body>
// 		  <div>
// 		    <h1>Learn Go Lang</h1>
// 		    <p style="text-align: center;">Coming soon...</p>
// 		  </div>
// 		</body>

// 		</html>`
// 	if respString != expected {
// 		t.Errorf("Response should be %s got %s", expected, respString)
// 	}

// }

func TestRouterForNonExistentRoute(t *testing.T) {
	r := approutes.NewRouter()
	mockServer := httptest.NewServer(r)
	// Most of the code is similar. The only difference is that now we make a
	//request to a route we know we didn't define, like the `POST /hello` route.
	resp, err := http.Post(mockServer.URL+"/hello", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 405 (method not allowed)
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405, got %d", resp.StatusCode)
	}

	// The code to test the body is also mostly the same, except this time, we
	// expect an empty body
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString := string(b)
	expected := ""

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}

}
