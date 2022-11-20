package footballdata

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

type mockHttpClient struct{}

func (m *mockHttpClient) Do(req *http.Request) (*http.Response, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	if req.URL.Path == "/matches" {
		b, err := ioutil.ReadFile(path + "/data/footballdata/mock/matches.json")
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(b)),
		}, nil
	}

	if req.URL.Path == "/teams" {
		b, err := ioutil.ReadFile(path + "/data/footballdata/mock/standings.json")
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(b)),
		}, nil
	}

	return nil, errors.New("unsupported mock request")
}
