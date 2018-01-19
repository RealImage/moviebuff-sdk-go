// moviego is Go SDK for Moviebuff.
//
// The Moviebuff is a service that offers information about movies, people and entities.
//
// All resources are identified by a UUID, which Moviebuff.com uniquely and randomly generates.
// Since it may be difficult to get the Moviebuff UUID of any resource without prior knowledge,
// the API also allows substitution of the UUID with the URL identifier of the resource.
// For instance, if when attempting to get information for 12 Years a Slave (moviebuff.com/12-years-a-slave),
// either its UUID or the identifier 12-years-a-slave may be used.
package moviego

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Type of resources supported by moviebuff
type ResourceType string

const (
	RESOURCE_TYPE_PEOPLE   ResourceType = "people"
	RESOURCE_TYPE_MOVIES   ResourceType = "movies"
	RESOURCE_TYPE_ENTITIES ResourceType = "entities"
)

var (
	ErrInvalidToken         = errors.New("access denied")
	ErrResponseNotReceived  = errors.New("could not receive valid response")
	ErrResourceDoesNotExist = errors.New("resource does not exist")
)

// Moviebuff allows to access to information in moviebuff using resource ids.

// Before accessing any API it need to be initialized.
// The Moviebuff is a service that offers information about movies, people, entities.
type Moviebuff struct {
	token string
	l     logger
}

// Init initialize the Moviebuff.
// token is mandatory. logger is optional however.
// Return the same Moviebuff object back
func (m *Moviebuff) Init(token string, l logger) *Moviebuff {
	if l == nil {
		l = log.New(new(devNull), "", 0)
	}

	m.l = l
	m.token = token
	return m
}

// GetMovie fetch a movie and its basic details for given resource uuid.

// Instead of the UUID, this can also be the URL of the movie as seen on moviebuff.com, like 12-years-a-slave
// Details include release dates, certifications, cast, crew, trailers, posters, purchase links etc.
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
		if res.StatusCode == http.StatusForbidden {
			return nil, ErrInvalidToken
		}
		if res.StatusCode == http.StatusForbidden {
			return nil, ErrInvalidToken
		}
		m.l.Println("Got invalid res code: ", r, res.StatusCode)
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

// GetPerson fetch a person and his/her basic details with UUID.

// Instead of the UUID, this can also be the URL of the person as seen on moviebuff.com, like amitabh-bachchan
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
		if res.StatusCode == http.StatusForbidden {
			return nil, ErrInvalidToken
		}
		m.l.Println("Got invalid res code: ", res.StatusCode)
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

// GetEntity fetch an entity and its basic details for UUID.

// Instead of the UUID, this can also be the URL of the company as seen on moviebuff.com: yash-raj-films .
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
		if res.StatusCode == http.StatusForbidden {
			return nil, ErrInvalidToken
		}
		m.l.Println("Got invalid res code: ", res.StatusCode)
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

// GetResource fetch list of resource of type resourceType.
//
// resourceType may be one of the following value defined by ResourceType
// page and limit is optional. Provide zero values to ignore them.
//
// limit represents the number of records to fetch in a single request.
// The actual count can be lower than the provided limit. Max value is 50.
// page represents the page number in the pagination. It starts from 1.
func (m *Moviebuff) GetResources(resourceType ResourceType, limit, page int) (*Resources, error) {
	u := "/resources/" + string(resourceType) + "?"
	if limit != 0 {
		u += "limit=" + strconv.Itoa(limit)
	}

	if page != 0 {
		u += "page=" + strconv.Itoa(page)
	}

	r, err := prepareRequest(m.token, u)
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
		if res.StatusCode == http.StatusForbidden {
			return nil, ErrInvalidToken
		}
		m.l.Println("Got invalid res code: ", res.StatusCode)
		return nil, ErrResponseNotReceived
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		m.l.Println("Unable to read res body: ", err)
		return nil, err
	}

	resources := new(Resources)
	err = json.Unmarshal(content, resources)
	if err != nil {
		m.l.Println("Unable to unmarshal res body: ", err)
		return nil, err
	}

	return resources, nil
}
