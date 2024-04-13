package httpclient

import (
	"bytes"
)

type (
	query struct {
		params map[string]string
	}

	QueryBuilder interface {
		Add(key string, value string)
		Remove(key string)
		Build() string
	}
)

func NewQuery() QueryBuilder {
	return &query{params: make(map[string]string)}
}

func (q *query) Add(key string, value string) {
	if len(key) == 0 || len(value) == 0 {
		return
	}
	q.params[key] = value
}

func (q *query) Remove(key string) {
	if len(key) == 0 {
		return
	}
	delete(q.params, key)
}

func (q *query) Build() string {
	if len(q.params) == 0 {
		return ""
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("?")

	index := 0
	for key, value := range q.params {
		index++
		var buffer bytes.Buffer

		buffer.WriteString(key)
		buffer.WriteString("=")
		buffer.WriteString(value)
		queryBuffer.WriteString(buffer.String())

		if index == len(q.params) {
			break
		}
		queryBuffer.WriteString("&")
	}

	return queryBuffer.String()
}
