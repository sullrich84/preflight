package preflight

import (
	"crypto/tls"
	"github.com/sullrich84/preflight/util/array"
	"net/http"
	"strings"
)

const (
	AccessControlAllowOrigin  = "Access-Control-Allow-Origin"
	AccessControlAllowMethods = "Access-Control-Allow-Methods"
	AccessControlAllowHeaders = "Access-Control-Allow-Headers"
)

var client = http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{},
	},
}

// Preflight will perform the CORS preflight the browser would usually do.
func Preflight(recipe Recipe, callback func(method string, results []bool)) {
	for _, method := range recipe.Methods {
		var results []bool
		for _, origin := range recipe.Origins {
			request := buildRequest(recipe.Target, origin, method, recipe.Headers)
			response := doRequest(request)

			originAllowed := originAllowed(response, origin)
			methodAllowed := methodAllowed(response, method)
			headersAllowed := headersAllowed(response, recipe.Headers)

			allowed := originAllowed && methodAllowed && headersAllowed
			results = append(results, allowed)
		}
		callback(method, results)
	}
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
