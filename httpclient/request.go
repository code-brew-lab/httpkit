package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
	defaultTimeout = 2 * time.Second
)

func MakeRequest(ctx context.Context, req *http.Request, timeout ...time.Duration) (*http.Response, error) {
	var (
		t          time.Duration
		respStream = make(chan *http.Response)
		errStream  = make(chan error)
	)

	if len(timeout) == 0 {
		t = defaultTimeout
	} else {
		t = timeout[0]
	}

	ctx, cancel := context.WithTimeout(ctx, t)
	defer cancel()
	defer close(respStream)
	defer close(errStream)

	go func(respStream chan<- *http.Response, errStream chan<- error) {
		client := &http.Client{}
		defer client.CloseIdleConnections()

		resp, err := client.Do(req)
		if err != nil {
			errStream <- err
			return
		}
		respStream <- resp
	}(respStream, errStream)

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case resp := <-respStream:
			return resp, nil
		case err := <-errStream:
			return nil, err
		}
	}

}

func MarshallBody(body interface{}) (io.Reader, error) {
	var (
		data []byte
		err  error
	)

	data, err = json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(data), nil
}
