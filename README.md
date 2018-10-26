# moviebuff-sdk-go

Go SDK for [Moviebuff](www.moviebuff.com)

The Moviebuff is a service that offers information about
 movies, people and entities.

All resources are identified by a UUID,
which Moviebuff.com uniquely and randomly generates.
Since it may be difficult to get the Moviebuff UUID of
 any resource without prior knowledge, the API also
 allows substitution of the UUID with the URL
 identifier of the resource. For instance, if when
 attempting to get information for 12 Years a Slave
  (moviebuff.com/12-years-a-slave), either its UUID
  or the identifier 12-years-a-slave may be used.

Sample usage:

    movie, err := new(moviebuff.Moviebuff).Init(os.Getenv("API_TOKEN"), nil).GetMovie("padmaavat")
    if err!=nil{
        return err
    }

    log.Println("Padmaavat was release in year ", movie.GetEarliestReleaseYear())

## Information Available via Moviebuff

For Detailed Code Documentation click [here](https://godoc.org/github.com/RealImage/moviebuff-sdk-go) .

* Get Movie Information

    Movie contain basic movie details like release dates,
     certifications, cast, crew, trailers, posters,
     purchase links etc. Here movies may include feature films,
     documentaries, short films etc.

* Get Person Information

    Person contain basic details of a celebrity.
    The people in the database include actors, directors,
    support personnel, etc. Moviebuff aims to document most,
     if not all, of the individuals involved in a film.

* Get Entity Information

    Entity contain basic details of an entity. Entities are
     usually organizations like production companies,
     service providers, etc.

* Get Certifications
    A Certification contains a readable code as well as UUID along with whether the certification indicates that the movie is safe for children and the country the certification is applicable for. A movie has multiple certifications, one for each country it is released in.

