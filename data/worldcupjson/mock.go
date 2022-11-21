package worldcupjson

import (
	"bytes"
	"errors"
	"io"
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
		b, err := os.ReadFile(path + "/data/worldcupjson/mock/matches.json")
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewReader(b)),
		}, nil
	}

	if req.URL.Path == "/teams" {
		b, err := os.ReadFile(path + "/data/worldcupjson/mock/teams.json")
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewReader(b)),
		}, nil
	}

	return nil, errors.New("unsupported mock request")
}
