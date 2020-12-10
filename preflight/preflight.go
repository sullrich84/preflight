package preflight

import (
	"crypto/tls"
	"github.com/sullrich84/preflight/util/array"
	"github.com/sullrich84/preflight/util/terminal"
	"net/http"
	"strings"
)

const (
	AccessControlAllowOrigin      = "Access-Control-Allow-Origin"
	AccessControlAllowMethods     = "Access-Control-Allow-Methods"
	AccessControlAllowHeaders     = "Access-Control-Allow-Headers"
	AccessControlAllowCredentials = "Access-Control-Allow-Credentials"
)

var client = http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{},
	},
}

func Preflight(target string, origins []string, methods []string, headers []string) {
	terminal.PrettyPrintResultsHeader(origins)
	for _, method := range methods {
		var results []bool
		for _, origin := range origins {
			request := buildRequest(target, origin, method, headers)
			response := doRequest(request)

			originAllowed := originAllowed(response, origin)
			methodAllowed := methodAllowed(response, method)
			headersAllowed := headersAllowed(response, headers)
			//credentialsAllowed := credentialsAllowed(response)

			allowed := originAllowed && methodAllowed && headersAllowed
			results = append(results, allowed)
		}
		terminal.PrettyPrintResults(method, results)
	}
	terminal.PrettyPrintResultsFooter(origins)
}

func originAllowed(response *http.Response, origin string) bool {
	allowedOrigins := response.Header.Get(AccessControlAllowOrigin)
	return allowedOrigins == "*" || strings.Contains(allowedOrigins, origin)
}

func methodAllowed(response *http.Response, origin string) bool {
	allowedMethods := response.Header.Get(AccessControlAllowMethods)
	return allowedMethods == "*" || strings.Contains(allowedMethods, origin)
}

func headersAllowed(response *http.Response, headers []string) bool {
	allowedHeaders := response.Header.Get(AccessControlAllowHeaders)
	if allowedHeaders == "*" {
		return true
	}

	allowedHeadersArray := strings.Split(allowedHeaders, ",")
	return array.ContainsAll(allowedHeadersArray, headers)
}

func credentialsAllowed(response *http.Response) bool {
	return response.Header.Get(AccessControlAllowCredentials) == "true"
}
