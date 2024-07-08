package core

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

var httpClient *http.Client

func init() {
	transport := &http.Transport{}
	transport.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
	httpClient = &http.Client{Transport: transport}
}

func getZip(uriAddress string) (zipFile io.ReaderAt, size int64, err error) {
	res, err := httpClient.Get(uriAddress)
	if err != nil {
		return
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		bodyStr, _ := io.ReadAll(res.Body)
		err = fmt.Errorf("http exception(%v): %v", res.StatusCode, string(bodyStr))
		return
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	zipFile = bytes.NewReader(data)
	size = int64(len(data))
	return
}
