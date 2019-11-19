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
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
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
type Moviebuff interface {
	GetMovie(ctx context.Context, id string) (*Movie, error)
	GetPerson(ctx context.Context, id string) (*Person, error)
	GetEntity(ctx context.Context, id string) (*Entity, error)
	GetResources(ctx context.Context, resourceType ResourceType, limit, page int) (*Resources, error)
	GetCertifications(ctx context.Context, country string) ([]Certification, error)
	GetHolidayCalendar(ctx context.Context, countryID string) (*Calendar, error)
}

type Config struct {
	HostURL     string
	StaticToken string
	Client      *http.Client
}

// Before accessing any API it need to be initialized.
// The Moviebuff is a service that offers information about movies, people, entities.
type moviebuff struct {
	Config
}

// New returns a Moviebuff interface.
func New(config Config) Moviebuff {
	if config.Client == nil {
		config.Client = http.DefaultClient
	}
	return &moviebuff{
		Config: config,
	}
}

// GetMovie fetch a movie and its basic details for given resource uuid.

// Instead of the UUID, this can also be the URL of the movie as seen on moviebuff.com, like 12-years-a-slave
// Details include release dates, certifications, cast, crew, trailers, posters, purchase links etc.
// Here movies may include feature films, documentaries, short films etc.
func (m *moviebuff) GetMovie(ctx context.Context, id string) (*Movie, error) {
	r, err := prepareRequest(ctx, m.HostURL, m.StaticToken, "/resources/movies/"+id)
	if err != nil {
		return nil, err
	}

	res, err := m.Client.Do(r)
	if err != nil {
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
		return nil, ErrResponseNotReceived
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	movie := new(Movie)
	err = json.Unmarshal(content, movie)
	if err != nil {
		return nil, err
	}

	if movie.Type != "movie" {
		return nil, ErrResourceDoesNotExist
	}

	return movie, nil
}

// GetPerson fetch a person and his/her basic details with UUID.

// Instead of the UUID, this can also be the URL of the person as seen on moviebuff.com, like amitabh-bachchan
// The people in the database include actors, directors, support personnel, etc.
// Moviebuff aims to document most, if not all, of the individuals involved in a film.
func (m *moviebuff) GetPerson(ctx context.Context, id string) (*Person, error) {
	r, err := prepareRequest(ctx, m.HostURL, m.StaticToken, "/resources/people/"+id)
	if err != nil {
		return nil, err
	}

	res, err := m.Client.Do(r)
	if err != nil {
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
		return nil, ErrResponseNotReceived
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	person := new(Person)
	err = json.Unmarshal(content, person)
	if err != nil {
		return nil, err
	}

	if person.Type != "person" {
		return nil, ErrResourceDoesNotExist
	}

	return person, nil
}

// GetEntity fetch an entity and its basic details for UUID.

// Instead of the UUID, this can also be the URL of the company as seen on moviebuff.com: yash-raj-films .
// Entities are usually organizations like production companies, service providers, etc.
func (m *moviebuff) GetEntity(ctx context.Context, id string) (*Entity, error) {
	r, err := prepareRequest(ctx, m.HostURL, m.StaticToken, "/resources/entities/"+id)
	if err != nil {
		return nil, err
	}

	res, err := m.Client.Do(r)
	if err != nil {
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
		return nil, ErrResponseNotReceived
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	entity := new(Entity)
	err = json.Unmarshal(content, entity)
	if err != nil {
		return nil, err
	}

	if entity.Type != "entity" {
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
func (m *moviebuff) GetResources(ctx context.Context, resourceType ResourceType, limit, page int) (*Resources, error) {
	u := "/resources/" + string(resourceType) + "?"
	if limit != 0 {
		u += "limit=" + strconv.Itoa(limit)
	}

	if page != 0 {
		u += "page=" + strconv.Itoa(page)
	}

	r, err := prepareRequest(ctx, m.HostURL, m.StaticToken, u)
	if err != nil {
		return nil, err
	}

	res, err := m.Client.Do(r)
	if err != nil {
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
		return nil, ErrResponseNotReceived
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resources := new(Resources)
	err = json.Unmarshal(content, resources)
	if err != nil {
		return nil, err
	}

	return resources, nil
}

// GetCertifications fetches a list of all certifications available on Moviebuff

//GetCertifications takes an optional argument country which can be the Qube Wire Cinemas country UUID or the ISO 2-digit code for this country, eg "IN". If country is provided, GetCertifications returns certifications available for the given country
//Pass empty value for country to get a list of all certifications across countries
func (m *moviebuff) GetCertifications(ctx context.Context, country string) ([]Certification, error) {
	r, err := prepareRequest(ctx, m.HostURL, m.StaticToken, "/certifications")
	if err != nil {
		return nil, err
	}
	if country != "" {
		addQueryParams(r, map[string]string{"country": country})
	}

	res, err := m.Client.Do(r)
	if err != nil {
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
		return nil, ErrResponseNotReceived
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	certifications := struct {
		Data []Certification `json:"data"`
	}{}
	err = json.Unmarshal(content, &certifications)
	if err != nil {
		return nil, err
	}

	return certifications.Data, nil
}

func (m *moviebuff) GetHolidayCalendar(ctx context.Context, countryID string) (*Calendar, error) {
	r, err := prepareRequest(ctx, m.HostURL, m.StaticToken, "/holidays/"+countryID)
	if err != nil {
		return nil, err
	}

	res, err := m.Client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusForbidden:
		return nil, ErrInvalidToken

	case http.StatusNotFound:
		return nil, ErrResourceDoesNotExist

	case http.StatusOK:
		calendarResp, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		calendarInfo := new(Calendar)
		err = json.Unmarshal(calendarResp, calendarInfo)
		if err != nil {
			return nil, err
		}
		return calendarInfo, nil

	default:
		return nil, ErrResponseNotReceived
	}

}
