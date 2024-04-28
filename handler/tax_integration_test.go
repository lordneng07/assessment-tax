//go:build integration
// +build integration

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

	t.Run("validate allowanceType is not correct", func(t *testing.T) {
		var verr Err

		reqBody := bytes.NewBufferString(`{
			"totalIncome": 500000.0,
			"wht": 0.0,
			"allowances": [
			  {
				"allowanceType": "a",
				"amount": 0.0
			  }
			]}`)

		res := request(http.MethodPost, uri("tax/calculations"), reqBody)

		err := res.Decode(&verr)
		if err != nil {
			t.Fatal("cannot parse tax request", err)
		}

		assert.Nil(t, err)
		assert.EqualValues(t, verr.Message, "allowanceType is not correct")
		assert.EqualValues(t, http.StatusBadRequest, res.StatusCode)

	})

	t.Run("total income 50000, who = 0, allowances = 0 tax should be 29000", func(t *testing.T) {
		var re TaxResponse
		body := bytes.NewBufferString(`{
		"totalIncome": 500000.0,
		"wht": 0.0,
		"allowances": [
		  {
			"allowanceType": "donation",
			"amount": 0.0
		  }
		]}`)

		res := request(http.MethodPost, uri("tax/calculations"), body)
		err := res.Decode(&re)
		if err != nil {
			t.Fatal("cannot calulate tax", err)
		}

		assert.Nil(t, err)
		assert.EqualValues(t, re.Tax, 29000)
		assert.EqualValues(t, http.StatusOK, res.StatusCode)
	})

	t.Run("total income 50000, who = 0, allowances = 0 tax should be 4000", func(t *testing.T) {
		var re TaxResponse
		body := bytes.NewBufferString(`{
			"totalIncome": 500000.0,
			"wht": 25000.0,
			"allowances": [
			  {
				"allowanceType": "donation",
				"amount": 0.0
			  }
			]
		  }`)

		res := request(http.MethodPost, uri("tax/calculations"), body)
		err := res.Decode(&re)
		if err != nil {
			t.Fatal("cannot calulate tax", err)
		}

		assert.Nil(t, err)
		assert.EqualValues(t, re.Tax, 4000)
		assert.EqualValues(t, http.StatusOK, res.StatusCode)
	})
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
