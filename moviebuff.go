// moviebuff-sdk-go is Go SDK for Moviebuff.
//
// The Moviebuff is a service that offers information about movies, people and entities.
//
// All resources are identified by a UUID, which Moviebuff.com uniquely and randomly generates.
// Since it may be difficult to get the Moviebuff UUID of any resource without prior knowledge,
// the API also allows substitution of the UUID with the URL identifier of the resource.
// For instance, if when attempting to get information for 12 Years a Slave (moviebuff.com/12-years-a-slave),
// either its UUID or the identifier 12-years-a-slave may be used.
package moviebuff

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

type Logger interface {
	Println(v ...interface{})
}

type Moviebuff interface {
	GetMovie(id string) (*Movie, error)
	GetPerson(id string) (*Person, error)
	GetEntity(id string) (*Entity, error)
	GetResources(resourceType ResourceType, limit, page int) (*Resources, error)
}

type Config struct {
	HostURL     string
	StaticToken string
	Logger      Logger
}

// Before accessing any API it need to be initialized.
// The Moviebuff is a service that offers information about movies, people, entities.
type moviebuff struct {
	Config
}

//New returns a QubeAccount interface with a nil logger.
func New(config Config) Moviebuff {
	mb := &moviebuff{Config: config}
	mb.init(config.Logger, config.HostURL, config.StaticToken)
	return mb
}

// Init initialize the Moviebuff.
// token is mandatory. logger is optional however.
// Return the same Moviebuff object back
func (m *moviebuff) init(l Logger, hostURL, staticToken string) {
	if l == nil {
		l = log.New(new(devNull), "", 0)
	}

	m.Logger = l
	m.HostURL = hostURL
	m.StaticToken = staticToken
}

// GetMovie fetch a movie and its basic details for given resource uuid.

// Instead of the UUID, this can also be the URL of the movie as seen on moviebuff.com, like 12-years-a-slave
// Details include release dates, certifications, cast, crew, trailers, posters, purchase links etc.
// Here movies may include feature films, documentaries, short films etc.
func (m *moviebuff) GetMovie(id string) (*Movie, error) {
	r, err := prepareRequest(m.HostURL, m.StaticToken, "/resources/movies/"+id)
	if err != nil {
		m.Logger.Println("Unable to create Request:", err)
		return nil, err
	}

	res, err := new(http.Client).Do(r)
	if err != nil {
		m.Logger.Println("Unable to make Request:", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Println("res.StatusCode...", res.StatusCode)
		if res.StatusCode == http.StatusForbidden {
			return nil, ErrInvalidToken
		}
		if res.StatusCode == http.StatusNotFound {
			return nil, ErrResourceDoesNotExist
		}
		m.Logger.Println("Got invalid res code: ", r, res.StatusCode)
		return nil, ErrResponseNotReceived
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		m.Logger.Println("Unable to read res body: ", err)
		return nil, err
	}

	movie := new(Movie)
	err = json.Unmarshal(content, movie)
	if err != nil {
		m.Logger.Println("Unable to unmarshal res body: ", err)
		return nil, err
	}

	if movie.Type != "movie" {
		m.Logger.Println("Resource is of type", movie.Type)
		return nil, ErrResourceDoesNotExist
	}

	return movie, nil
}

// GetPerson fetch a person and his/her basic details with UUID.

// Instead of the UUID, this can also be the URL of the person as seen on moviebuff.com, like amitabh-bachchan
// The people in the database include actors, directors, support personnel, etc.
// Moviebuff aims to document most, if not all, of the individuals involved in a film.
func (m *moviebuff) GetPerson(id string) (*Person, error) {
	r, err := prepareRequest(m.HostURL, m.StaticToken, "/resources/people/"+id)
	if err != nil {
		m.Logger.Println("Unable to create Request:", err)
		return nil, err
	}

	res, err := new(http.Client).Do(r)
	if err != nil {
		m.Logger.Println("Unable to make Request:", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusForbidden {
			return nil, ErrInvalidToken
		}
		if res.StatusCode == http.StatusNotFound {
			return nil, ErrResourceDoesNotExist
		}
		m.Logger.Println("Got invalid res code: ", res.StatusCode)
		return nil, ErrResponseNotReceived
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		m.Logger.Println("Unable to read res body: ", err)
		return nil, err
	}

	person := new(Person)
	err = json.Unmarshal(content, person)
	if err != nil {
		m.Logger.Println("Unable to unmarshal res body: ", err)
		return nil, err
	}

	if person.Type != "person" {
		m.Logger.Println("Resource is of type", person.Type)
		return nil, ErrResourceDoesNotExist
	}

	return person, nil
}

// GetEntity fetch an entity and its basic details for UUID.

// Instead of the UUID, this can also be the URL of the company as seen on moviebuff.com: yash-raj-films .
// Entities are usually organizations like production companies, service providers, etc.
func (m *moviebuff) GetEntity(id string) (*Entity, error) {
	r, err := prepareRequest(m.HostURL, m.StaticToken, "/resources/entities/"+id)
	if err != nil {
		m.Logger.Println("Unable to create Request:", err)
		return nil, err
	}

	res, err := new(http.Client).Do(r)
	if err != nil {
		m.Logger.Println("Unable to make Request:", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusForbidden {
			return nil, ErrInvalidToken
		}
		if res.StatusCode == http.StatusNotFound {
			return nil, ErrResourceDoesNotExist
		}
		m.Logger.Println("Got invalid res code: ", res.StatusCode)
		return nil, ErrResponseNotReceived
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		m.Logger.Println("Unable to read res body: ", err)
		return nil, err
	}

	entity := new(Entity)
	err = json.Unmarshal(content, entity)
	if err != nil {
		m.Logger.Println("Unable to unmarshal res body: ", err)
		return nil, err
	}

	if entity.Type != "entity" {
		m.Logger.Println("Resource is of type", entity.Type)
		return nil, ErrResourceDoesNotExist
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
func (m *moviebuff) GetResources(resourceType ResourceType, limit, page int) (*Resources, error) {
	u := "/resources/" + string(resourceType) + "?"
	if limit != 0 {
		u += "limit=" + strconv.Itoa(limit)
	}

	if page != 0 {
		u += "page=" + strconv.Itoa(page)
	}

	r, err := prepareRequest(m.HostURL, m.StaticToken, u)
	if err != nil {
		m.Logger.Println("Unable to create Request:", err)
		return nil, err
	}

	res, err := new(http.Client).Do(r)
	if err != nil {
		m.Logger.Println("Unable to make Request:", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusForbidden {
			return nil, ErrInvalidToken
		}
		if res.StatusCode == http.StatusNotFound {
			return nil, ErrResourceDoesNotExist
		}
		m.Logger.Println("Got invalid res code: ", res.StatusCode)
		return nil, ErrResponseNotReceived
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		m.Logger.Println("Unable to read res body: ", err)
		return nil, err
	}

	resources := new(Resources)
	err = json.Unmarshal(content, resources)
	if err != nil {
		m.Logger.Println("Unable to unmarshal res body: ", err)
		return nil, err
	}

	return resources, nil
}
