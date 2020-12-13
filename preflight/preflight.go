package preflight

import (
	"crypto/tls"
	"github.com/thoas/go-funk"
	"log"
	"net/http"
	"strings"
)

type PreFlight struct {
	target string
	origin string
	method string
	header []string
}

// client is the default http client with https support,
// preflight will use to make it's pre-flights.
var client = http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{},
	},
}

// NewPreFlight initializes a new CORS pre-flight scenario.
func NewPreFlight(target string, origin string, method string, header []string) (*PreFlight, error) {
	return &PreFlight{target, origin, method, header}, nil
}

// PreFly will do the actual CORS pre-flight.
//
// An options-request will be send to the target having origin
// as requester. The resolving response will be evaluated an
// results in the returned success  state which will either be
// true (stating a successful pre-flight) or false stating an
// unsuccessful pre-flight.
func (preFlight *PreFlight) PreFly() bool {
	request, reqErr := preFlight.newRequest()
	if reqErr != nil {
		log.Println(reqErr)
		return false
	}

	response, resErr := client.Do(request)
	if resErr != nil {
		log.Println(resErr)
		return false
	}

	// Assume failed failed when status is not OK
	if response.StatusCode != 200 {
		return false
	}

	originAllowed := preFlight.originAllowed(response)
	methodAllowed := preFlight.methodAllowed(response)
	headersAllowed := preFlight.headersAllowed(response)

	return originAllowed && methodAllowed && headersAllowed
}

// newRequest will initialize a new request containing all relevant CORS headers.
func (preFlight *PreFlight) newRequest() (*http.Request, error) {
	req, err := http.NewRequest(http.MethodOptions, preFlight.target, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Origin", preFlight.origin)
	req.Header.Add("Access-Control-Request-Method", preFlight.method)

	headers := strings.Join(preFlight.header, ",")
	req.Header.Add("Access-Control-Request-Headers", headers)

	return req, err
}

// originAllowed verifies the origin is listed in the cors response header.
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin
func (preFlight *PreFlight) originAllowed(res *http.Response) bool {
	origins := res.Header.Get("Access-Control-Allow-Origin")
	return origins == "*" || strings.Contains(origins, preFlight.origin)
}

// originAllowed verifies the method is listed in the cors response header.
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Methods
func (preFlight *PreFlight) methodAllowed(res *http.Response) bool {
	methods := res.Header.Get("Access-Control-Allow-Methods")
	return methods == "*" || strings.Contains(methods, preFlight.method)
}

// headersAllowed verifies the headers are listed in the cors response header.
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Headers
func (preFlight *PreFlight) headersAllowed(res *http.Response) bool {
	headers := res.Header.Get("Access-Control-Allow-Headers")
	if headers == "*" {
		return true
	}

	allowedHeadersArray := strings.Split(headers, ",")
	intersect := funk.IntersectString(preFlight.header, allowedHeadersArray)
	return len(intersect) > 0
}
