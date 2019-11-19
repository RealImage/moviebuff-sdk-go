package moviebuff

// Entity contain basic details of an entity.
// Entities are usually organizations like production companies, service providers, etc.
type Entity struct {
	// Name of the entity.
	Name string `json:"name"`

	// Poster image of the entity.
	Poster string `json:"poster"`

	// Tags of the entity.
	Tags []string `json:"tags"`

	// The primary slug url of the entity.
	URL string `json:"url"`

	// The UUID of the entity.
	UUID string `json:"uuid"`

	// The links to various social media platforms, websites of the entity.
	Links []struct {
		DisplayClass string `json:"displayClass"`
		Name         string `json:"name"`
		URL          string `json:"url"`
	} `json:"links"`

	// Trivia of the entity.
	Trivia []string `json:"trivia"`

	// The main services provided by the entity.
	Services []string `json:"services"`

	// The profile of the entity as a company.
	CompanyProfile string `json:"companyProfile"`

	// The credits for the entity from various movies grouped by the department.
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
			Poster         interface{}       `json:"poster"`
			MoviebuffURL   string            `json:"moviebuffUrl"`
			APIPath        string            `json:"apiPath"`
			Role           string            `json:"role"`
			Department     string            `json:"department"`
			Primary        bool              `json:"primary"`
			Character      interface{}       `json:"character"`
		} `json:"roles"`
	} `json:"credits"`

	// The Alternate slug URLs of the entity.
	AlternateUrls []string `json:"alternateUrls"`

	// value is fixed as entity
	Type string `json:"type"`

	// An list of posters of the entity.
	Posters []struct {
		Featured bool   `json:"featured"`
		URL      string `json:"url"`
		Key      string `json:"key"`
		Caption  string `json:"caption"`
		Type     string `json:"type"`
	} `json:"posters"`

	// An list of videos of the entity.
	Videos []struct {
		Featured  bool   `json:"featured"`
		URL       string `json:"url"`
		EmbedURL  string `json:"embedUrl"`
		Key       string `json:"key"`
		Caption   string `json:"caption"`
		Thumbnail string `json:"thumbnail"`
		Type      string `json:"type"`
	} `json:"videos"`

	// An list of stills of the entity.
	Stills []struct {
		Featured bool   `json:"featured"`
		URL      string `json:"url"`
		Key      string `json:"key"`
		Caption  string `json:"caption"`
		Type     string `json:"type"`
	} `json:"stills"`

	// The path to the entity in the current version of the API.
	APIPath string `json:"apiPath"`

	// The url to the moviebuff page of the entity.
	MoviebuffURL string `json:"moviebuffUrl"`
}
