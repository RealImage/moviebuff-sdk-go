package moviebuff

// Resources contain a paginated list of various resources like movies, people, entities in moviebuff.
// prev - The previous page in the paginated api. If null there are no more pages before this page.
// data - The array of items requested by the client. Each item in the array will contain the following fields
// name, url, uuid, type, poster, moviebuffUrl, apiPath.
// he next page in the paginated api. If null there are no more pages at the end.
type Resources struct {
	Prev string     `json:"prev"`
	Data []Resource `json:"data"`
	Next string     `json:"next"`
}

//Resource contains information about specific type of the resource like movie, people, entity in moviebuff.
type Resource struct {
	// Name of the resource
	Name string `json:"name"`

	// slug url of the resource
	URL string `json:"url"`

	// uuid of the resource
	UUID string `json:"uuid"`

	// type of resource. Possible type: movies, people, entities, theatres
	Type string `json:"type"`

	// Poster url of the resource
	Poster string `json:"poster"`

	// The path to the person in the current version of the api.
	APIPath string `json:"apiPath"`

	// The url pointing to the moviebuff page of the person.
	MoviebuffURL string `json:"moviebuffUrl"`
}
