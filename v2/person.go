package moviebuff

// Person contain basic details of a celebrity.
// The people in the database include actors, directors, support personnel, etc.
// Moviebuff aims to document most, if not all, of the individuals involved in a film.
type Person struct {
	// Name of the given person.
	Name string `json:"name"`

	// Profile photo of person
	Poster string `json:"poster"`

	// The URL for the poster of the given person.
	URL string `json:"url"`

	// The alternate URLs of the given person.
	AlternateUrls []string `json:"alternateUrls"`

	// An list of alternate names of the given person.
	AlternateNames []string `json:"alternateNames"`

	// Tags for the given person.
	Tags []string `json:"tags"`

	// value is fixed as person
	Type string `json:"type"`

	// The UUID of the person.
	UUID string `json:"uuid"`

	// The Biography of the given person.
	Biography string `json:"biography"`

	// Other popular name and titles of the given person.
	Akas []string `json:"akas"`

	// Height of the given person.
	Height string `json:"height"`

	// The date of birth of the given person.
	Birthday string `json:"birthday"`

	// The date of death of the given person.
	Deathday string `json:"deathday"`

	// The birth place of the given person.
	Birthplace string `json:"birthplace"`

	// An list containing the links of the given person's social media accounts, websites etc.
	Links []struct {
		DisplayClass string `json:"displayClass"`
		Name         string `json:"name"`
		URL          string `json:"url"`
	} `json:"links"`

	// The trivia of the given person.
	Trivia []string `json:"trivia"`

	// An list of posters of the given person.
	Posters []struct {
		Featured bool   `json:"featured"`
		URL      string `json:"url"`
		Key      string `json:"key"`
		Caption  string `json:"caption"`
		Type     string `json:"type"`
	} `json:"posters"`

	// An list of videos of the given person.
	Videos []struct {
		Featured  bool   `json:"featured"`
		URL       string `json:"url"`
		EmbedURL  string `json:"embedUrl"`
		Key       string `json:"key"`
		Caption   string `json:"caption"`
		Thumbnail string `json:"thumbnail"`
		Type      string `json:"type"`
	} `json:"videos"`

	// An list of stills of the given person.
	Stills []struct {
		Featured bool   `json:"featured"`
		URL      string `json:"url"`
		Key      string `json:"key"`
		Caption  string `json:"caption"`
		Type     string `json:"type"`
	} `json:"stills"`

	// The roles of the person in various movies grouped by department name.
	Credits []struct {
		Department string `json:"department"`
		Roles      []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
			// Release Dates mapped against their corresponding country code like "IN" : "2013-12-20".
			ReleaseDates map[string]string `json:"releaseDates"`

			// Certifications mapped against their corresponding country code like "IN" : "A".
			Certifications map[string]string `json:"certifications"`
			Language       string            `json:"language"`
			Type           string            `json:"type"`
			UUID           string            `json:"uuid"`
			Poster         string            `json:"poster"`
			MoviebuffURL   string            `json:"moviebuffUrl"`
			APIPath        string            `json:"apiPath"`
			Role           string            `json:"role"`
			Department     string            `json:"department"`
			Primary        bool              `json:"primary"`
			Character      string            `json:"character"`
		} `json:"roles"`
	} `json:"credits"`

	// The path to the person in the current version of the api.
	APIPath string `json:"apiPath"`

	// The url pointing to the moviebuff page of the person.
	MoviebuffURL string `json:"moviebuffUrl"`
}
