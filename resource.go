package moviebuff

import v2 "github.com/RealImage/moviebuff-sdk-go/v2"

// Resources contain a paginated list of various resources like movies, people, entities in moviebuff.
// prev - The previous page in the paginated api. If null there are no more pages before this page.
// data - The array of items requested by the client. Each item in the array will contain the following fields
// name, url, uuid, type, poster, moviebuffUrl, apiPath.
// he next page in the paginated api. If null there are no more pages at the end.
type Resources = v2.Resources

//Resource contains information about specific type of the resource like movie, people, entity in moviebuff.
type Resource = v2.Resource
