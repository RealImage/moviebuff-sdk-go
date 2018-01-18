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
func (m *Moviebuff) GetMovie(id string) (*Movie, error) {
	r, err := prepareRequest(m.token, "/resources/movies/"+id)
	if err != nil {
		m.l.Println("Unable to create Request:", err)
		return nil, err
	}

	res, err := new(http.Client).Do(r)
	if err != nil {
		m.l.Println("Unable to make Request:", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		m.l.Println("Got invalid res code. Status: ", res.StatusCode)
		return nil, ErrResponseNotReceived
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		m.l.Println("Unable to read res body: ", err)
		return nil, err
	}

	movie := new(Movie)
	err = json.Unmarshal(content, movie)
	if err != nil {
		m.l.Println("Unable to unmarshal res body: ", err)
		return nil, err
	}

	return movie, nil
}

// GetPerson fetch a person and his/her basic details.
// The people in the database include actors, directors, support personnel, etc.
// Moviebuff aims to document most, if not all, of the individuals involved in a film.
func (m *Moviebuff) GetPerson(id string) (*Person, error) {
	r, err := prepareRequest(m.token, "/resources/people/"+id)
	if err != nil {
		m.l.Println("Unable to create Request:", err)
		return nil, err
	}

	res, err := new(http.Client).Do(r)
	if err != nil {
		m.l.Println("Unable to make Request:", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		m.l.Println("Got invalid res code. Status: ", res.StatusCode)
		return nil, ErrResponseNotReceived
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		m.l.Println("Unable to read res body: ", err)
		return nil, err
	}

	person := new(Person)
	err = json.Unmarshal(content, person)
	if err != nil {
		m.l.Println("Unable to unmarshal res body: ", err)
		return nil, err
	}

	return person, nil
}

// GetEntity fetch an entity and its basic details.
// Entities are usually organizations like production companies, service providers, etc.
func (m *Moviebuff) GetEntity(id string) (*Entity, error) {
	r, err := prepareRequest(m.token, "/resources/entities/"+id)
	if err != nil {
		m.l.Println("Unable to create Request:", err)
		return nil, err
	}

	res, err := new(http.Client).Do(r)
	if err != nil {
		m.l.Println("Unable to make Request:", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		m.l.Println("Got invalid res code. Status: ", res.StatusCode)
		return nil, ErrResponseNotReceived
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		m.l.Println("Unable to read res body: ", err)
		return nil, err
	}

	entity := new(Entity)
	err = json.Unmarshal(content, entity)
	if err != nil {
		m.l.Println("Unable to unmarshal res body: ", err)
		return nil, err
	}

	return entity, nil
}
