package request

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"

	"github.com/andybalholm/brotli"
)

type APIRequest struct {
	url      string
	params   url.Values
	headers  http.Header
	data     map[string]io.Reader
	response *http.Response
}

func NewAPIRequest(url string, params url.Values, headers http.Header, data map[string]io.Reader) *APIRequest {
	return &APIRequest{
		url:     url,
		params:  params,
		headers: headers,
		data:    data,
	}
}

func (r *APIRequest) SendRequest() (*APIRequest, error) {
	var req *http.Request
	var err error

	if r.data == nil {
		req, err = http.NewRequest("GET", r.url, nil)
		if err != nil {
			return r, err
		}

		if r.params != nil {
			req.URL.RawQuery = r.params.Encode()
		}
		req.Header = r.headers
	} else {
		b, contentType, err := prepareForm(r.data)
		if err != nil {
			return r, err
		}
		req, err = http.NewRequest("POST", r.url, b)
		if err != nil {
			return r, err
		}
		req.Header = r.headers
		req.Header.Set("Content-Type", contentType)
	}

	client := &http.Client{}
	r.response, err = client.Do(req)
	if err != nil {
		return r, err
	}

	return r, err
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

func (r *APIRequest) GetCookie(cookie string) (*http.Cookie, error) {
	cookies := r.response.Cookies()

	// Find the specific cookie by name
	for _, c := range cookies {
		if c.Name == cookie {
			return c, nil
		}
	}

	// Return an error if the cookie is not found
	return nil, fmt.Errorf("cookie not found: %s", cookie)
}

func prepareForm(values map[string]io.Reader) (*bytes.Buffer, string, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		var err error
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		if x, ok := r.(*os.File); ok {
			// Add file
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return nil, "", err
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return nil, "", err
			}
		}
		if _, err := io.Copy(fw, r); err != nil {
			return nil, "", err
		}

	}
	defer w.Close()

	return &b, w.FormDataContentType(), nil
}

func (r *APIRequest) getResponseContent() ([]byte, error) {
	contentEncoding := r.response.Header.Get("Content-Encoding")
	content, err := io.ReadAll(r.response.Body)
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

func parseJSONContent(content []byte) (map[string]interface{}, error) {
	var parsedContent map[string]interface{}
	err := json.Unmarshal(content, &parsedContent)
	if err != nil {
		return nil, err
	}

	return parsedContent, nil
}

func brotliDecode(content []byte) ([]byte, error) {
	r := brotli.NewReader(bytes.NewReader(content))
	return io.ReadAll(r)
}

func gzipDecode(content []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(content))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return io.ReadAll(r)
}
