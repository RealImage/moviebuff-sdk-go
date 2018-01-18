package moviego

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	ErrInvalidToken         = errors.New("access denied")
	ErrResponseNotReceived  = errors.New("could not receive valid response")
	ErrResourceDoesNotExist = errors.New("resource does not exist")
)

type Moviebuff struct {
	token string
	l     logger
}

func (m *Moviebuff) Init(token string, l logger) {
	if l == nil {
		l = log.New(new(devNull), "", 0)
	}

	m.l = l
	m.token = token
}

// GetMovie fetch a movie and its basic details
// like release dates, certifications, cast, crew, trailers, posters, purchase links etc.
// Here movies may include feature films, documentaries, short films etc.
func (m Moviebuff) GetMovie(id string) (*Movie, error) {
	path := "/resources/movies/" + id
	r, err := prepareRequest(m.token, path)
	if err != nil {
		m.l.Println("Unable to create Request:", hostUrl+path, err)
		return nil, err
	}

	c := new(http.Client)
	res, err := c.Do(r)
	if err != nil {
		m.l.Println("Unable to make Request:", hostUrl+path, err)
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		m.l.Println("Got invalid res code. Status: ", res.StatusCode)
		return nil, ErrResponseNotReceived
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		m.l.Println("Unable to read res body, Err:", err)
		return nil, err
	}

	movie := new(Movie)
	err = json.Unmarshal(content, movie)
	if err != nil {
		m.l.Println("Unable to unmarshal res body, Err:", err)
		return nil, err
	}

	return movie, nil
}
