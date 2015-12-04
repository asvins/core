package main

import (
	"io"
	"net/url"

	"github.com/asvins/warehouse/decoder"
)

type QueryMapping func(url.Values) url.Values

func BuildStructFromQueryString(dst interface{}, queryString url.Values, mappingfs ...QueryMapping) error {
	v := queryString
	for _, f := range mappingfs {
		v = f(v)
	}
	decoder := decoder.NewDecoder()
	return decoder.DecodeURLValues(dst, v)
}

func BuildStructFromReqBody(dst interface{}, body io.ReadCloser) error {
	decoder := decoder.NewDecoder()
	return decoder.DecodeReqBody(dst, body)
}
