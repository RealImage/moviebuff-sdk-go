package moviego

type Resources struct {
	Prev string `json:"prev"`
	Data []struct {
		Name         string `json:"name"`
		URL          string `json:"url"`
		UUID         string `json:"uuid"`
		Type         string `json:"type"`
		Poster       string `json:"poster"`
		MoviebuffURL string `json:"moviebuffUrl"`
		APIPath      string `json:"apiPath"`
	} `json:"data"`
	Next string `json:"next"`
}
