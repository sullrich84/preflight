package preflight

import (
	"log"
	"net/http"
	"strings"
)

const (
	Origin                      = "Origin"
	AccessControlRequestMethod  = "Access-Control-Request-Method"
	AccessControlRequestHeaders = "Access-Control-Request-Headers"
)

func newRequest(target string) *http.Request {
	request, err := http.NewRequest(http.MethodOptions, target, nil)
	if err != nil {
		log.Fatal(err)
	}

	return request
}

func buildRequest(target string, origin string, method string, header []string) *http.Request {
	request := newRequest(target)

	request.Header.Add(Origin, origin)
	request.Header.Add(AccessControlRequestMethod, method)
	request.Header.Add(AccessControlRequestHeaders, strings.Join(header, ","))

	return request
}

func doRequest(request *http.Request) *http.Response {
	response, requestErr := client.Do(request)
	if requestErr != nil {
		log.Fatalf("Could not request %s", request.Host)
	}

	return response
}
