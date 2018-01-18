package moviego

type Person struct {
	Name           string      `json:"name"`
	Poster         string      `json:"poster"`
	URL            string      `json:"url"`
	AlternateUrls  []string    `json:"alternateUrls"`
	AlternateNames []string    `json:"alternateNames"`
	Tags           []string    `json:"tags"`
	Type           string      `json:"type"`
	UUID           string      `json:"uuid"`
	Biography      string      `json:"biography"`
	Akas           []string    `json:"akas"`
	Height         interface{} `json:"height"`
	Birthday       string      `json:"birthday"`
	Deathday       interface{} `json:"deathday"`
	Birthplace     string      `json:"birthplace"`
	Links          []struct {
		DisplayClass string `json:"displayClass"`
		Name         string `json:"name"`
		URL          string `json:"url"`
	} `json:"links"`
	Trivia  []string `json:"trivia"`
	Posters []struct {
		Featured bool   `json:"featured"`
		URL      string `json:"url"`
		Key      string `json:"key"`
		Caption  string `json:"caption"`
		Type     string `json:"type"`
	} `json:"posters"`
	Videos []struct {
		Featured  bool   `json:"featured"`
		URL       string `json:"url"`
		EmbedURL  string `json:"embedUrl"`
		Key       string `json:"key"`
		Caption   string `json:"caption"`
		Thumbnail string `json:"thumbnail"`
		Type      string `json:"type"`
	} `json:"videos"`
	Stills []struct {
		Featured bool   `json:"featured"`
		URL      string `json:"url"`
		Key      string `json:"key"`
		Caption  string `json:"caption"`
		Type     string `json:"type"`
	} `json:"stills"`
	Credits []struct {
		Department string `json:"department"`
		Roles      []struct {
			Name         string `json:"name"`
			URL          string `json:"url"`
			ReleaseDates struct {
			} `json:"releaseDates"`
			Certifications struct {
			} `json:"certifications"`
			Language     string      `json:"language"`
			Type         string      `json:"type"`
			UUID         string      `json:"uuid"`
			Poster       interface{} `json:"poster"`
			MoviebuffURL string      `json:"moviebuffUrl"`
			APIPath      string      `json:"apiPath"`
			Role         string      `json:"role"`
			Department   string      `json:"department"`
			Primary      bool        `json:"primary"`
			Character    interface{} `json:"character"`
		} `json:"roles"`
	} `json:"credits"`
	APIPath      string `json:"apiPath"`
	MoviebuffURL string `json:"moviebuffUrl"`
}
