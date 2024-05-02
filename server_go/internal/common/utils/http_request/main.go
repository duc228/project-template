package http_request

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Header struct {
	Key   string
	Value string
}

type Param struct {
	Name  string
	Value interface{}
}

func Send(method string, link string, headers map[string][]string, paramMap map[string][]Param) (*http.Response, error) {
	client := &http.Client{
		// Timeout: time.Second * 10,
	}

	var body io.Reader = nil
	if paramMap != nil {

		for key, params := range paramMap {
			data := url.Values{}
			for _, param := range params {
				data.Add(param.Name, fmt.Sprint(param.Value))
			}

			switch key {
			case "query":
				link = fmt.Sprintf("%s?%s", link, data.Encode())

			case "body":
				body = strings.NewReader(data.Encode())

			default:
				fmt.Println("Nothing To DO")
			}
		}
	}

	req, reqErr := http.NewRequest(method, link, body)
	req.Header = headers

	if reqErr != nil {
		return nil, fmt.Errorf("Got error %s", reqErr.Error())
	}
	res, resErr := client.Do(req)
	if resErr != nil {
		return nil, fmt.Errorf("Got error %s", resErr.Error())
	}

	// defer res.Body.Close()

	return res, resErr
}

func Get(url string, headers map[string][]string, paramMap map[string][]Param) (*http.Response, error) {
	return Send("GET", url, headers, paramMap)
}

func Post(url string, headers map[string][]string, paramMap map[string][]Param) (*http.Response, error) {
	return Send("POST", url, headers, paramMap)
}
