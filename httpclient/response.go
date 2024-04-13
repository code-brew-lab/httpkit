package httpclient

import (
	"encoding/json"
	"errors"
	"io"
)

func UnmarshalBody[TData interface{}](body io.ReadCloser) (*TData, error) {
	var (
		bytes   = make([]byte, 0)
		bodyObj = new(TData)
		err     error
	)

	if body == nil {
		return nil, errors.New("body is missing")
	}

	defer body.Close()

	bytes, err = io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, bodyObj)
	if err != nil {
		return nil, err
	}

	return bodyObj, nil
}
