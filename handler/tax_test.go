package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaxCalculate(t *testing.T) {
	// Setup
	var tr TaxRequest
	body := bytes.NewBufferString(`{
		"totalIncome": 500000.0,
		"wht": 0.0
	  }`)

	res := request(http.MethodPost, uri("api/tax/calculations"), body)
	err := res.Decode(&tr)
	if err != nil {
		t.Fatal("cannot calulate tax", err)
	}

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, res.StatusCode)
}

func uri(paths ...string) string {
	host := "http://localhost:8080"
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return &Response{res, err}
}

type Response struct {
	*http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	return json.NewDecoder(r.Body).Decode(v)
}
