package moviebuff

import (
	"context"
	"net/http"

	v1 "github.com/RealImage/moviebuff-sdk-go"
)

//Calendar - contains the info about the country's public holidays
type Calendar = v1.Calendar

//Holiday - contains id, date and name of the holiday
type Holiday = v1.Holiday

// Certification contains a readable code as well as UUID along with whether the certification indicates that the movie is safe for children and the country it is applicable for.
// A movie has multiple certifications, one for each country it is released in.
type Certification = v1.Certification

//Country has country information as available on Qube Wire Cinemas
type Country = v1.Country

// Entity contain basic details of an entity.
// Entities are usually organizations like production companies, service providers, etc.
type Entity = v1.Entity

// Movie contain basic movie details like release dates, certifications, cast, crew, trailers, posters,
// purchase links etc.
// Here movies may include feature films, documentaries, short films etc.
type Movie = v1.Movie

// Person contain basic details of a celebrity.
// The people in the database include actors, directors, support personnel, etc.
// Moviebuff aims to document most, if not all, of the individuals involved in a film.
type Person = v1.Person

// Resources contain a paginated list of various resources like movies, people, entities in moviebuff.
// prev - The previous page in the paginated api. If null there are no more pages before this page.
// data - The array of items requested by the client. Each item in the array will contain the following fields
// name, url, uuid, type, poster, moviebuffUrl, apiPath.
// he next page in the paginated api. If null there are no more pages at the end.
type Resources = v1.Resources

//Resource contains information about specific type of the resource like movie, people, entity in moviebuff.
type Resource = v1.Resource

const apiKey = "X-Api-Key"

func prepareRequest(ctx context.Context, hostURL, staticToken, path string) (*http.Request, error) {

	r, err := http.NewRequestWithContext(ctx, http.MethodGet, hostURL+path, nil)
	if err != nil {
		return nil, err
	}

	r.Header.Add(apiKey, staticToken)
	return r, nil
}

func addQueryParams(r *http.Request, queryParams map[string]string) {
	q := r.URL.Query()
	for k, v := range queryParams {
		q.Add(k, v)
	}
	r.URL.RawQuery = q.Encode()
}
