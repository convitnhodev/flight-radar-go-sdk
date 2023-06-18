package request

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/andybalholm/brotli"
)

type APIRequest struct {
	url      string
	params   url.Values
	headers  http.Header
	data     url.Values
	response *http.Response
}

func (r *APIRequest) sendRequest() error {
	var req *http.Request
	var err error

	if r.data == nil {
		req, err = http.NewRequest("GET", r.url, nil)
		if err != nil {
			return err
		}

		if r.params != nil {
			req.URL.RawQuery = r.params.Encode()
		}
	} else {
		req, err = http.NewRequest("POST", r.url, strings.NewReader(r.data.Encode()))
		if err != nil {
			return err
		}
	}

	req.Header = r.headers

	client := &http.Client{}
	r.response, err = client.Do(req)
	return err
}

func NewAPIRequest(url string, params url.Values, headers http.Header, data url.Values) (*APIRequest, error) {
	request := &APIRequest{
		url:     url,
		params:  params,
		headers: headers,
		data:    data,
	}
	if err := request.sendRequest(); err != nil {
		return nil, err
	}
	return request, nil
}

func (r *APIRequest) GetContent() (interface{}, error) {
	content, err := r.getResponseContent()
	if err != nil {
		return nil, err
	}

	contentType := r.response.Header.Get("Content-Type")
	if contentType == "application/json" {
		return parseJSONContent(content)
	}

	return content, nil
}

func (r *APIRequest) getResponseContent() ([]byte, error) {
	contentEncoding := r.response.Header.Get("Content-Encoding")
	content, err := ioutil.ReadAll(r.response.Body)
	if err != nil {
		return nil, err
	}

	if contentEncoding == "br" {
		return brotliDecode(content)
	} else if contentEncoding == "gzip" {
		return gzipDecode(content)
	}

	return content, nil
}

func (api *APIRequest) GetCookie(cookie string) (*http.Cookie, error) {
	cookies := api.response.Cookies()

	// Find the specific cookie by name
	for _, c := range cookies {
		if c.Name == cookie {
			return c, nil
		}
	}

	// Return an error if the cookie is not found
	return nil, fmt.Errorf("Cookie not found: %s", cookie)
}

func parseJSONContent(content []byte) (interface{}, error) {
	var parsedContent interface{}
	err := json.Unmarshal(content, &parsedContent)
	if err != nil {
		return nil, err
	}

	return parsedContent, nil
}

func brotliDecode(content []byte) ([]byte, error) {
	r := brotli.NewReader(bytes.NewReader(content))
	return ioutil.ReadAll(r)
}

func gzipDecode(content []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(content))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return ioutil.ReadAll(r)
}

//func main() {
//	// Example usage
//	params := url.Values{}
//	params.Set("key1", "value1")
//	params.Set("key2", "value2")
//
//	headers := http.Header{}
//	headers.Set("Content-Type", "application/json")
//	headers.Set("Authorization", "Bearer token123")
//
//	request := NewAPIRequest("https://example.com/api", params, headers, nil)
//	err := request.sendRequest()
//	if err != nil {
//		panic(err)
//	}
//
//	content, err := request.GetContent()
//	if err != nil {
//		panic(err)
//	}
//
//	// Handle the content accordingly
//	switch content.(type) {
//	case []byte:
//		// Content is a byte array
//		// Handle as needed
//	case map[string]interface{}:
//		// Content is a JSON object
//		// Handle as needed
//	default:
//		// Content has an unknown type
//
//	}
//}
