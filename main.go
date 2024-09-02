package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const serverPort = 8500

// eofReaderFunc is an io.Reader that runs itself, and then returns io.EOF.
type eofReaderFunc func()

func (f eofReaderFunc) Read(p []byte) (n int, err error) {
	f()
	return 0, io.EOF
}

func main() {
	requestURL := fmt.Sprintf("http://localhost:%d/post", serverPort)
	var req *http.Request
	body, _ := os.Open("data/body.dat")
	req, err := http.NewRequest(http.MethodPost, requestURL, io.MultiReader(
		body,
		eofReaderFunc(func() {
			req.Trailer["Client-Trailer-A"] = []string{"valuea"}
		}),
		eofReaderFunc(func() {
			req.Trailer["Client-Trailer-B"] = []string{"valueb"}
		}),
	))
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}
	req.Trailer = http.Header{
		"Client-Trailer-A": nil, //  to be set later
		"Client-Trailer-B": nil, //  to be set later
	}
	req.ContentLength = -1

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", resBody)
}
