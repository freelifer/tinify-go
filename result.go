package tinify

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type Result struct {
	data []byte
	*ResultMeta
}

type ResultMeta struct {
	meta http.Header
}

func NewResultMeta(meta http.Header) *ResultMeta {
	r := new(ResultMeta)
	r.meta = meta
	return r
}

func NewResult(meta http.Header, data []byte) *Result {
	r := new(Result)
	r.ResultMeta = NewResultMeta(meta)
	r.data = data
	return r
}

func (r *Result) ToFile(path string) error {
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, r.data, os.ModePerm)
	return err
}
