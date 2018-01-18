package moviego

type Movie struct {
	Name          string   `json:"name"`
	Poster        string   `json:"poster"`
	AlternateUrls []string `json:"alternateUrls"`
	URL           string   `json:"url"`
	Type          string   `json:"type"`
	UUID          string   `json:"uuid"`
	ReleaseDates  struct {
		IN string `json:"IN"`
	} `json:"releaseDates"`
	Certifications struct {
		IN string `json:"IN"`
	} `json:"certifications"`
	Language      string   `json:"language"`
	FilmType      string   `json:"filmType"`
	Featured      bool     `json:"featured"`
	Synopsis      string   `json:"synopsis"`
	ShortSynopsis string   `json:"shortSynopsis"`
	Storyline     string   `json:"storyline"`
	Genres        []string `json:"genres"`
	LocalName     string   `json:"localName"`
	RunningTime   int      `json:"runningTime"`
	Trailer       struct {
		Featured  bool   `json:"featured"`
		URL       string `json:"url"`
		EmbedURL  string `json:"embedUrl"`
		Key       string `json:"key"`
		Caption   string `json:"caption"`
		Thumbnail string `json:"thumbnail"`
		Type      string `json:"type"`
	} `json:"trailer"`
	AlternateTitles []string `json:"alternateTitles"`
	Taglines        []string `json:"taglines"`
	Links           []struct {
		DisplayClass string `json:"displayClass"`
		Name         string `json:"name"`
		URL          string `json:"url"`
	} `json:"links"`
	PurchaseLinks []struct {
		DisplayClass string `json:"displayClass"`
		Name         string `json:"name"`
		URL          string `json:"url"`
	} `json:"purchaseLinks"`
	TechDetails []struct {
		Name string   `json:"name"`
		Data []string `json:"data"`
	} `json:"techDetails"`
	Trivia      []string `json:"trivia"`
	MovieRating struct {
		Value string `json:"value"`
		Count int    `json:"count"`
	} `json:"movieRating"`
	MusicRating struct {
		Value string `json:"value"`
		Count int    `json:"count"`
	} `json:"musicRating"`
	Cast []struct {
		Name         string `json:"name"`
		Poster       string `json:"poster"`
		Type         string `json:"type"`
		URL          string `json:"url"`
		UUID         string `json:"uuid"`
		Role         string `json:"role"`
		Department   string `json:"department"`
		Primary      bool   `json:"primary"`
		Character    string `json:"character"`
		MoviebuffURL string `json:"moviebuffUrl"`
		APIPath      string `json:"apiPath"`
	} `json:"cast"`
	Crew []struct {
		Department string `json:"department"`
		Roles      []struct {
			Name         string      `json:"name"`
			Poster       interface{} `json:"poster"`
			Type         string      `json:"type"`
			URL          string      `json:"url"`
			UUID         string      `json:"uuid"`
			Role         string      `json:"role"`
			Department   string      `json:"department"`
			Primary      bool        `json:"primary"`
			Character    interface{} `json:"character"`
			MoviebuffURL string      `json:"moviebuffUrl"`
			APIPath      string      `json:"apiPath"`
		} `json:"roles"`
	} `json:"crew"`
	MusicLabels []struct {
		Name         string      `json:"name"`
		Poster       interface{} `json:"poster"`
		URL          string      `json:"url"`
		UUID         string      `json:"uuid"`
		Type         string      `json:"type"`
		APIPath      string      `json:"apiPath"`
		MoviebuffURL string      `json:"moviebuffUrl"`
	} `json:"musicLabels"`
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
	News []struct {
		Poster  string `json:"poster"`
		Summary string `json:"summary"`
		Date    string `json:"date"`
		URL     string `json:"url"`
		Writer  string `json:"writer"`
	} `json:"news"`
	Connections []struct {
		Name         string `json:"name"`
		URL          string `json:"url"`
		ReleaseDates struct {
			IN string `json:"IN"`
		} `json:"releaseDates"`
		Certifications struct {
			IN string `json:"IN"`
		} `json:"certifications"`
		Language       string      `json:"language"`
		Type           string      `json:"type"`
		UUID           string      `json:"uuid"`
		Poster         interface{} `json:"poster"`
		MoviebuffURL   string      `json:"moviebuffUrl"`
		APIPath        string      `json:"apiPath"`
		ConnectionType string      `json:"connectionType"`
	} `json:"connections"`
	ReleaseStatuses struct {
		AE string `json:"AE"`
	} `json:"releaseStatuses"`
	MoviebuffURL string `json:"moviebuffUrl"`
	APIPath      string `json:"apiPath"`
}
