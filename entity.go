package moviego

type Entity struct {
	Name   string   `json:"name"`
	Poster string   `json:"poster"`
	Tags   []string `json:"tags"`
	URL    string   `json:"url"`
	UUID   string   `json:"uuid"`
	Links  []struct {
		DisplayClass string `json:"displayClass"`
		Name         string `json:"name"`
		URL          string `json:"url"`
	} `json:"links"`
	Trivia         []string `json:"trivia"`
	Services       []string `json:"services"`
	CompanyProfile string   `json:"companyProfile"`
	Credits        []struct {
		Department string `json:"department"`
		Roles      []struct {
			Name         string `json:"name"`
			URL          string `json:"url"`
			ReleaseDates struct {
				IN string `json:"IN"`
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
	AlternateUrls []string `json:"alternateUrls"`
	Type          string   `json:"type"`
	Posters       []struct {
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
	APIPath      string `json:"apiPath"`
	MoviebuffURL string `json:"moviebuffUrl"`
}
