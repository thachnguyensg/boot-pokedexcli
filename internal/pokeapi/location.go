package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationEndpoint struct {
	Previous *string
	Next     *string
	Url      string
}

func NewLocationEndpoint(url string) *LocationEndpoint {
	return &LocationEndpoint{
		Url: url,
	}
}

func (l *LocationEndpoint) GetNext() (*NamedResp, error) {
	var url string
	if l.Next != nil {
		url = *l.Next
	} else {
		url = l.Url
	}

	namedResp, err := request(url)
	if err != nil {
		return nil, err
	}

	l.Previous = namedResp.Previous
	l.Next = namedResp.Next

	return namedResp, nil
}

func (l *LocationEndpoint) GetPrev() (*NamedResp, error) {
	if l.Previous == nil {
		return nil, fmt.Errorf("No previous page")
	}

	namedResp, err := request(*l.Previous)
	if err != nil {
		return nil, fmt.Errorf("%v, url: %v", err, *l.Previous)
	}

	l.Previous = namedResp.Previous
	l.Next = namedResp.Next

	return namedResp, nil
}

func request(url string) (*NamedResp, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}
	namedResp := NamedResp{}
	err = json.Unmarshal(body, &namedResp)
	if err != nil {
		return nil, err
	}
	return &namedResp, nil
}
