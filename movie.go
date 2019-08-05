package moviebuff

import (
	"time"
)

// Movie contain basic movie details like release dates, certifications, cast, crew, trailers, posters,
// purchase links etc.
// Here movies may include feature films, documentaries, short films etc.
type Movie struct {
	// Name of the movie.
	Name string `json:"name"`

	// Poster url of the movie.
	Poster string `json:"poster"`

	// Alternate URLS of the movie.
	AlternateUrls []string `json:"alternateUrls"`

	// Primary URL of the movie.
	URL string `json:"url"`

	// Type of the resource.
	Type string `json:"type"`

	// UUID of the movie.
	UUID string `json:"uuid"`

	// Release Dates mapped against their corresponding country code like "IN" : "2013-12-20".
	ReleaseDates map[string]string `json:"releaseDates"`

	// Certifications mapped against their corresponding country code like "IN" : "A".
	Certifications map[string]string `json:"certifications"`

	// The primary language of the movie.
	Language string `json:"language"`

	// The language details of the movie.
	LanguageData struct {
		Name string `json:"name"`
		UUID string `json:"uuid"`
	} `json:"languageData"`

	// The type of the movie like Feature Film, Short Film etc.
	FilmType string `json:"filmType"`

	Featured bool `json:"featured"`

	// The synopsis of the movie.
	Synopsis string `json:"synopsis"`

	// The short synopsis of the movie.
	ShortSynopsis string `json:"shortSynopsis"`

	// The storyline of the movie.
	Storyline string `json:"storyline"`

	// An list of genres of the movie.
	Genres []string `json:"genres"`

	// The name of the movie in the local language.
	LocalName string `json:"localName"`

	// The running time of the movie in seconds.
	RunningTime int `json:"runningTime"`

	// Details of the trailer of the movie.
	Trailer struct {
		Featured  bool   `json:"featured"`
		URL       string `json:"url"`
		EmbedURL  string `json:"embedUrl"`
		Key       string `json:"key"`
		Caption   string `json:"caption"`
		Thumbnail string `json:"thumbnail"`
		Type      string `json:"type"`
	} `json:"trailer"`

	// An list containing the alternate titles of the movie.
	AlternateTitles []string `json:"alternateTitles"`

	// An list containing the taglines of the movie.
	TagLines []string `json:"taglines"`

	// An list containing links to movie's twitter, facebook pages etc.
	Links []struct {
		DisplayClass string `json:"displayClass"`
		Name         string `json:"name"`
		URL          string `json:"url"`
	} `json:"links"`

	// An list containing the links to various purchase channels of movie, its songs etc.
	PurchaseLinks []struct {
		DisplayClass string `json:"displayClass"`
		Name         string `json:"name"`
		URL          string `json:"url"`
	} `json:"purchaseLinks"`

	// An list of objects each having a name and data. Here name contains the type of the tech data and data which may
	// be an list of strings or a string.
	TechDetails []struct {
		Name string      `json:"name"`
		Data interface{} `json:"data"`
	} `json:"techDetails"`

	// An list containing the trivia of the movie.
	Trivia []string `json:"trivia"`

	// An object containing the value and the count of the ratings for the movie.
	MovieRating struct {
		Value string `json:"value"`
		Count int    `json:"count"`
	} `json:"movieRating"`

	// An object containing the value and the count of the ratings for the movie's tracks.
	MusicRating struct {
		Value string `json:"value"`
		Count int    `json:"count"`
	} `json:"musicRating"`

	// An list containing the cast of the movie.
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

	// An list containing the crew of the movie grouped by department.
	Crew []struct {
		Department string `json:"department"`
		Roles      []struct {
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
		} `json:"roles"`
	} `json:"crew"`

	// An list containing the music labels for the movies tracks.
	MusicLabels []struct {
		Name         string `json:"name"`
		Poster       string `json:"poster"`
		URL          string `json:"url"`
		UUID         string `json:"uuid"`
		Type         string `json:"type"`
		APIPath      string `json:"apiPath"`
		MoviebuffURL string `json:"moviebuffUrl"`
	} `json:"musicLabels"`

	// An list containing the posters of the movie.
	Posters []struct {
		Featured bool   `json:"featured"`
		URL      string `json:"url"`
		Key      string `json:"key"`
		Caption  string `json:"caption"`
		Type     string `json:"type"`
	} `json:"posters"`

	// An list containing the videos of the movie.
	Videos []struct {
		Featured  bool   `json:"featured"`
		URL       string `json:"url"`
		EmbedURL  string `json:"embedUrl"`
		Key       string `json:"key"`
		Caption   string `json:"caption"`
		Thumbnail string `json:"thumbnail"`
		Type      string `json:"type"`
	} `json:"videos"`

	// An list containing the stills of the movie.
	Stills []struct {
		Featured bool   `json:"featured"`
		URL      string `json:"url"`
		Key      string `json:"key"`
		Caption  string `json:"caption"`
		Type     string `json:"type"`
	} `json:"stills"`

	// An list containing the news articles related to the movie.
	News []struct {
		Poster  string `json:"poster"`
		Summary string `json:"summary"`
		Date    string `json:"date"`
		URL     string `json:"url"`
		Writer  string `json:"writer"`
	} `json:"news"`

	// An list containing the various connections of the movie with other movie.
	// It is an list of objects with each object having a connectionType that
	// specifies how the two movies are related to each other.
	Connections []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
		// Release Dates mapped against their corresponding country code like "IN" : "2013-12-20".
		ReleaseDates map[string]string `json:"releaseDates"`

		// Certifications mapped against their corresponding country code like "IN" : "A".
		Certifications map[string]string `json:"certifications"`

		Language       string `json:"language"`
		Type           string `json:"type"`
		UUID           string `json:"uuid"`
		Poster         string `json:"poster"`
		MoviebuffURL   string `json:"moviebuffUrl"`
		APIPath        string `json:"apiPath"`
		ConnectionType string `json:"connectionType"`
	} `json:"connections"`

	// An object where each release status is mapped against its corresponding country code.
	ReleaseStatuses struct {
		AE string `json:"AE"`
	} `json:"releaseStatuses"`

	ThirdPartyIdentifiers []struct {
		IDs    []string `json:"ids"`
		Source struct {
			UUID string `json:"uuid"`
			Name string `json:"name"`
		} `json:"source"`
	} `json:"thirdPartyIdentifiers"`

	MoviebuffURL string `json:"moviebuffUrl"`
	APIPath      string `json:"apiPath"`
}

// GetEarliestReleaseYear return the year at which movie was first released anywhere in the world.
// Returns 0 if release date is not available
func (m *Movie) GetEarliestReleaseYear() (releaseYear int) {
	for _, v := range m.ReleaseDates {
		if t, err := time.Parse("2006-01-02", v); err == nil && (t.Year() < releaseYear || releaseYear == 0) {
			releaseYear = t.Year()
		}
	}
	return
}

// GetThirdPartyIDsBySource returns third-party IDs for provided source.
// Returns nil if source ID is not available
func (m *Movie) GetThirdPartyIDsBySource(sourceID string) []string {
	for _, id := range m.ThirdPartyIdentifiers {
		if id.Source.UUID == sourceID {
			return id.IDs
		}
	}
	return nil
}
