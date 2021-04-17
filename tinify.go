package tinify

import (
	"errors"
	"io/ioutil"
	"net/http"
)

type Source struct {
	url      string
	commands map[string]interface{}
}

func FromFile(path string) (s *Source, err error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	return FromBuffer(buf)
}

func FromBuffer(buf []byte) (s *Source, err error) {
	response, err := GetClient().Request(http.MethodPost, "/shrink", buf)
	if err != nil {
		return
	}

	s, err = getSourceFromResponse(response)
	return
}

func getSourceFromResponse(response *http.Response) (s *Source, err error) {
	location := response.Header["Location"]
	url := ""
	if len(location) > 0 {
		url = location[0]
	}

	s = newSource(url, nil)
	return
}

func newSource(url string, commands map[string]interface{}) *Source {
	s := new(Source)
	s.url = url
	if commands != nil {
		s.commands = commands
	} else {
		s.commands = make(map[string]interface{})
	}

	return s
}

func (s *Source) ToFile(path string) error {
	result, err := s.toResult()
	if err != nil {
		return err
	}

	return result.ToFile(path)
}

func (s *Source) toResult() (r *Result, err error) {
	if len(s.url) == 0 {
		err = errors.New("url is empty")
		return
	}

	//body := make([]byte, 0)
	//if len(s.commands) > 0 {
	//	body, err = json.Marshal(s.commands)
	//	if err != nil {
	//		return
	//	}
	//}
	response, err := GetClient().Request(http.MethodGet, s.url, s.commands)
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	r = NewResult(response.Header, data)
	return
}
