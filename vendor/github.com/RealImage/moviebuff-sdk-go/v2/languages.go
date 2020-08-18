package moviebuff

type Language struct {
	UUID       string   `json:"uuid"`
	Name       string   `json:"name"`
	ISO639_1   string   `json:"iso639_1"`
	ISO639_2   string   `json:"iso639_2"`
	IANA       string   `json:"iana"`
	NativeName string   `json:"native_name"`
	SpokedIn   []string `json:"spoken_in"`
}
