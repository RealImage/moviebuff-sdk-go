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

	v2 "github.com/RealImage/moviebuff-sdk-go/v2"
)

// Type of resources supported by moviebuff
type ResourceType = v2.ResourceType

const (
	RESOURCE_TYPE_PEOPLE   ResourceType = v2.RESOURCE_TYPE_PEOPLE
	RESOURCE_TYPE_MOVIES   ResourceType = v2.RESOURCE_TYPE_MOVIES
	RESOURCE_TYPE_ENTITIES ResourceType = v2.RESOURCE_TYPE_ENTITIES
)

var (
	ErrInvalidToken         = v2.ErrInvalidToken
	ErrResponseNotReceived  = v2.ErrResponseNotReceived
	ErrResourceDoesNotExist = v2.ErrResourceDoesNotExist
)

// Moviebuff allows to access to information in moviebuff using resource ids.
type Moviebuff interface {
	GetMovie(id string) (*Movie, error)
	GetPerson(id string) (*Person, error)
	GetEntity(id string) (*Entity, error)
	GetResources(resourceType ResourceType, limit, page int) (*Resources, error)
	GetCertifications(country string) ([]Certification, error)
	GetHolidayCalendar(countryID string) (*Calendar, error)
}

type Config struct {
	HostURL     string
	StaticToken string
}

// Before accessing any API it need to be initialized.
// The Moviebuff is a service that offers information about movies, people, entities.
type moviebuff struct {
	Config
	v2 v2.Moviebuff
}

// New returns a Moviebuff interface.
func New(config Config) Moviebuff {
	v2Config := v2.Config{
		HostURL:     config.HostURL,
		StaticToken: config.StaticToken,
	}
	return &moviebuff{
		Config: config,
		v2:     v2.New(v2Config),
	}
}

// GetMovie fetch a movie and its basic details for given resource uuid.

// Instead of the UUID, this can also be the URL of the movie as seen on moviebuff.com, like 12-years-a-slave
// Details include release dates, certifications, cast, crew, trailers, posters, purchase links etc.
// Here movies may include feature films, documentaries, short films etc.
func (m *moviebuff) GetMovie(id string) (*Movie, error) {
	return m.v2.GetMovie(context.Background(), id)
}

// GetPerson fetch a person and his/her basic details with UUID.

// Instead of the UUID, this can also be the URL of the person as seen on moviebuff.com, like amitabh-bachchan
// The people in the database include actors, directors, support personnel, etc.
// Moviebuff aims to document most, if not all, of the individuals involved in a film.
func (m *moviebuff) GetPerson(id string) (*Person, error) {
	return m.v2.GetPerson(context.Background(), id)
}

// GetEntity fetch an entity and its basic details for UUID.

// Instead of the UUID, this can also be the URL of the company as seen on moviebuff.com: yash-raj-films .
// Entities are usually organizations like production companies, service providers, etc.
func (m *moviebuff) GetEntity(id string) (*Entity, error) {
	return m.v2.GetEntity(context.Background(), id)
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
	return m.v2.GetResources(context.Background(), resourceType, limit, page)
}

// GetCertifications fetches a list of all certifications available on Moviebuff

//GetCertifications takes an optional argument country which can be the Qube Wire Cinemas country UUID or the ISO 2-digit code for this country, eg "IN". If country is provided, GetCertifications returns certifications available for the given country
//Pass empty value for country to get a list of all certifications across countries
func (m *moviebuff) GetCertifications(country string) ([]Certification, error) {
	return m.v2.GetCertifications(context.Background(), country)
}

func (m *moviebuff) GetHolidayCalendar(countryID string) (*Calendar, error) {
	return m.v2.GetHolidayCalendar(context.Background(), countryID)

}
